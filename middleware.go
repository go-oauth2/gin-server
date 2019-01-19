package ginserver

import (
	"github.com/gin-gonic/gin"
)

type (
	// ErrorHandleFunc error handling function
	ErrorHandleFunc func(*gin.Context, error)
	// Config defines the config for Session middleware
	Config struct {
		// error handling when starting the session
		ErrorHandleFunc ErrorHandleFunc
		// keys stored in the context
		TokenKey string
		// defines a function to skip middleware.Returning true skips processing
		// the middleware.
		Skipper func(*gin.Context) bool
	}
)

var (
	// DefaultConfig is the default middleware config.
	DefaultConfig = Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			ctx.AbortWithError(500, err)
		},
		TokenKey: "github.com/go-oauth2/gin-server/access-token",
		Skipper: func(_ *gin.Context) bool {
			return false
		},
	}
)

// HandleTokenVerify Verify the access token of the middleware
func HandleTokenVerify(config ...Config) gin.HandlerFunc {
	cfg := DefaultConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.ErrorHandleFunc == nil {
		cfg.ErrorHandleFunc = DefaultConfig.ErrorHandleFunc
	}

	tokenKey := cfg.TokenKey
	if tokenKey == "" {
		tokenKey = DefaultConfig.TokenKey
	}

	return func(c *gin.Context) {
		if cfg.Skipper != nil && cfg.Skipper(c) {
			c.Next()
			return
		}
		ti, err := gServer.ValidationBearerToken(c.Request)
		if err != nil {
			cfg.ErrorHandleFunc(c, err)
			return
		}

		c.Set(tokenKey, ti)
		c.Next()
	}
}
