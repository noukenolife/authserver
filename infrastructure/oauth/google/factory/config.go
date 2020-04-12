package factory

import (
	"os"

	"github.com/noukenolife/authserver/infrastructure/oauth"
	"github.com/noukenolife/authserver/infrastructure/oauth/errors"
	"github.com/noukenolife/authserver/infrastructure/oauth/factory"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuthConfigFactory struct {
	factory.OAuthConfigFactoryInterface
}

func (s GoogleOAuthConfigFactory) Create(scopes []string) (config oauth.OAuthConfigInterface, err error) {
	ClientID, exists := os.LookupEnv("GOOGLE_OAUTH_CLIENT_ID")
	if !exists {
		err = &errors.InvalidConfigError{Message: "Google oauth client id is not specified."}
		return
	}
	ClientSecret, exists := os.LookupEnv("GOOGLE_OAUTH_CLIENT_SECRET")
	if !exists {
		err = &errors.InvalidConfigError{Message: "Google oauth client secret is not specified."}
		return
	}
	RedirectURL, exists := os.LookupEnv("GOOGLE_OAUTH_REDIRECT_URL")
	if !exists {
		err = &errors.InvalidConfigError{Message: "Google oauth redirect url is not specified."}
		return
	}

	config = &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  RedirectURL,
		Endpoint:     google.Endpoint,
		Scopes:       scopes,
	}
	return
}
