package service

import (
	"github.com/noukenolife/authserver/application/errors"
	"github.com/noukenolife/authserver/application/oauth/port"
)

type GetAccessTokenInput struct {
	OAuthToken    string
	OAuthVerifier string
}

type GetAccessTokenOutput struct {
	OAuthToken       string
	OAuthTokenSecret string
}

type GetAccessTokenInterface interface {
	Invoke(input GetAccessTokenInput) (GetAccessTokenOutput, error)
}

type GetAccessToken struct {
	GetAccessTokenInterface
	GetAccessToken port.GetAccessToken
}

func (s GetAccessToken) Invoke(input GetAccessTokenInput) (output GetAccessTokenOutput, err error) {
	pOutput, pErr := s.GetAccessToken.Invoke(port.GetAccessTokenInput{})
	if pErr != nil {
		err = &errors.UnexpectedError{
			Message: "Failed to get an access token",
			Cause:   pErr,
		}
		return
	}
	output = GetAccessTokenOutput{
		OAuthToken:       pOutput.OAuthToken,
		OAuthTokenSecret: pOutput.OAuthTokenSecret,
	}
	return
}
