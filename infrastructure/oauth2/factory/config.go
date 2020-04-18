package factory

import "github.com/noukenolife/authserver/infrastructure/oauth2"

type OAuthConfigFactoryInterface interface {
	Create(scopes []string) (oauth2.OAuthConfigInterface, error)
}
