package factory

import (
	"os"
	"testing"

	"github.com/noukenolife/authserver/infrastructure/oauth2/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func TestNewGoogleOAuthConfig(t *testing.T) {
	t.Run("should create an oauth config object successfully", func(t *testing.T) {
		factory := GoogleOAuthConfigFactory{}

		os.Setenv("GOOGLE_OAUTH_CLIENT_ID", "CLIENT_ID")
		os.Setenv("GOOGLE_OAUTH_CLIENT_SECRET", "CLIENT_SECRET")
		os.Setenv("GOOGLE_OAUTH_REDIRECT_URL", "http://127.0.0.1:3000")

		scopes := []string{"email"}
		iconfig, _ := factory.Create(scopes)
		config := iconfig.(*oauth2.Config)

		assert.Equal(t, "CLIENT_ID", config.ClientID)
		assert.Equal(t, "CLIENT_SECRET", config.ClientSecret)
		assert.Equal(t, "http://127.0.0.1:3000", config.RedirectURL)
		assert.Equal(t, google.Endpoint, config.Endpoint)
		assert.Equal(t, scopes, config.Scopes)

		os.Clearenv()
	})
	t.Run("should error out when there's missing env vars", func(t *testing.T) {
		factory := GoogleOAuthConfigFactory{}
		scopes := []string{"email"}

		var err error
		_, err = factory.Create(scopes)
		assert.IsType(t, &errors.InvalidConfigError{}, err)

		os.Setenv("GOOGLE_OAUTH_CLIENT_ID", "CLIENT_ID")
		_, err = factory.Create(scopes)
		assert.IsType(t, &errors.InvalidConfigError{}, err)

		os.Setenv("GOOGLE_OAUTH_CLIENT_SECRET", "CLIENT_SECRET")
		_, err = factory.Create(scopes)
		assert.IsType(t, &errors.InvalidConfigError{}, err)

		os.Clearenv()
	})
}
