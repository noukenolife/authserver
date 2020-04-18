package adaptor

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/joho/godotenv"
	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/noukenolife/authserver/helper"
)

func TestTwitterGetOAuthURL(t *testing.T) {
	godotenv.Load(helper.ProjectRootPath() + "/.env")
	adaptor := TwitterGetOAuthURL{}

	output, _ := adaptor.Invoke(port.GetAuthURLInput{})

	assert.MatchRegex(t, output.URL, "https:\\/\\/api.twitter.com\\/oauth\\/authorize\\?oauth_token=.+?$")
}
