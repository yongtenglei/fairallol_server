package router

import (
	"github.com/gin-gonic/gin"
	"rey.com/fairallol/logic"
	"rey.com/fairallol/middleware"
)

func SetUp() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())
	// ef1
	router.POST("/saveworld", logic.SaveWorld)

	//v1Group := router.Group("/v1/api/user", middlewares.JWTAuth())
	//{
	//// info
	//v1Group.GET("/info/:mobile", logic.InfoHandler)
	//// update
	//v1Group.PUT("/update", logic.UpdateHandler)
	//// delete
	//v1Group.DELETE("/delete/:mobile", logic.DeleteHandler)
	//// re password
	//v1Group.PUT("/repassword", logic.RePasswordHandler)

	//}
	return router
}
