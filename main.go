package main

import (
	_ "APINoiTu/routers"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

func main() {

	// beego.BConfig.WebConfig.DirectoryIndex = true
	// beego.BConfig.WebConfig.StaticDir["/v1/noitu/swagger"] = "swagger"

	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type", "Content-Type", ""},
	// 	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
	// 	AllowCredentials: true,
	// }))
	// beego.Run()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.HttpPort = port
	}

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/v1/noitu/swagger"] = "swagger"

	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowAllOrigins:  true,
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type", "Content-Type", ""},
	// 	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type"},
	// 	AllowCredentials: true,
	// }))
	beego.Run()

}
