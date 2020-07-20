package controllers

import (
	"fmt"
)

// @Title ViewRoom
// @Description Get roomId
// @Param	roomId		path 	string	true	"roomId"
// @Param	username		path 	string	true	"username"
// @Success 200 {object} responses.BoolResponse
// @router /:roomId/:username [get]
func (c *RoomController) GETRoomId(roomId, username string) {
	fmt.Println("[GETRoomId]", roomId)
	c.Data["roomId"] = roomId
	c.Data["username"] = username
	c.TplName = "chat.html"
}
