package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

type OAuthConfigInterface interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}
