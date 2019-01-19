# Gin OAuth 2.0 Server

> Using Gin framework implementation OAuth 2.0 services

[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![GoDoc][GoDoc-Image]][GoDoc-Url]

## Quick Start

### Download and install

``` bash
$ go get -u github.com/go-oauth2/gin-server
```

### Create file `server.go`

``` go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {
	manager := manage.NewDefaultManager()

	// token store
	manager.MustTokenStorage(store.NewFileTokenStore("data.db"))

	// client store
	clientStore := store.NewClientStore()
	clientStore.Set("000000", &models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})
	manager.MapClientStorage(clientStore)

	// Initialize the oauth2 service
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	g := gin.Default()

	auth := g.Group("/oauth2")
	{
		auth.GET("/token", ginserver.HandleTokenRequest)
	}

	api := g.Group("/api")
	{
		api.Use(ginserver.HandleTokenVerify())
		api.GET("/test", func(c *gin.Context) {
			ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
			if exists {
				c.JSON(http.StatusOK, ti)
				return
			}
			c.String(http.StatusOK, "not found")
		})
	}

	g.Run(":9096")
}
```

### Build and run

``` bash
$ go build server.go
$ ./server
```

### Open in your web browser

#### The token information

```
http://localhost:9096/oauth2/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read
```

``` json
{
    "access_token": "AJPNSQO2PCITABYX0RFLWG",
    "expires_in": 7200,
    "scope": "read",
    "token_type": "Bearer"
}
```

#### The authentication token

```
http://localhost:9096/api/test?access_token=AJPNSQO2PCITABYX0RFLWG
```

``` json
{
    "ClientID": "000000",
    "UserID": "",
    "RedirectURI": "",
    "Scope": "read",
    "Code": "",
    "CodeCreateAt": "0001-01-01T00:00:00Z",
    "CodeExpiresIn": 0,
    "Access": "AJPNSQO2PCITABYX0RFLWG",
    "AccessCreateAt": "2016-11-29T09:00:52.617250916+08:00",
    "AccessExpiresIn": 7200000000000,
    "Refresh": "",
    "RefreshCreateAt": "0001-01-01T00:00:00Z",
    "RefreshExpiresIn": 0
}
```

## MIT License

```
Copyright (c) 2016 Lyric
```

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/go-oauth2/gin-server
[ReportCard-Image]: https://goreportcard.com/badge/github.com/go-oauth2/gin-server
[GoDoc-Url]: https://godoc.org/github.com/go-oauth2/gin-server
[GoDoc-Image]: https://godoc.org/github.com/go-oauth2/gin-server?status.svg
