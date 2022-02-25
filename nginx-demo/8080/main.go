package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/", http.Dir("./public"))
	router.Run("0.0.0.0:8080")
}
