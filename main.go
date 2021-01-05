package main

import (
	"logOperation/routers"
)

func init() {

}

func main() {
	router := routers.Setup()

	router.Run(":8088")

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", app.Config.Server.HttpPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ServerSetting.ReadTimeout,
	// 	WriteTimeout:   setting.ServerSetting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// s.ListenAndServe()
}
