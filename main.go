package main

import (
	_ "APINoiTu/routers"
	"log"
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

	log.Println("Env $PORT :", os.Getenv("PORT"))
	if os.Getenv("PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
			log.Fatal("$PORT must be set")
		}
		log.Println("port : ", port)
		beego.BConfig.Listen.HTTPPort = port
		beego.BConfig.Listen.HTTPSPort = port
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
