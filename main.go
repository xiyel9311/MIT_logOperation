package main

import (
	// "logOperation/models/setting"
	"logOperation/routers"
	// "github.com/aWildProgrammer/fconf"
)

// const (
// 	settingPath = "./setting.ini"
// )

// var globalSetting *setting.Setting

// func getSetting() {
// 	c, err := fconf.NewFileConf(settingPath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	globalSetting = new(setting.Setting)
// 	globalSetting.ThisPort = c.String("Setting.this_port")
// 	globalSetting.MainIP = c.String("Setting.main_ip")
// }

func init() {
	// getSetting()
}

func main() {
	router := routers.Setup()

	router.Run(":8885")

	// routers.SetupOri()

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", app.Config.Server.HttpPort),
	// 	Handler:        router,
	// 	ReadTimeout:    setting.ServerSetting.ReadTimeout,
	// 	WriteTimeout:   setting.ServerSetting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// s.ListenAndServe()
}
