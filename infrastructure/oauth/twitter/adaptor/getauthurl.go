package adaptor

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/noukenolife/authserver/infrastructure/errors"
	oautherrs "github.com/noukenolife/authserver/infrastructure/oauth/errors"
)

type TwitterGetOAuthURL struct {
	port.GetAuthURL
}

func (s TwitterGetOAuthURL) Invoke(input port.GetAuthURLInput) (output port.GetAuthURLOutput, err error) {
	requestURL := "https://api.twitter.com/oauth/request_token"
	requestMethod := "POST"

	consumerKey, exists := os.LookupEnv("TWITTER_OAUTH_CONSUMER_KEY")
	if !exists {
		err = &oautherrs.InvalidConfigError{Message: "Twitter consumer key is not specified."}
		return
	}

	consumerSecret, exists := os.LookupEnv("TWITTER_OAUTH_CONSUMER_SECRET")
	if !exists {
		err = &oautherrs.InvalidConfigError{Message: "Twitter consumer secret is not specified."}
		return
	}

	redirectURL, exists := os.LookupEnv("TWITTER_OAUTH_REDIRECT_URL")
	if !exists {
		err = &oautherrs.InvalidConfigError{Message: "Twitter oauth redirect url is not specified."}
		return
	}

	nonce := base64.RawURLEncoding.EncodeToString([]byte(uuid.New().String()))
	timestamp := time.Now().Unix()
	oauthValueMap := map[string]string{
		"oauth_nonce":            nonce,
		"oauth_callback":         redirectURL,
		"oauth_signature_method": "HMAC-SHA1",
		"oauth_timestamp":        strconv.FormatInt(timestamp, 10),
		"oauth_consumer_key":     consumerKey,
		"oauth_version":          "1.0",
	}

	// Create signature base string
	u := url.URL{}
	q := u.Query()
	for k, v := range oauthValueMap {
		q.Add(k, v)
	}
	signatureBaseString := strings.Join([]string{requestMethod, url.QueryEscape(requestURL), url.QueryEscape(q.Encode())}, "&")
	h := hmac.New(sha1.New, []byte(url.QueryEscape(consumerSecret)+"&"))
	h.Write([]byte(signatureBaseString))
	oauthValueMap["oauth_signature"] = base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Create authorization header
	var authHeaderValues []string
	for k, v := range oauthValueMap {
		authHeaderValues = append(authHeaderValues, k+"=\""+url.QueryEscape(v)+"\"")
	}
	authHeader := "OAuth " + strings.Join(authHeaderValues, ", ")

	// Create request for token request
	req, err := http.NewRequestWithContext(
		context.TODO(),
		requestMethod,
		requestURL,
		nil,
	)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", authHeader)

	// Post the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		err = &errors.UnexpectedError{Message: string(body)}
		return
	}

	tokenQuery, err := url.ParseQuery(string(body))
	if err != nil {
		return
	}

	authURL, err := url.Parse("https://api.twitter.com/oauth/authorize")
	if err != nil {
		return
	}

	authQuery := authURL.Query()
	authQuery.Set("oauth_token", tokenQuery.Get("oauth_token"))
	authURL.RawQuery = authQuery.Encode()

	output = port.GetAuthURLOutput{
		URL: authURL.String(),
	}
	return
}
