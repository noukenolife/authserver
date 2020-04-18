package google

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noukenolife/authserver/application/oauth2/service"
)

type GetGoogleAuthURL struct {
	GetAuthURL service.GetAuthURLInterface
}

func (s GetGoogleAuthURL) Invoke(c *gin.Context) {
	var input service.GetAuthURLInput
	err := c.BindJSON(&input)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	output, err := s.GetAuthURL.Invoke(input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, output)
}
