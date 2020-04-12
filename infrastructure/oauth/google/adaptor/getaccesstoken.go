package adaptor

import (
	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/noukenolife/authserver/infrastructure/oauth/factory"
	"golang.org/x/oauth2"
)

type GoogleGetAccessToken struct {
	OAuthConfigFactory factory.OAuthConfigFactoryInterface
}

func (s GoogleGetAccessToken) Invoke(input port.GetAccessTokenInput) (output port.GetAccessTokenOutput, err error) {
	config, err := s.OAuthConfigFactory.Create([]string{})
	if err != nil {
		return
	}

	token, err := config.Exchange(oauth2.NoContext, input.Code)
	if err != nil {
		return
	}

	output = port.GetAccessTokenOutput{
		AccessToken: token.AccessToken,
	}
	return
}
