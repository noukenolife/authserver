package service

import (
	"testing"

	"github.com/noukenolife/authserver/application/oauth2/port"
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
		GetAuthURL: MockGetAuthUrl{Ouput: port.GetAuthURLOutput{"http://example.com/authurl"}},
	}
	output, _ := service.Invoke(GetAuthURLInput{Scopes: []string{}})
	assert.Equal(t, "http://example.com/authurl", output.URL)
}
