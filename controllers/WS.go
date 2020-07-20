package controllers

import (
	"APINoiTu/models"
	"log"
)

// @Param	roomId		path 	string	true	"roomId"
// @Param	username		path 	string	true	"username"
// @router /ws/:roomId/:username [get]
func (c *RoomController) WS(roomId string, username string) {
	log.Println("[WS] = ", roomId, username)
	ip := c.Ctx.Input.IP()
	models.ServeWs(c.Ctx.ResponseWriter, c.Ctx.Request, roomId, username, ip)

	c.ServeJSON()
}
