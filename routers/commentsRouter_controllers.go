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

    beego.GlobalControllerRouter["APINoiTu/controllers:RoomController"] = append(beego.GlobalControllerRouter["APINoiTu/controllers:RoomController"],
        beego.ControllerComments{
            Method: "GETRoomId",
            Router: "/:roomId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("roomId", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["APINoiTu/controllers:RoomController"] = append(beego.GlobalControllerRouter["APINoiTu/controllers:RoomController"],
        beego.ControllerComments{
            Method: "WS",
            Router: "/ws/:roomId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("roomId", param.IsRequired, param.InPath),
			),
            Filters: nil,
            Params: nil})

}
