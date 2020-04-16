package main

import (
	"github.com/astaxie/beego"
	_ "woaiziji/routers"
	"woaiziji/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

