package main

import (
	"fmt"
	"memos/comm"
	"memos/dao/mysql"
	"memos/routes"
)

func main() {
	//var a []map[string]string
	//a = []map[string]string{
	//	{"apa2": "sss"},
	//	{"apa22": "s222ss"},
	//	{"apa222": "sss"},
	//}
	//comm.WriteJson(a, "./data/t.json")

	// 初始化读取配置
	if err := comm.InitViperCfg(); err != nil {
		return
	}
	cfgLoader := comm.CfgLoader

	// 初始化日志
	if err := comm.InitLogger(); err != nil {
		return
	}
	log := comm.Logger

	//- 初始化mysql连接
	if err := mysql.InitCfg(); err != nil {
		fmt.Println("init mysql failed!")
		return
	}
	defer mysql.Close()

	// 初始化路由
	r := routes.Init()
	log.Info().Msg("Router init...")

	env := cfgLoader.GetString("app.env") + ".url"
	if err := r.Run(cfgLoader.GetString(env)); err != nil {
		log.Error().Msg("App Run Error")
		return
	}
}
