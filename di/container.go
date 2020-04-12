package di

import (
	"github.com/noukenolife/authserver/application/oauth/service"
	"github.com/noukenolife/authserver/infrastructure/oauth/google/adaptor"
	"github.com/noukenolife/authserver/infrastructure/oauth/google/factory"
	"github.com/noukenolife/authserver/interfaces"
	igoogle "github.com/noukenolife/authserver/interfaces/oauth/google"
	"github.com/noukenolife/authserver/interfaces/oauth/google/token"
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
