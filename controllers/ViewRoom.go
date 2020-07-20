package controllers

import (
	"APINoiTu/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// @Title ViewRoom
// @Description Get roomId
// @Param	roomId		path 	string	true	"roomId"
// @Success 200 {object} responses.BoolResponse
// @router /:roomId [get]
func (c *RoomController) GETRoomId(roomId string) {
	fmt.Println("[GETRoomId]", roomId)
	c.Data["roomId"] = roomId
	c.TplName = "index.html"
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @Param	roomId		path 	string	true	"roomId"
// @router /ws/:roomId [get]
func (c *RoomController) WS(roomId string) {
	log.Println("[WS] = ", roomId)
	// ws, err := upGrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	// if err != nil {
	// 	fmt.Println("[Channel] ", ws, err)
	// 	return
	// }
	// defer ws.Close()

	models.ServeWs(c.Ctx.ResponseWriter, c.Ctx.Request, roomId)
}
