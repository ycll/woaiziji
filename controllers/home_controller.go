package controllers

import (
	"fmt"
	"woaiziji/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get()  {
		page, _ := c.GetInt("page")
		if page <= 0 {
			page = 1
		}
		var artList []models.Article
		artList, _ = models.FindArticleWithPage(page)
		c.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		c.Data["HasFooter"] = true
		//c.Data["PageCode"] = 1
		c.Data["HasFooter"] = true

		fmt.Println("IsLogin:", c.IsLogin, c.LoginUser)
		c.Data["Content"] = models.MakeHomeBlocks(artList, c.IsLogin)

		c.TplName = "home.html"
}