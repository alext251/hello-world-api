package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// StartAPI starts up an API using the gin framework
// and takes in a boolean parameter to enable debug logging.
func StartAPI(debugMode bool) {
	r := setupRouter(debugMode)
	r.Run()
}

func setupRouter(debugMode bool) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	if debugMode {
		r.Use(logger())
	}

	r.GET("/", func(c *gin.Context) {
		if v, ok := c.Request.Header["Accept"]; ok && v[0] == "application/json" {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello, World",
			})
			return
		}
		c.Data(http.StatusOK, "text/html", []byte("<p>Hello, World</p>"))
	})

	r.POST("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", []byte("<p>Hello, World</p>"))
	})

	return r
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logMessage := fmt.Sprintf("- %s%s", c.Request.Host, c.Request.RequestURI)
		log.Println(logMessage)
	}
}
