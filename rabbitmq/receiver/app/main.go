package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "haha")
}

func main() {
	router := gin.Default()
	router.GET("/", get)
	router.Run("0.0.0.0:8080")
}
