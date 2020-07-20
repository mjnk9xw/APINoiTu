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

	beego.AddNamespace(ns)
}
