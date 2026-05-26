package protocol

import (
	"net/url"
)

// apiVersion is the version of the Sentry API.
const apiVersion = "7"

type scheme string

const (
	SchemeHTTP  scheme = "http"
	SchemeHTTPS scheme = "https"
)

func (scheme scheme) defaultPort() int { _ = "STUB: not implemented"; return 0 }

// DsnParseError represents an error that occurs if a Sentry
// DSN cannot be parsed.
type DsnParseError struct {
	Message string
}

func (e DsnParseError) Error() string { _ = "STUB: not implemented"; return "" }

// Dsn is used as the remote address source to client transport.
type Dsn struct {
	scheme    scheme
	publicKey string
	secretKey string
	host      string
	port      int
	path      string
	projectID string
	orgID     uint64
}

// NewDsn creates a Dsn by parsing rawURL. Most users will never call this
// function directly. It is provided for use in custom Transport
// implementations.
func NewDsn(rawURL string) (*Dsn, error) {
	_ = "STUB: not implemented"
	// Parse
	return nil, nil
}

// Scheme

// PublicKey

// SecretKey

// Host

// OrgID (optional)

// Port

// ProjectID

// Path

// String formats Dsn struct into a valid string url.
func (dsn Dsn) String() string { _ = "STUB: not implemented"; return "" }

// Get the scheme of the DSN.
func (dsn Dsn) GetScheme() string { _ = "STUB: not implemented"; return "" }

// Get the public key of the DSN.
func (dsn Dsn) GetPublicKey() string { _ = "STUB: not implemented"; return "" }

// Get the secret key of the DSN.
func (dsn Dsn) GetSecretKey() string { _ = "STUB: not implemented"; return "" }

// Get the host of the DSN.
func (dsn Dsn) GetHost() string {
	_ = "STUB: not implemented"

	// Get the port of the DSN.
	return ""
}

func (dsn Dsn) GetPort() int {
	_ = "STUB: not implemented"

	// Get the path of the DSN.
	return 0
}

func (dsn Dsn) GetPath() string {
	_ = "STUB: not implemented"

	// Get the project ID of the DSN.
	return ""
}

func (dsn Dsn) GetProjectID() string { _ = "STUB: not implemented"; return "" }

// GetOrgID returns the orgID that was parsed from the DSN.
func (dsn Dsn) GetOrgID() uint64 {
	_ = "STUB: not implemented"

	// SetOrgID sets the orgID used for trace continuation.
	//
	// This function is used for overriding the orgID parsed from the DSN.
	return 0
}

func (dsn *Dsn) SetOrgID(orgID uint64) {
	_ = "STUB: not implemented"

	// GetAPIURL returns the URL of the envelope endpoint of the project
	// associated with the DSN.
	return
}

func (dsn Dsn) GetAPIURL() *url.URL { _ = "STUB: not implemented"; return nil }

// RequestHeaders returns all the necessary headers that have to be used in the transport when sending events
// to the /store endpoint.
//
// Deprecated: This method shall only be used if you want to implement your own transport that sends events to
// the /store endpoint. If you're using the transport provided by the SDK, all necessary headers to authenticate
// against the /envelope endpoint are added automatically.
func (dsn Dsn) RequestHeaders(sdkVersion string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// MarshalJSON converts the Dsn struct to JSON.
func (dsn Dsn) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON converts JSON data to the Dsn struct.
func (dsn *Dsn) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
