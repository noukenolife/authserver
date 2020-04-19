package service

import (
	"errors"
	"testing"

	"github.com/noukenolife/authserver/application/oauth2/port"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGetAuthURL struct {
	mock.Mock
	GetAuthURLInterface
}

func (s *MockGetAuthURL) Invoke(input port.GetAuthURLInput) (output port.GetAuthURLOutput, err error) {
	r := s.Called(input)
	if o, ok := r.Get(0).(port.GetAuthURLOutput); ok {
		output = o
	} else {
		err = r.Error(1)
	}
	return
}

func TestGetAuthURL(t *testing.T) {
	t.Run("should get an auth url successfully", func(t *testing.T) {
		pInput := port.GetAuthURLInput{Scopes: []string{}}
		pOutput := port.GetAuthURLOutput{"http://example.com/authurl"}

		mockGetAuthURL := new(MockGetAuthURL)
		mockGetAuthURL.On("Invoke", pInput).Return(pOutput, nil)

		service := GetAuthURL{
			GetAuthURL: mockGetAuthURL,
		}
		output, _ := service.Invoke(GetAuthURLInput{Scopes: []string{}})
		assert.Equal(t, "http://example.com/authurl", output.URL)
	})
	t.Run("should fail when failed to get an auth url", func(t *testing.T) {
		pInput := port.GetAuthURLInput{Scopes: []string{}}

		mockGetAuthURL := new(MockGetAuthURL)
		mockGetAuthURL.On("Invoke", pInput).Return(nil, errors.New("Some Error"))

		service := GetAuthURL{
			GetAuthURL: mockGetAuthURL,
		}
		_, err := service.Invoke(GetAuthURLInput{Scopes: []string{}})
		assert.NotNil(t, err)
	})
}
