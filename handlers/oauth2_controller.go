package handlers

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/moroz/line-login-go/config"
)

type oauth2Controller struct {
	clientId     string
	clientSecret string
}

func OAuth2Controller(clientId, clientSecret string) *oauth2Controller {
	return &oauth2Controller{clientId, clientSecret}
}

const LineBaseUrl = "https://access.line.me/oauth2/v2.1/authorize"

func generateState() string {
	var bytes = make([]byte, 4)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (c *oauth2Controller) Init(w http.ResponseWriter, r *http.Request) {
	params := url.Values{
		"response_type": {"code"},
		"client_id":     {c.clientId},
		"redirect_uri":  {config.LineCallbackUri},
		"state":         {generateState()},
		"scope":         {"profile openid"},
	}

	url := LineBaseUrl + "?" + params.Encode()
	http.Redirect(w, r, url, http.StatusFound)
}

const LineRedeemAccessTokenUrl = "https://api.line.me/oauth2/v2.1/token"

type LineAccessTokenResponse struct {
	IDToken     string `json:"id_token"`
	AccessToken string `json:"access_token"`
}

func (c *oauth2Controller) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	params := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"client_id":     {c.clientId},
		"client_secret": {c.clientSecret},
		"redirect_uri":  {config.LineCallbackUri},
	}

	body := bytes.NewBufferString(params.Encode())

	req, _ := http.NewRequest("POST", LineRedeemAccessTokenUrl, body)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error obtaining access token: %s", err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	var result LineAccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error obtaining access token: %s", err)
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprint(w, result.IDToken)
}
