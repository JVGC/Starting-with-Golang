package config

import (
	"sales-tool-clone/src/main/factories"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	addUserController := factories.MakeAddUserController()
	router.POST("/user", addUserController.Handle)
}
