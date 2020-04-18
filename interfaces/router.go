package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/noukenolife/authserver/interfaces/oauth2/google"
	"github.com/noukenolife/authserver/interfaces/oauth2/google/token"
)

type Router struct {
	GetGoogleAuthURL     google.GetGoogleAuthURL
	GetGoogleAccessToken token.GetGoogleAccessToken
}

func (s Router) InitRoutes(r *gin.Engine) {
	g := r.Group("/api")
	g.GET("/oauth/google", s.GetGoogleAuthURL.Invoke)
	g.GET("/oauth/google/token", s.GetGoogleAccessToken.Invoke)
}
