package adaptor

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/noukenolife/authserver/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAccessToken(t *testing.T) {
	t.Run("should get an access token successfully", func(t *testing.T) {
		expectedOutput := port.GetAccessTokenOutput{
			OAuthToken:       "6253282-eWudHldSbIaelX7swmsiHImEL4KinwaGloHANdrY",
			OAuthTokenSecret: "2EEfA6BG5ly3sR3XjE0IBSnlQu4ZrUzPiYTmrkVU",
		}

		mockHTTPClient := new(helper.MockHTTPClient)
		mockHTTPClient.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("oauth_token=6253282-eWudHldSbIaelX7swmsiHImEL4KinwaGloHANdrY&oauth_token_secret=2EEfA6BG5ly3sR3XjE0IBSnlQu4ZrUzPiYTmrkVU&user_id=6253282&screen_name=twitterapi")),
		}, nil)

		adaptor := TwitterGetAccessToken{
			HTTPClient: mockHTTPClient,
		}
		output, _ := adaptor.Invoke(port.GetAccessTokenInput{
			OAuthToken:    "qLBVyoAAAAAAx72QAAATZxQWU6P",
			OAuthVerifier: "ghLM8lYmAxDbaqL912RZSRjCCEXKDIzx",
		})

		assert.Equal(t, expectedOutput, output)
	})
	t.Run("should fail when failed to send http request", func(t *testing.T) {
		mockHTTPClient := new(helper.MockHTTPClient)
		mockHTTPClient.On("Do", mock.Anything).Return(nil, errors.New("HTTP Error"))

		adaptor := TwitterGetAccessToken{
			HTTPClient: mockHTTPClient,
		}
		_, err := adaptor.Invoke(port.GetAccessTokenInput{
			OAuthToken:    "qLBVyoAAAAAAx72QAAATZxQWU6P",
			OAuthVerifier: "ghLM8lYmAxDbaqL912RZSRjCCEXKDIzx",
		})
		assert.NotNil(t, err)
	})
	t.Run("should fail when response status is not 200", func(t *testing.T) {
		mockHTTPClient := new(helper.MockHTTPClient)
		mockHTTPClient.On("Do", mock.Anything).Return(&http.Response{
			StatusCode: http.StatusUnauthorized,
			Body:       ioutil.NopCloser(bytes.NewBufferString("FAILED")),
		}, nil)

		adaptor := TwitterGetAccessToken{
			HTTPClient: mockHTTPClient,
		}
		_, err := adaptor.Invoke(port.GetAccessTokenInput{
			OAuthToken:    "qLBVyoAAAAAAx72QAAATZxQWU6P",
			OAuthVerifier: "ghLM8lYmAxDbaqL912RZSRjCCEXKDIzx",
		})
		assert.NotNil(t, err)
	})
}
