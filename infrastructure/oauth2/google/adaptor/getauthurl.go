package adaptor

import (
	"github.com/google/uuid"
	"github.com/noukenolife/authserver/application/oauth2/port"
	"github.com/noukenolife/authserver/infrastructure/oauth2/factory"
	"golang.org/x/oauth2"
)

type GoogleGetAuthURL struct {
	OAuthConfigFactory factory.OAuthConfigFactoryInterface
}

func (s GoogleGetAuthURL) Invoke(input port.GetAuthURLInput) (output port.GetAuthURLOutput, err error) {
	state := uuid.New().String()

	config, err := s.OAuthConfigFactory.Create(input.Scopes)
	if err != nil {
		return
	}

	authURL := config.AuthCodeURL(state, oauth2.ApprovalForce)
	return port.GetAuthURLOutput{
		URL: authURL,
	}, nil
}
