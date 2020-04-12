package service

import (
	"github.com/noukenolife/authserver/application/errors"
	"github.com/noukenolife/authserver/application/oauth/port"
)

type GetAccessTokenInput struct {
	State string `form:"state" binding:"required"`
	Code  string `form:"code" binding:"required"`
}

type GetAccessTokenOutput struct {
	AccessToken string
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
		Code: input.Code,
	})
	if pErr != nil {
		err = &errors.UnexpectedError{
			Message: "Failed to get an access token",
			Cause:   pErr,
		}
		return
	}
	output = GetAccessTokenOutput{
		AccessToken: pOutput.AccessToken,
	}
	return
}
