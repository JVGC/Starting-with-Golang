package main

import (
	"sales-tool-clone/src/main/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.Routes(router)
	router.Run("localhost:5000")
}
