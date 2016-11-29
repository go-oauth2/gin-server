package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	aserver "gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	server.InitServer(manager)
	server.SetAllowGetAccessRequest(true)
	server.SetClientInfoHandler(aserver.ClientFormHandler)

	g := gin.Default()

	auth := g.Group("/oauth2")
	{
		auth.GET("/token", server.HandleTokenRequest)
	}

	api := g.Group("/api")
	{
		api.Use(server.HandleTokenVerify())
		api.GET("/test", func(c *gin.Context) {
			ti, exists := c.Get("AccessToken")
			if exists {
				c.JSON(http.StatusOK, ti)
				return
			}
			c.String(http.StatusOK, "not found")
		})
	}

	g.Run(":9096")
}
