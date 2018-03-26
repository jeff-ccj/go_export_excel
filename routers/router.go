package routers

import (
	"export_excel/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/excel",
			beego.NSInclude(
				&controllers.ExcelController{},
			),
		),
	)
	beego.AddNamespace(ns)
}