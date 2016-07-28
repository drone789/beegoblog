package controllers

import . "jikeblog/models/class"

type UserController struct {
	//	beego.Controller
	BaseController
}

func (c *UserController) Profile() {

	id := c.Ctx.Input.Param(":id")
	u := &User{Id: id}
	u.ReadDB()

	c.Data["u"] = u

	as := Article{Author: u}.Gets()
	replys := Reply{Author: u}.Gets()

	c.Data["articles"] = as
	c.Data["replys"] = replys

	c.TplName = "user/profile.html"
}

func (c *UserController) PageJoin() {
	c.TplName = "user/join.html"
}

func (c *UserController) PageSetting() {
	c.CheckLogin()
	c.TplName = "user/setting.html"
}
