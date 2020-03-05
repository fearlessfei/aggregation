package http

import (
	"github.com/gin-gonic/gin"

	"solo/solo-go-common/net/http/middleware"
)

// InitRouter init router
func InitRouter() (r *gin.Engine) {
	// Creates a router without any middleware by default
	r = gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(middleware.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Set serviceInstance
	r.Use(func(c *gin.Context) {
		c.Set("srv", srv)
	})


	// api v1
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/user", User)
	}

	return
}
