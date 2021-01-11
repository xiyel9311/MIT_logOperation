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

	logController := &controllers.LogController{}

	r.PUT("/log_output", logController.UpdateOutput)
	r.GET("/log_output", logController.GetOutput)
	r.GET("/log_message", logController.LogMessage)

	return r
}
