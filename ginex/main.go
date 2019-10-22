package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xautjzd/ginex/endpoints"
)

func main() {
	r := gin.Default()
	r.GET("/version", endpoints.Version)

	r.Run()
}
