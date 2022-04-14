package presentation

import "github.com/gin-gonic/gin"

type Controller interface{
	Handle(c *gin.Context)
}