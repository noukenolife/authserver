package adaptor

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/noukenolife/authserver/helper"
	"github.com/noukenolife/authserver/infrastructure/errors"
)

type TwitterGetAccessToken struct {
	port.GetAccessToken
	HTTPClient helper.HTTPClientInterface
}

func (s TwitterGetAccessToken) Invoke(input port.GetAccessTokenInput) (output port.GetAccessTokenOutput, err error) {
	// Create a request url
	accessTokenRequestURL, _ := url.Parse("https://api.twitter.com/oauth/access_token")
	accessTokenRequestQuery := accessTokenRequestURL.Query()
	accessTokenRequestQuery.Set("oauth_token", input.OAuthToken)
	accessTokenRequestQuery.Set("oauth_verifier", input.OAuthVerifier)
	accessTokenRequestURL.RawQuery = accessTokenRequestQuery.Encode()

	// Send a request for access token
	req, err := http.NewRequestWithContext(
		context.TODO(),
		"POST",
		accessTokenRequestURL.String(),
		nil,
	)
	if err != nil {
		return
	}

	res, err := s.HTTPClient.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		err = &errors.UnexpectedError{Message: string(body)}
		return
	}

	// Return oauth token and oauth token secret
	accessTokenResponseQuery, err := url.ParseQuery(string(body))
	if err != nil {
		return
	}

	output = port.GetAccessTokenOutput{
		OAuthToken:       accessTokenResponseQuery.Get("oauth_token"),
		OAuthTokenSecret: accessTokenResponseQuery.Get("oauth_token_secret"),
	}
	return
}
