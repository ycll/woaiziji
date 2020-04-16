package controllers

import (
	"strconv"
	"woaiziji/models"
)

type ShowArticleController struct {
	BaseController
}

func (c *ShowArticleController) Get()  {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	article := models.QueryArticleWithId(id)

	c.Data["Title"] = article.Title
	c.Data["Content"] = article.Content


	c.TplName = "show_article.html"
}
