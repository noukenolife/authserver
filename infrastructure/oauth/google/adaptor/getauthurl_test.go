package adaptor

import (
	"os"
	"testing"

	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/noukenolife/authserver/helper"
	"github.com/noukenolife/authserver/infrastructure/oauth/google/factory"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	helper.LoadDotEnv()
	os.Exit(m.Run())
}

func TestGoogleAuthURL(t *testing.T) {
	t.Run("Should get a google auth url successfully", func(t *testing.T) {
		adaptor := GoogleGetAuthURL{
			OAuthConfigFactory: factory.GoogleOAuthConfigFactory{},
		}
		output, _ := adaptor.Invoke(port.GetAuthURLInput{
			Scopes: []string{"email"},
		})
		assert.Contains(t, output.URL, "https://accounts.google.com/o/oauth2/auth")
		assert.Contains(t, output.URL, "scope=email")
	})
}
