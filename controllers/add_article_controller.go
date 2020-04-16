package controllers

import (
	"fmt"
	"time"
	"woaiziji/models"
)

type AddArticleController struct {
	BaseController
}

func (c *AddArticleController) Get()  {
	c.TplName = "write_article.html"
}

func (c *AddArticleController) Post() {

	//获取浏览器传输的数据，通过表单的name属性获取值
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中

	fmt.Println(c.LoginUser)
	art := models.Article{Id:0, Title: title, Tags: tags, Short: short, Content: content, Author: c.LoginUser.(string), Createtime: time.Now().Unix()}
	_, err := models.AddArticle(art)

	//返回数据给浏览器
	var response map[string]interface{}
	if err == nil {
		//无误
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	c.Data["json"] = response
	c.ServeJSON()

}
