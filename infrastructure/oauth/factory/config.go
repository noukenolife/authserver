package factory

import "github.com/noukenolife/authserver/infrastructure/oauth"

type OAuthConfigFactoryInterface interface {
	Create(scopes []string) (oauth.OAuthConfigInterface, error)
}
