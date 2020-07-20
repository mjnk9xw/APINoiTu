package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["APINoiTu/controllers:Controller"] = append(beego.GlobalControllerRouter["APINoiTu/controllers:Controller"],
        beego.ControllerComments{
            Method: "Check",
            Router: `/check`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
