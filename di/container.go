package di

import (
	oauth_service "github.com/noukenolife/authserver/application/oauth/service"
	oauth2_service "github.com/noukenolife/authserver/application/oauth2/service"
	twitter_adaptor "github.com/noukenolife/authserver/infrastructure/oauth/twitter/adaptor"
	google_adaptor "github.com/noukenolife/authserver/infrastructure/oauth2/google/adaptor"
	"github.com/noukenolife/authserver/infrastructure/oauth2/google/factory"
	"github.com/noukenolife/authserver/interfaces"
	interface_twitter "github.com/noukenolife/authserver/interfaces/oauth/twitter"
	interface_google "github.com/noukenolife/authserver/interfaces/oauth2/google"
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
			GetGoogleAuthURL: interface_google.GetGoogleAuthURL{
				GetAuthURL: oauth2_service.GetAuthURL{
					GetAuthURL: google_adaptor.GoogleGetAuthURL{
						OAuthConfigFactory: googleOAuthConfigFactory,
					},
				},
			},
			GetGoogleAccessToken: token.GetGoogleAccessToken{
				GetAccessToken: oauth2_service.GetAccessToken{
					GetAccessToken: google_adaptor.GoogleGetAccessToken{
						OAuthConfigFactory: googleOAuthConfigFactory,
					},
				},
			},
			GetTwitterAuthURL: interface_twitter.GetTwitterAuthURL{
				GetAuthURL: oauth_service.GetAuthURL{
					GetAuthURL: twitter_adaptor.TwitterGetOAuthURL{},
				},
			},
		},
	}
	return
}
