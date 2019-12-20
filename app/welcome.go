/*
@Time : 2019/12/20 16:04
@Software: GoLand
@File : welcome
@Author : Bingo <airplayx@gmail.com>
*/
package _app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"net/url"
)

func WelcomeScreen() fyne.CanvasObject {
	link, err := url.Parse("https://github.com/airplayx")
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), layout.NewSpacer()),
		widget.NewHyperlinkWithStyle("github.com/airplayx", link, fyne.TextAlignCenter, fyne.TextStyle{}),
		layout.NewSpacer(),
	)
}