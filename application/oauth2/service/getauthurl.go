package service

import (
	"github.com/noukenolife/authserver/application/errors"
	"github.com/noukenolife/authserver/application/oauth2/port"
)

type GetAuthURLInput struct {
	Scopes []string `json:"scopes" binding:"required"`
}

type GetAuthURLOutput struct {
	URL string `json:"auth_url"`
}

type GetAuthURLInterface interface {
	Invoke(input GetAuthURLInput) (GetAuthURLOutput, error)
}

type GetAuthURL struct {
	GetAuthURLInterface
	GetAuthURL port.GetAuthURL
}

func (s GetAuthURL) Invoke(input GetAuthURLInput) (output GetAuthURLOutput, err error) {
	pOutput, pErr := s.GetAuthURL.Invoke(port.GetAuthURLInput{
		Scopes: input.Scopes,
	})
	if pErr != nil {
		err = &errors.UnexpectedError{Message: "Failed to get auth url.", Cause: err}
		return
	}
	output = GetAuthURLOutput{
		URL: pOutput.URL,
	}
	return
}
