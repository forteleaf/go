package main

import (
	"io/ioutil"

	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
)

func appMain(driver gxui.Driver) {
	g := new(gxuitter)
	g.LoadConfig()

	theme := dark.CreateTheme(driver)

	fontFile := g.ConfigString("FontFile")
	fontSize := g.ConfigInt("FontSize")
	if fontSize <= 0 {
		fontSize = 12
	}
	if fontFile != "" {
		b, err := ioutil.ReadFile(fontFile)
		font, err := driver.CreateFont(b, fontSize)
		if err == nil {
			theme.SetDefaultFont(font)
		}
	}
}
