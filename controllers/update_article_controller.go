package controllers

import "woaiziji/models"

type UpdateArticleController struct {
	BaseController
}

func (c *UpdateArticleController) Get()  {
	id, _ := c.GetInt("id")

	article := models.QueryArticleWithId(id)
	c.Data["Title"] = article.Title
	c.Data["Tags"] = article.Tags
	c.Data["Short"] = article.Short
	c.Data["Content"] = article.Content
	c.Data["Id"] = article.Id
	c.TplName = "write_article.html"
}


//修改文章
func (c *UpdateArticleController) Post() {
	id, _ := c.GetInt("id")

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")

	//实例化model，修改数据库
	art := models.Article{id, title, tags, short, content, "", 0}
	_, err := models.UpdateArticle(art)

	//返回数据给浏览器
	if err == nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	}
	c.ServeJSON()
}