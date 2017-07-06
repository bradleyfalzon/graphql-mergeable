
# Introduction

`githubql-mergeable` is an example Go application using GitHub's [GraphQL](https://developer.github.com/v4/) API via
Dmitri Shuralyov's ([@shurcool](https://github.com/shurcooL)) [githubql](https://github.com/shurcooL/githubql) library.

It's purpose was to test and illustrate the `githubql` library, it's not designed to be a useful program.

Running `githubql-mergeable` will show you all open GitHub Pull Requests for a user and their mergeable state.

# Running

Fetch the application using go get.

```
go get -u github.com/bradleyfalzon/githubql-mergeable
```

GitHub's GraphQL API requires authentication, the simplest way is to use a [Personal Access
Token](https://github.com/settings/tokens), and setting the environment's `GITHUB_TOKEN` to this value.

```
export GITHUB_TOKEN=aabbcc...ddeeff
```

`githubql-mergeable` then takes the GitHub's user's login as the first and only argument, and it's required.

```
githubql-mergeable bradleyfalzon
```

# Example

```
MERGEABLE   | "2017-01-31 00:07:45 +0000 UTC" | "Feed + Folder View / Refreshing (#30)"                             | https://github.com/bradleyfalzon/hydrocarbon/pull/1
MERGEABLE   | "2017-06-23 10:13:55 +0000 UTC" | "Stacktrace the culprit should be the received error not the cause" | https://github.com/evalphobia/logrus_sentry/pull/45
CONFLICTING | "2017-06-30 11:25:41 +0000 UTC" | "lots files"                                                        | https://github.com/bf-test/gopherci-dev1/pull/136
```

# License

Licensed under Creative Commons Zero (CC0) license.
