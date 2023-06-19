/*
! @Author: fengzhilaoling_Go fengzhilaoling_go@163.com
! @Date: 2023-06-16 20:48:51

 ! @LastEditors: fengzhilaoling_Go fengzhilaoling_go@163.com
 ! @LastEditTime: 2023-06-19 20:39:23

! @FilePath: \frne\main.go
! @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
*/
package main

import (
	"fmt"
	"fynetime/utils"
	"image/color"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

//	var supportZones map[string]string = map[string]string{
//		"太平洋夏季时间":      "PDT",
//		"太平洋标准时区(冬令时)": "PST",
//	}

func main() {
	os.Setenv("FYNE_FONT", "C:\\Windows\\Fonts\\STZHONGS.TTF")
	myApp := app.New()
	myWindow := myApp.NewWindow("太平洋标准时区转换系统")
	black := color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}
	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
	red := color.NRGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	// -------------------标题设置-------------------------
	title := canvas.NewText("时区转换", black)
	title.TextSize = 26
	title.TextStyle.Bold = true
	contentTitle := container.New(layout.NewCenterLayout(), title)
	// ---------------------内容设置------------------------------
	input := widget.NewEntry()
	input.SetPlaceHolder("时间格式: 2006-01-02 15:04:05 & 15:04:05")
	name, _ := time.Now().Zone()
	nowZone := fmt.Sprintf("当前时区: %s", name)
	zoneText := canvas.NewText(nowZone, blue)
	prompt := canvas.NewText("转换为太平洋标准时间(自动处理东/夏令时)", blue)

	osResult := canvas.NewText("", blue)
	button := widget.NewButton("转换", func() {
		// 检查input字符串是否符合格式
		if !utils.Check_input(input) {
			input.Text = time.Now().Format("2006-01-02 15:04:05")
			input.Refresh()
			return
		}
		// layout1 := "2006-01-02 15:04:05"
		// layout2 := "15:04:05"
		layout1Bool, layout2Bool := utils.Check_input_format(input.Text)
		if !layout1Bool {
			if !layout2Bool {
				osResult.Text = fmt.Sprintf("%v", "时间字符串不符合格式请重新输入!")
				osResult.Color = red
				osResult.Refresh()
				return
			}
		}
		// 获取input中字符串
		date, _ := utils.GetUTCTime(input.Text)
		log.Println("date", date)
		resultDate := utils.GetLosAngeles(date)
		osResult.Text = fmt.Sprintf("%v", resultDate)
		osResult.Color = green
		osResult.Refresh()

		log.Println("button", resultDate)
	})

	content1 := container.New(layout.NewVBoxLayout(), input, zoneText, prompt, button, osResult)

	losInput := widget.NewEntry()
	losInput.SetPlaceHolder("时间格式: 2006-01-02 15:04:05 & 15:04:05")
	// loc, _ := time.LoadLocation("America/Los_Angeles")
	// nowTime := time.Now()
	// time.ParseInLocation()
	nowTimeStr := time.Now().Format("2006-01-02 15:04:05")
	losZone := fmt.Sprintf("当前时区: %s", nowTimeStr)
	losZoneText := canvas.NewText(losZone, blue)
	losPrompt := canvas.NewText("转换为太平洋标准时间(自动处理东/夏令时)", blue)
	result := canvas.NewText("", blue)
	losButton := widget.NewButton("转换", func() {
		// 检查input字符串是否符合格式
		if !utils.Check_input(input) {
			input.Text = time.Now().Format("2006-01-02 15:04:05")
			input.Refresh()
			return
		}
		// layout1 := "2006-01-02 15:04:05"
		// layout2 := "15:04:05"
		layout1Bool, layout2Bool := utils.Check_input_format(input.Text)
		if !layout1Bool {
			if !layout2Bool {
				result.Text = fmt.Sprintf("%v", "时间字符串不符合格式请重新输入!")
				result.Color = red
				result.Refresh()
				return
			}
		}
		// 获取input中字符串
		date, _ := utils.GetUTCTime(input.Text)
		log.Println("date", date)
		resultDate := utils.GetLosAngeles(date)
		result.Text = fmt.Sprintf("%v", resultDate)
		result.Color = green
		result.Refresh()

		log.Println("button", resultDate)
	})

	content2 := container.New(layout.NewVBoxLayout(), losInput, losZoneText, losPrompt, losButton, result)

	line := canvas.NewLine(color.Black)
	line.StrokeWidth = 1
	line.Refresh()

	content := fyne.NewContainerWithLayout(
		layout.NewMaxLayout(),
		line,
	)
	content.


	body := container.New(layout.NewGridLayout(3), content1, line, content2)

	context := container.New(layout.NewVBoxLayout(), contentTitle, body)
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.SetContent(context)
	myWindow.ShowAndRun()
}
