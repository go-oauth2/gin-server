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
	"gopkg.in/oauth2.v3/store"
)

func main() {
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	manager.MapClientStorage(store.NewTestClientStore())

	// Initialize the oauth2 service
	server.InitServer(manager)
	server.SetAllowGetAccessRequest(true)

	g := gin.Default()
	g.GET("/token", server.HandleTokenRequest)
	api := g.Group("/api")
	{
		api.Use(server.TokenAuth(func(c *gin.Context) string {
			return c.Query("access_token")
		}))
		api.GET("/test", func(c *gin.Context) {
			ti, _ := c.Get("Token")
			c.JSON(http.StatusOK, ti)
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
http://localhost:9096/token?grant_type=client_credentials&client_id=1&client_secret=11&scope=read
```

``` json
{
    "access_token": "ZF1M7NKDNWUUX2TCDIMZZG",
    "expires_in": 7200,
    "scope": "read",
    "token_type": "Bearer"
}
```

#### The authentication token

```
http://localhost:9096/api/test?access_token=ZF1M7NKDNWUUX2TCDIMZZG
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