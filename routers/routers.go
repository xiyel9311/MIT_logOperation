package routers

import (
	"logOperation/controllers"
	"logOperation/controllers/logControllers"
	"logOperation/controllers/rs485DataControllers"

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

	logController := &logControllers.LogController{}

	r.PUT("/log_output", logController.UpdateOutput)
	r.GET("/log_output", logController.Output)
	r.GET("/log_message", logController.LogMessage)

	ew := r.Group("/expansion_word")
	{
		rs485dataCtrl := &rs485DataControllers.RS485Controller{}
		ew.GET("com_:comID", rs485dataCtrl.ComData)
		ew.GET("com_:comID/ch_:chID", rs485dataCtrl.ChData)

		ew.GET("temprature_and_humidity", rs485dataCtrl.TempratureAndHumidity)
	}

	return r
}
