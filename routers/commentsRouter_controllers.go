package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["APINoiTu/controllers:Controller"] = append(beego.GlobalControllerRouter["APINoiTu/controllers:Controller"],
        beego.ControllerComments{
            Method: "Check",
            Router: "/check",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APINoiTu/controllers:HomeController"] = append(beego.GlobalControllerRouter["APINoiTu/controllers:HomeController"],
        beego.ControllerComments{
            Method: "ViewHome",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APINoiTu/controllers:RoomController"] = append(beego.GlobalControllerRouter["APINoiTu/controllers:RoomController"],
        beego.ControllerComments{
            Method: "GETRoomId",
            Router: "/:roomId/:username",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("roomId", param.IsRequired, param.InPath),
				param.New("username", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APINoiTu/controllers:RoomController"] = append(beego.GlobalControllerRouter["APINoiTu/controllers:RoomController"],
        beego.ControllerComments{
            Method: "WS",
            Router: "/ws/:roomId/:username",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("roomId", param.IsRequired, param.InPath),
				param.New("username", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

}
