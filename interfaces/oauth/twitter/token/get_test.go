package token

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/noukenolife/authserver/application/oauth/service"
	"github.com/stretchr/testify/mock"
)

type MockGetAccessToken struct {
	service.GetAccessTokenInterface
	mock.Mock
}

func (s *MockGetAccessToken) Invoke(input service.GetAccessTokenInput) (output service.GetAccessTokenOutput, err error) {
	r := s.Called(input)
	if o, ok := r.Get(0).(service.GetAccessTokenOutput); ok {
		output = o
	} else {
		err = r.Error(1)
	}
	return
}

func TestGetGoogleAccessToken(t *testing.T) {
	t.Run("should return an access token successfully", func(t *testing.T) {
		input := service.GetAccessTokenInput{
			OAuthToken:    "OAUTH_TOKEN",
			OAuthVerifier: "OAUTH_VERIFIER",
		}
		output := service.GetAccessTokenOutput{
			OAuthToken:       "OAUTH_TOKEN",
			OAuthTokenSecret: "OAUTH_TOKEN_SECRET",
		}

		mockGetAccessToken := new(MockGetAccessToken)
		mockGetAccessToken.On(
			"Invoke",
			input,
		).Return(output, nil)

		getTwitterAccessToken := GetTwitterAccessToken{
			GetAccessToken: mockGetAccessToken,
		}

		r := gin.Default()
		r.GET("/oauth/twitter/token", getTwitterAccessToken.Invoke)

		w := httptest.NewRecorder()
		url, _ := url.Parse("/oauth/twitter/token")
		q := url.Query()
		q.Set("oauth_token", "OAUTH_TOKEN")
		q.Set("oauth_verifier", "OAUTH_VERIFIER")
		url.RawQuery = q.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		r.ServeHTTP(w, req)

		var response service.GetAccessTokenOutput
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, output, response)
	})
	t.Run("should respond 400 when query string is invalid", func(t *testing.T) {
		getTwitterAccessToken := GetTwitterAccessToken{}

		r := gin.Default()
		r.GET("/oauth/twitter/token", getTwitterAccessToken.Invoke)

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", "/oauth/twitter/token", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
	t.Run("should respond 500 when failed to get an access token", func(t *testing.T) {
		input := service.GetAccessTokenInput{
			OAuthToken:    "OAUTH_TOKEN",
			OAuthVerifier: "OAUTH_VERIFIER",
		}

		mockGetAccessToken := new(MockGetAccessToken)
		mockGetAccessToken.On(
			"Invoke",
			input,
		).Return(nil, errors.New("Some Error"))

		getTwitterAccessToken := GetTwitterAccessToken{
			GetAccessToken: mockGetAccessToken,
		}

		r := gin.Default()
		r.GET("/oauth/twitter/token", getTwitterAccessToken.Invoke)

		w := httptest.NewRecorder()
		url, _ := url.Parse("/oauth/twitter/token")
		q := url.Query()
		q.Set("oauth_token", "OAUTH_TOKEN")
		q.Set("oauth_verifier", "OAUTH_VERIFIER")
		url.RawQuery = q.Encode()

		req, _ := http.NewRequest("GET", url.String(), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}
