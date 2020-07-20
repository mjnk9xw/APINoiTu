package controllers

import "log"

// @router / [get]
func (c *HomeController) ViewHome() {
	log.Println("[ViewHome]")
	// ip := c.Ctx.Input.IP()
	c.TplName = "home.html"
}
