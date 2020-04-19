package token

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noukenolife/authserver/application/oauth/service"
)

type GetTwitterAccessToken struct {
	GetAccessToken service.GetAccessTokenInterface
}

func (s GetTwitterAccessToken) Invoke(c *gin.Context) {
	var input service.GetAccessTokenInput
	err := c.BindQuery(&input)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	output, err := s.GetAccessToken.Invoke(input)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, output)
}
