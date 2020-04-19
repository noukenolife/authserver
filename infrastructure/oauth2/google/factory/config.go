package factory

import (
	"os"

	infoauth2 "github.com/noukenolife/authserver/infrastructure/oauth2"
	"github.com/noukenolife/authserver/infrastructure/oauth2/errors"
	"github.com/noukenolife/authserver/infrastructure/oauth2/factory"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuthConfigFactory struct {
	factory.OAuthConfigFactoryInterface
}

func (s GoogleOAuthConfigFactory) Create(scopes []string) (config infoauth2.OAuthConfigInterface, err error) {
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
