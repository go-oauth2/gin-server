package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/server"
)

var (
	gServer *server.Server
)

// InitServer Initialize the service
func InitServer(manager oauth2.Manager) {
	if err := manager.CheckInterface(); err != nil {
		panic(err)
	}
	gServer = server.NewDefaultServer(manager)
}

// HandleAuthorizeRequest the authorization request handling
func HandleAuthorizeRequest(c *gin.Context) {
	err := gServer.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

// HandleTokenRequest token request handling
func HandleTokenRequest(c *gin.Context) {
	err := gServer.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

// HandleTokenVerify Verify the access token of the middleware
func HandleTokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {

		ti, err := gServer.ValidationBearerToken(c.Request)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("AccessToken", ti)
		c.Next()
	}
}
