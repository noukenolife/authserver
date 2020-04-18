package service

import (
	"testing"

	"github.com/noukenolife/authserver/application/oauth/port"
	"github.com/stretchr/testify/assert"
)

type MockGetAuthUrl struct {
	Ouput port.GetAuthURLOutput
}

func (s MockGetAuthUrl) Invoke(input port.GetAuthURLInput) (output port.GetAuthURLOutput, err error) {
	output = s.Ouput
	return
}

func TestGetAuthUrl_Invoke(t *testing.T) {
	service := GetAuthURL{
		GetAuthURL: MockGetAuthUrl{Ouput: port.GetAuthURLOutput{"https://api.twitter.com/oauth/authorize?oauth_token=fake_oauth_token"}},
	}
	output, _ := service.Invoke(GetAuthURLInput{})
	assert.Equal(t, "https://api.twitter.com/oauth/authorize?oauth_token=fake_oauth_token", output.URL)
}
