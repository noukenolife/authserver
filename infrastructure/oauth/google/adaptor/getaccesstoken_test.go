package adaptor

import (
	"context"
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/noukenolife/authserver/infrastructure/oauth"
	"github.com/noukenolife/authserver/infrastructure/oauth/factory"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

type MockOAuthConfigFactory struct {
	factory.OAuthConfigFactoryInterface
	MockConfig *MockConfig
}

func (s MockOAuthConfigFactory) Create(scopes []string) (oauth.OAuthConfigInterface, error) {
	return s.MockConfig, nil
}

type MockConfig struct {
	oauth.OAuthConfigInterface
	mock.Mock
}

func (m *MockConfig) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (token *oauth2.Token, err error) {
	r := m.Called(ctx, code, opts)
	t := r.Get(0)
	if t != nil {
		token = t.(*oauth2.Token)
	}
	err = r.Error(1)
	return
}

func TestGoogleGetAccessToken(t *testing.T) {
	t.Run("should get an access token successfully", func(t *testing.T) {
		mockConfig := new(MockConfig)
		mockConfig.On(
			"Exchange",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(&oauth2.Token{AccessToken: "ACCESS_TOKEN"}, nil)

		adaptor := GoogleGetAccessToken{
			OAuthConfigFactory: MockOAuthConfigFactory{
				MockConfig: mockConfig,
			},
		}
		output, _ := adaptor.Invoke(port.GetAccessTokenInput{})
		assert.Equal(t, "ACCESS_TOKEN", output.AccessToken)
	})
	t.Run("should fail when an access token could not be obtained", func(t *testing.T) {
		expectedErr := errors.New("Some Error")
		mockConfig := new(MockConfig)
		mockConfig.On(
			"Exchange",
			mock.Anything,
			mock.Anything,
			mock.Anything,
		).Return(nil, expectedErr)

		adaptor := GoogleGetAccessToken{
			OAuthConfigFactory: MockOAuthConfigFactory{
				MockConfig: mockConfig,
			},
		}
		_, err := adaptor.Invoke(port.GetAccessTokenInput{})
		assert.Equal(t, expectedErr, err)
	})
}
