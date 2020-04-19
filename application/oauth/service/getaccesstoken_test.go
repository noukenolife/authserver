package service

import (
	"errors"
	"testing"

	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGetAccessToken struct {
	mock.Mock
	GetAccessTokenInterface
}

func (s *MockGetAccessToken) Invoke(input port.GetAccessTokenInput) (output port.GetAccessTokenOutput, err error) {
	r := s.Called(input)
	if o, ok := r.Get(0).(port.GetAccessTokenOutput); ok {
		output = o
	} else {
		err = r.Error(1)
	}
	return
}

func TestMockGetAccessToken(t *testing.T) {
	t.Run("should get an access token successfully", func(t *testing.T) {
		pInput := port.GetAccessTokenInput{}
		pOutput := port.GetAccessTokenOutput{
			OAuthToken:       "OAUTH_TOKEN",
			OAuthTokenSecret: "OAUTH_TOKEN_SECRET",
		}

		mockGetAccessToken := new(MockGetAccessToken)
		mockGetAccessToken.On("Invoke", pInput).Return(pOutput, nil)

		service := GetAccessToken{
			GetAccessToken: mockGetAccessToken,
		}
		output, _ := service.Invoke(GetAccessTokenInput{})
		assert.Equal(t, "OAUTH_TOKEN", output.OAuthToken)
		assert.Equal(t, "OAUTH_TOKEN_SECRET", output.OAuthTokenSecret)
	})
	t.Run("should fail when failed to get an access token", func(t *testing.T) {
		pInput := port.GetAccessTokenInput{}

		mockGetAccessToken := new(MockGetAccessToken)
		mockGetAccessToken.On("Invoke", pInput).Return(nil, errors.New("Some Error"))

		service := GetAccessToken{
			GetAccessToken: mockGetAccessToken,
		}
		_, err := service.Invoke(GetAccessTokenInput{})
		assert.NotNil(t, err)
	})
}
