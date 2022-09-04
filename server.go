package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

type IPAddress struct {
	IP string `json:"ip"`
}

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", func(c *gin.Context) {
		forward := ReadUserIP(c)

		// Call the HTML method of the Context to render a template
		c.HTML(

			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title":    "La HOMEPAGE",
				"remoteIP": forward,
			},
		)

	})
	router.GET("/api/my-ip/", func(c *gin.Context) {
		forward := ReadUserIP(c)

		var ip IPAddress
		ip.IP = forward
		// Call the HTML method of the Context to render a template
		c.JSON(http.StatusOK, ip)
	})

	// Start serving the application
	router.Run()

}

func ReadUserIP(r *gin.Context) string {
	IPAddress := r.Request.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Request.Header.Get("X-Real-Ip")
	}
	if IPAddress == "" {
		IPAddress = r.Request.RemoteAddr
	}
	ip := strings.Split(IPAddress, ":")
	return ip[0]
}
