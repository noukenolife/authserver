package service

import (
	"testing"

	"github.com/noukenolife/authserver/application/oauth2/port"
	"github.com/stretchr/testify/assert"
)

type MockGetAccessToken struct {
	Output port.GetAccessTokenOutput
}

func (s MockGetAccessToken) Invoke(input port.GetAccessTokenInput) (output port.GetAccessTokenOutput, err error) {
	output = s.Output
	return
}

func TestMockGetAccessToken_Invoke(t *testing.T) {
	service := GetAccessToken{
		GetAccessToken: MockGetAccessToken{
			Output: port.GetAccessTokenOutput{AccessToken: "ACCESS_TOKEN"},
		},
	}
	output, _ := service.Invoke(GetAccessTokenInput{Code: "CODE"})
	assert.Equal(t, "ACCESS_TOKEN", output.AccessToken)
}
