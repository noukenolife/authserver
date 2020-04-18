package di

import (
	"github.com/noukenolife/authserver/application/oauth2/service"
	"github.com/noukenolife/authserver/infrastructure/oauth2/google/adaptor"
	"github.com/noukenolife/authserver/infrastructure/oauth2/google/factory"
	"github.com/noukenolife/authserver/interfaces"
	igoogle "github.com/noukenolife/authserver/interfaces/oauth2/google"
	"github.com/noukenolife/authserver/interfaces/oauth2/google/token"
)

type Container struct {
	Router interfaces.Router
}

func NewContainer() (container Container, err error) {
	if err != nil {
		return
	}

	googleOAuthConfigFactory := factory.GoogleOAuthConfigFactory{}

	container = Container{
		interfaces.Router{
			GetGoogleAuthURL: igoogle.GetGoogleAuthURL{
				GetAuthURL: service.GetAuthURL{
					GetAuthURL: adaptor.GoogleGetAuthURL{
						OAuthConfigFactory: googleOAuthConfigFactory,
					},
				},
			},
			GetGoogleAccessToken: token.GetGoogleAccessToken{
				GetAccessToken: service.GetAccessToken{
					GetAccessToken: adaptor.GoogleGetAccessToken{
						OAuthConfigFactory: googleOAuthConfigFactory,
					},
				},
			},
		},
	}
	return
}
