package transport

import (
	"net/http"

	authutils "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/auth/utils"
)

type AuthHeadersTransport struct {
	T http.RoundTripper
}

// NewAuthHeadersTransport creates a new transport to attach context authentication values into request headers
func NewAuthHeadersTransport() Middleware {
	return func(nxt http.RoundTripper) http.RoundTripper {
		return &AuthHeadersTransport{
			T: nxt,
		}
	}
}

func (t *AuthHeadersTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	authutils.AddAuthorizationHeader(req)
	return t.T.RoundTrip(req)
}
