package router

import (
	"github.com/gin-gonic/gin"
	"rey.com/fairallol/logic"
	"rey.com/fairallol/middleware"
)

func SetUp() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())

	// Saveworld
	router.POST("/saveworld", logic.SaveWorld)

	// Ef1
	router.POST("/ef1", logic.SaveWorld)

	// AdjustedWinner
	router.POST("/adjustedwinner", logic.AdjustedWinner)

	// AdjustedWinner
	router.POST("/dividechoose", logic.DivideAndChoose)

	// RoundRobin
	router.POST("/roundrobin", logic.RoundRobin)

	return router
}
