package controllers

import "fmt"

// @Title ViewRoom
// @Description Get roomId
// @Param	roomId		query 	string	true	"roomId"
// @Param	username	query 	string	true	"username"
// @router / [get]
func (c *RoomController) ViewRoom() {
	var (
		roomId, username string
	)

	c.Ctx.Input.Bind(&roomId, "roomId")
	c.Ctx.Input.Bind(&username, "username")

	fmt.Println(roomId, username)
	c.Data["roomId"] = roomId
	c.Data["username"] = username
	c.TplName = "chat.html"
}
