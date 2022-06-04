package clientinterface

import "net/http"

// Client - interface for the built-in `http.Client` so we can mock and test.
//
// Documentation for `http.Client`:
//
// A Client is an HTTP client. Its zero value (DefaultClient) is a
// usable client that uses DefaultTransport.
//
// The Client's Transport typically has internal state (cached TCP
// connections), so Clients should be reused instead of created as
// needed. Clients are safe for concurrent use by multiple goroutines.
//
// A Client is higher-level than a RoundTripper (such as Transport)
// and additionally handles HTTP details such as cookies and
// redirects.
//
// When following redirects, the Client will forward all headers set on the
// initial Request except:
//
// • when forwarding sensitive headers like "Authorization",
// "WWW-Authenticate", and "Cookie" to untrusted targets.
// These headers will be ignored when following a redirect to a domain
// that is not a subdomain match or exact match of the initial domain.
// For example, a redirect from "foo.com" to either "foo.com" or "sub.foo.com"
// will forward the sensitive headers, but a redirect to "bar.com" will not.
//
// • when forwarding the "Cookie" header with a non-nil cookie Jar.
// Since each redirect may mutate the state of the cookie jar,
// a redirect may possibly alter a cookie set in the initial request.
// When forwarding the "Cookie" header, any mutated cookies will be omitted,
// with the expectation that the Jar will insert those mutated cookies
// with the updated values (assuming the origin matches).
// If Jar is nil, the initial cookies are forwarded without change.
//
type Client interface {

	// Do sends an HTTP request and returns an HTTP response, following
	// policy (such as redirects, cookies, auth) as configured on the
	// client.
	//
	// An error is returned if caused by client policy (such as
	// CheckRedirect), or failure to speak HTTP (such as a network
	// connectivity problem). A non-2xx status code doesn't cause an
	// error.
	//
	// If the returned error is nil, the Response will contain a non-nil
	// Body which the user is expected to close. If the Body is not both
	// read to EOF and closed, the Client's underlying RoundTripper
	// (typically Transport) may not be able to re-use a persistent TCP
	// connection to the server for a subsequent "keep-alive" request.
	//
	// The request Body, if non-nil, will be closed by the underlying
	// Transport, even on errors.
	//
	// On error, any Response can be ignored. A non-nil Response with a
	// non-nil error only occurs when CheckRedirect fails, and even then
	// the returned Response.Body is already closed.
	//
	// Generally Get, Post, or PostForm will be used instead of Do.
	//
	// If the server replies with a redirect, the Client first uses the
	// CheckRedirect function to determine whether the redirect should be
	// followed. If permitted, a 301, 302, or 303 redirect causes
	// subsequent requests to use HTTP method GET
	// (or HEAD if the original request was HEAD), with no body.
	// A 307 or 308 redirect preserves the original HTTP method and body,
	// provided that the Request.GetBody function is defined.
	// The NewRequest function automatically sets GetBody for common
	// standard library body types.
	//
	// Any returned error will be of type *url.Error. The url.Error
	// value's Timeout method will report true if the request timed out.
	Do(req *http.Request) (*http.Response, error)
}
