package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	initOAuth2()

	g := gin.Default()

	g.GET("/authorize", server.HandleAuthorizeRequest)
	g.GET("/token", server.HandleTokenRequest)
	api := g.Group("/api")
	{
		api.Use(server.TokenAuth(tokenAuthHandle))
		api.GET("/test", testHandle)
	}

	g.Run(":9096")
}

func initOAuth2() {
	manager := manage.NewDefaultManager()
	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	// client store
	manager.MapClientStorage(store.NewTestClientStore(&models.Client{
		ID:     "999999",
		Secret: "999999",
	}))

	// Initialize the oauth2 service
	server.InitServer(manager)
	server.SetAllowGetAccessRequest(true)
}

func tokenAuthHandle(c *gin.Context) (token string) {
	token = c.Query("access_token")
	return
}

func testHandle(c *gin.Context) {
	ti, _ := c.Get("Token")
	c.JSON(http.StatusOK, ti)
}
