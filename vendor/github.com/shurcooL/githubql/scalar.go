package githubql

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// Note: These custom types are meant to be used in queries, but it's not required.
// They're here for convenience and documentation. If you use the base Go types,
// things will still work.
//
// TODO: In Go 1.9, consider using type aliases instead (for extra simplicity
//       and convenience).

type (
	// Boolean represents true or false values.
	Boolean bool

	// DateTime is an ISO-8601 encoded UTC date.
	DateTime struct{ time.Time }

	// Float represents signed double-precision fractional values as
	// specified by IEEE 754.
	Float float64

	// GitObjectID is a Git object ID. For example,
	// "912ec1990bd09f8fc128c3fa6b59105085aabc03".
	GitObjectID string

	// GitTimestamp is an ISO-8601 encoded date.
	// Unlike the DateTime type, GitTimestamp is not converted in UTC.
	GitTimestamp struct{ time.Time }

	// HTML is a string containing HTML code.
	HTML string

	// ID represents a unique identifier that is Base64 obfuscated. It
	// is often used to refetch an object or as key for a cache. The ID
	// type appears in a JSON response as a String; however, it is not
	// intended to be human-readable. When expected as an input type,
	// any string (such as "VXNlci0xMA==") or integer (such as 4) input
	// value will be accepted as an ID.
	ID interface{}

	// Int represents non-fractional signed whole numeric values.
	// Int can represent values between -(2^31) and 2^31 - 1.
	Int int32

	// String represents textual data as UTF-8 character sequences.
	// This type is most often used by GraphQL to represent free-form
	// human-readable text.
	String string

	// URI is an RFC 3986, RFC 3987, and RFC 6570 (level 4) compliant URI.
	URI struct{ *url.URL }

	// X509Certificate is a valid x509 certificate.
	X509Certificate struct{ *x509.Certificate }
)

// MarshalJSON implements the json.Marshaler interface.
// The URI is a quoted string.
func (u URI) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The URI is expected to be a quoted string.
func (u *URI) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	u.URL, err = url.Parse(s)
	return err
}

// MarshalJSON implements the json.Marshaler interface.
func (x X509Certificate) MarshalJSON() ([]byte, error) {
	// TODO: Implement.
	return nil, fmt.Errorf("X509Certificate.MarshalJSON: not implemented")
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (x *X509Certificate) UnmarshalJSON(data []byte) error {
	// TODO: Implement.
	return fmt.Errorf("X509Certificate.UnmarshalJSON: not implemented")
}

// NewBoolean is a helper to make a new *Boolean.
func NewBoolean(v Boolean) *Boolean { return &v }

// NewDateTime is a helper to make a new *DateTime.
func NewDateTime(v DateTime) *DateTime { return &v }

// NewFloat is a helper to make a new *Float.
func NewFloat(v Float) *Float { return &v }

// NewGitObjectID is a helper to make a new *GitObjectID.
func NewGitObjectID(v GitObjectID) *GitObjectID { return &v }

// NewGitTimestamp is a helper to make a new *GitTimestamp.
func NewGitTimestamp(v GitTimestamp) *GitTimestamp { return &v }

// NewHTML is a helper to make a new *HTML.
func NewHTML(v HTML) *HTML { return &v }

// NewID is a helper to make a new *ID.
func NewID(v ID) *ID { return &v }

// NewInt is a helper to make a new *Int.
func NewInt(v Int) *Int { return &v }

// NewString is a helper to make a new *String.
func NewString(v String) *String { return &v }

// NewURI is a helper to make a new *URI.
func NewURI(v URI) *URI { return &v }

// NewX509Certificate is a helper to make a new *X509Certificate.
func NewX509Certificate(v X509Certificate) *X509Certificate { return &v }
