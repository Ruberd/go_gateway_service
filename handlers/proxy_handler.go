package handlers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// ProxyHandler wraps the http.HandlerFunc to make it compatible with Gin
func ProxyHandler(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL, err := url.Parse(target)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid service URL"})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		proxy.ErrorHandler = func(resp http.ResponseWriter, req *http.Request, err error) {
			log.Printf("Proxy error: %v", err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Service unavailable"})
		}

		r := c.Request

		r.Host = targetURL.Host

		proxy.ServeHTTP(c.Writer, r)
	}
}
