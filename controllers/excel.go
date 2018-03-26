package controllers

import (
	//"encoding/json"

	"github.com/astaxie/beego"
	"export_excel/models"
	"fmt"
	"encoding/json"
)

// Operations about Excel
type ExcelController struct {
	beego.Controller
}




// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ExcelController) Post() {
	var data map[string][][]string
	a := o.Ctx.Input.RequestBody
	json.Unmarshal(a, &data)
	arr := data["data"]
	var createUrl string
	var err error
	resData := make(map[string]string)
	createUrl, err = models.Create(&arr)
	fmt.Println(resData)
	resData["url"] = createUrl
	if err == nil {
		o.Data["json"] = resData
	}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ExcelController) Get() {
	o.Ctx.WriteString("Get")

}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ExcelController) GetAll() {
	var a string
	var err error
	a1 := [][]string{
		{"a123", "a456", "a789", "a123", "a456", "a789", "a123", "a456", "a789", "a123", "a456", "a789"},
		{"a123", "a456", "a789", "a123", "a456", "a789", "a123", "a456", "a789", "a123", "a456", "a789"},
		{"a123", "a456", "a789", "a123", "a456", "a789", "a123", "a456", "a789", "a123", "a456", "a789"}}
	a, err = models.Create(&a1)
	fmt.Println(a)
	fmt.Println(err)
	if err == nil {
		o.Data["string"] = a
	}
	o.ServeJSON()
}
