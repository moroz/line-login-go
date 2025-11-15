package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"net/url"
)

type oauth2Controller struct {
	clientId     string
	clientSecret string
}

func OAuth2Controller(clientId, clientSecret string) *oauth2Controller {
	return &oauth2Controller{clientId, clientSecret}
}

const LineBaseUrl = "https://access.line.me/oauth2/v2.1/authorize"

// ?response_type=code&client_id=1234567890&redirect_uri=https%3A%2F%2Fexample.com%2Fauth%3Fkey%3Dvalue&state=12345abcde&scope=profile%20openid&nonce=09876xyz

func generateState() string {
	var bytes = make([]byte, 4)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (c *oauth2Controller) Init(w http.ResponseWriter, r *http.Request) {
	params := url.Values{
		"response_type": {"code"},
		"client_id":     {c.clientId},
		"redirect_uri":  {"http://localhost:3000/oauth/line/callback"},
		"state":         {generateState()},
		"scope":         {"profile openid"},
	}

	url := LineBaseUrl + "?" + params.Encode()
	http.Redirect(w, r, url, http.StatusFound)
}
