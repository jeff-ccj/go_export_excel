package models

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tealeg/xlsx"
	"strconv"
)

var (
	excelTemplate []string
)


func init() {
	// 标题模板
	excelTemplate = []string{"车牌号", "开票日期", "通行日期起", "通行日期止", "发票代码", "发票号码", "购方名称", "销方名称", "未税金额", "进项税额", "价税合计", "税率"}
}

// 创建函数
func Create(arr *[][]string) (fileName string, err error) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	arrData := *arr
	t := time.Now()
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}

	if err != nil {
		panic(err)
	}
	for i:=-1; i<len(arrData); i++ {
		row = sheet.AddRow()
		row.SetHeight(30)
		for j:=0; j<len(excelTemplate); j++ {
			cell = row.AddCell()
			if i == -1 {
				cell.Value = excelTemplate[j]
			} else {
				cell.Value = arrData[i][j]
			}
		}
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "备注：未税金额汇总、进项税额汇总、价税汇总、发票识别数量/未识别汇总"
	sheet.SetColWidth(0, 12, 30)
	fileName = t.Format("2006-01-02_150405_")
	fileName += strconv.Itoa(rand.Intn(1000000))
	fileName += ".xlsx"
	err = file.Save("./static/download/" + fileName)
	if err == nil {
		fileName = "/down/" + fileName
	}
	return
}
