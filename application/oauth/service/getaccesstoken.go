package service

import (
	"github.com/noukenolife/authserver/application/errors"
	"github.com/noukenolife/authserver/application/oauth/port"
)

type GetAccessTokenInput struct {
	OAuthToken    string `form:"oauth_token" binding:"required"`
	OAuthVerifier string `form:"oauth_verifier" binding:"required"`
}

type GetAccessTokenOutput struct {
	OAuthToken       string `json:"oauth_token"`
	OAuthTokenSecret string `json:"oauth_token_secret"`
}

type GetAccessTokenInterface interface {
	Invoke(input GetAccessTokenInput) (GetAccessTokenOutput, error)
}

type GetAccessToken struct {
	GetAccessTokenInterface
	GetAccessToken port.GetAccessToken
}

func (s GetAccessToken) Invoke(input GetAccessTokenInput) (output GetAccessTokenOutput, err error) {
	pOutput, pErr := s.GetAccessToken.Invoke(port.GetAccessTokenInput{
		OAuthToken:    input.OAuthToken,
		OAuthVerifier: input.OAuthVerifier,
	})
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
