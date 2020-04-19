package twitter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noukenolife/authserver/application/oauth/service"
)

type GetTwitterAuthURL struct {
	GetAuthURL service.GetAuthURLInterface
}

func (s GetTwitterAuthURL) Invoke(c *gin.Context) {
	output, err := s.GetAuthURL.Invoke(service.GetAuthURLInput{})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, output)
}
