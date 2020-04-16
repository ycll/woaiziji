package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"woaiziji/utils"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	LoginUser interface{}
}

func (c *BaseController) Prepare() {
	loginUser := c.GetSession(utils.LoginUser)
	fmt.Println("loginUser-->", loginUser)
	if loginUser != nil {
		c.IsLogin = true
		c.LoginUser = loginUser
	} else {
		c.IsLogin = false
	}
	c.Data["IsLogin"] = c.IsLogin
}