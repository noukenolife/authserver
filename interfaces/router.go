package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/noukenolife/authserver/interfaces/oauth/twitter"
	twitter_token "github.com/noukenolife/authserver/interfaces/oauth/twitter/token"
	"github.com/noukenolife/authserver/interfaces/oauth2/google"
	google_token "github.com/noukenolife/authserver/interfaces/oauth2/google/token"
)

type Router struct {
	GetGoogleAuthURL      google.GetGoogleAuthURL
	GetGoogleAccessToken  google_token.GetGoogleAccessToken
	GetTwitterAuthURL     twitter.GetTwitterAuthURL
	GetTwitterAccessToken twitter_token.GetTwitterAccessToken
}

func (s Router) InitRoutes(r *gin.Engine) {
	g := r.Group("/api")
	g.GET("/oauth/google", s.GetGoogleAuthURL.Invoke)
	g.GET("/oauth/google/token", s.GetGoogleAccessToken.Invoke)

	g.GET("/oauth/twitter", s.GetTwitterAuthURL.Invoke)
	g.GET("/oauth/twitter/token", s.GetTwitterAccessToken.Invoke)
}
