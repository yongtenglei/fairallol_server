package router

import (
	"github.com/gin-gonic/gin"
	"rey.com/fairallol/logic"
	"rey.com/fairallol/middleware"
)

func SetUp() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())

	// saveworld = ef1
	router.POST("/saveworld", logic.SaveWorld)
	router.POST("/ef1", logic.SaveWorld)

	return router
}
