package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(c *gin.Context) {
	c.Next()
}

// ErrorHandler is a middleware to controller errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()
	// TODO: Type it in a better way
	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": c.Errors,
		})
	}
}
