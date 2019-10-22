package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xautjzd/ginex/version"
)

// Version get version to gin example
func Version(c *gin.Context) {
	c.String(http.StatusOK, version.Version())
}
