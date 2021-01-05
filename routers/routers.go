package routers

import (
	"logOperation/controllers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	// r.LoadHTMLGlob("views/**/*")
	// r.LoadHTMLGlob("views/*")

	test := r.Group("/test")
	{
		testController := &controllers.TestController{}
		test.GET("", testController.Test)
	}

	log := r.Group("/log")
	{
		logController := &controllers.LogController{}

		log.PUT("/updateOutput", logController.UpdateOutput)
		log.GET("/getOutput", logController.GetOutput)
	}

	return r
}
