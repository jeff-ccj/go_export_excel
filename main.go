package main

import (
	_ "export_excel/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/down", "static/download")
	beego.Run()
}

