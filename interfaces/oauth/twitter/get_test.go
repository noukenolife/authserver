package twitter

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/noukenolife/authserver/application/oauth/service"
	"github.com/stretchr/testify/mock"
)

type MockGetAuthURL struct {
	service.GetAuthURLInterface
	mock.Mock
}

func (s *MockGetAuthURL) Invoke(input service.GetAuthURLInput) (output service.GetAuthURLOutput, err error) {
	r := s.Called(input)
	if o, ok := r.Get(0).(service.GetAuthURLOutput); ok {
		output = o
	} else {
		err = r.Error(1)
	}
	return
}

func TestGetTwitterAuthURL(t *testing.T) {
	t.Run("should return an auth url successfully", func(t *testing.T) {
		input := service.GetAuthURLInput{}
		output := service.GetAuthURLOutput{
			URL: "http://example.com/oauth",
		}

		mockGetAuthURL := new(MockGetAuthURL)
		mockGetAuthURL.On(
			"Invoke",
			input,
		).Return(output, nil)

		getGoogleAuthURL := GetTwitterAuthURL{
			GetAuthURL: mockGetAuthURL,
		}

		r := gin.Default()
		r.GET("/oauth/twitter", getGoogleAuthURL.Invoke)

		w := httptest.NewRecorder()
		reqbody, _ := json.Marshal(input)
		req, _ := http.NewRequest("GET", "/oauth/twitter", strings.NewReader(string(reqbody)))
		r.ServeHTTP(w, req)

		var response service.GetAuthURLOutput
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, 200, w.Code)
		assert.Equal(t, output, response)
	})
	t.Run("should respond 400 when request body is invalid", func(t *testing.T) {
		getGoogleAuthURL := GetTwitterAuthURL{}

		r := gin.Default()
		r.GET("/oauth/twitter", getGoogleAuthURL.Invoke)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/oauth/twitter", strings.NewReader("INVALID_REQUEST"))
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
	t.Run("should respond 500 when failed to get an auth url", func(t *testing.T) {
		input := service.GetAuthURLInput{}

		mockGetAuthURL := new(MockGetAuthURL)
		mockGetAuthURL.On(
			"Invoke",
			input,
		).Return(nil, errors.New("Some Error"))

		getGoogleAuthURL := GetTwitterAuthURL{
			GetAuthURL: mockGetAuthURL,
		}

		r := gin.Default()
		r.GET("/oauth/twitter", getGoogleAuthURL.Invoke)

		w := httptest.NewRecorder()
		reqbody, _ := json.Marshal(input)
		req, _ := http.NewRequest("GET", "/oauth/twitter", strings.NewReader(string(reqbody)))
		r.ServeHTTP(w, req)

		assert.Equal(t, 500, w.Code)
	})
}
