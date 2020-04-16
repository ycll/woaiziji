package controllers

import "woaiziji/utils"

type ExitController struct {
	BaseController
}

func (c *ExitController) Get()  {
	c.DelSession(utils.LoginUser)
	c.Redirect("/", 302)
}
