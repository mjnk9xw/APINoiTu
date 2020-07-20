// @APIVersion 1.0.0
// @Title API Nối từ
// @Description API by minh nguyễn dev
package routers

import (
	"APINoiTu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/noitu",
			beego.NSInclude(
				&controllers.Controller{},
			),
		),
	)
	// ns2 := beego.NewNamespace("/room",
	// 	beego.NSNamespace("/room",
	// 		beego.NSInclude(
	// 			&controllers.RoomController{},
	// 		),
	// 	),
	// )
	// beego.Router("/", &controllers.HomeController{})
	ns2 := beego.NewNamespace("/room", beego.NSInclude(&controllers.RoomController{}))
	ns3 := beego.NewNamespace("/home", beego.NSInclude(&controllers.HomeController{}))
	beego.AddNamespace(ns)
	beego.AddNamespace(ns2)
	beego.AddNamespace(ns3)

}
