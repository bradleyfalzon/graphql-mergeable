package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/shurcooL/githubql"

	"golang.org/x/oauth2"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("%v <GitHub user's login>", os.Args[0])
	}

	if os.Getenv("GITHUB_TOKEN") == "" {
		log.Fatal("GITHUB_TOKEN environment variable is not set")
	}

	err := run(os.Getenv("GITHUB_TOKEN"), os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}

func run(token, login string) error {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	ctx := context.Background()
	httpClient := oauth2.NewClient(ctx, src)
	client := githubql.NewClient(httpClient)

	// GraphQL query:
	/*
		query {
			user(login:"bradleyfalzon") {
				pullRequests(states:[OPEN], first:10, orderBy: {field: UPDATED_AT, direction: ASC}) {
					nodes {
						title
						url
						mergeable
					}
				}
			}
		}
	*/

	type pull struct {
		Title     githubql.String
		URL       githubql.String
		UpdatedAt githubql.DateTime
		Mergeable githubql.MergeableState
	}

	var pullsQuery struct {
		User struct {
			PullRequests struct {
				Nodes    []pull
				PageInfo struct {
					EndCursor   githubql.String
					HasNextPage githubql.Boolean
				}
			} `graphql:"pullRequests(states: [OPEN], first: 10, orderBy: {field: UPDATED_AT, direction: ASC}, after: $pullsCursor)"`
		} `graphql:"user(login: $login)"`
	}

	variables := map[string]interface{}{
		"login":       githubql.String(login),
		"pullsCursor": (*githubql.String)(nil),
	}

	// Paginate through results.
	var pulls []pull
	for {
		err := client.Query(ctx, &pullsQuery, variables)
		if err != nil {
			return err
		}
		pulls = append(pulls, pullsQuery.User.PullRequests.Nodes...)
		if !pullsQuery.User.PullRequests.PageInfo.HasNextPage {
			break
		}
		variables["pullsCursor"] = githubql.NewString(pullsQuery.User.PullRequests.PageInfo.EndCursor)
	}

	if len(pulls) == 0 {
		fmt.Println("No open pull requests found")
		return nil
	}

	// Display results.
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	for _, pull := range pulls {
		fmt.Fprintf(w, "%v\t %q\t %q\t %v\n", pull.Mergeable, pull.UpdatedAt, pull.Title, pull.URL)
	}
	w.Flush()

	return nil
}
