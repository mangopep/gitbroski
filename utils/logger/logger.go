package logger

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func Banner(text string) {
	myFigure := figure.NewFigure(text, "slant", true)
	c := color.New(color.FgHiMagenta)
	c.Println(myFigure.String())
}

func Text(text string) {
	c := color.New()
	c.AddRGB(135, 206, 250)
	c.Println(text)
}

func Error(text string) {
	c := color.New(color.FgHiRed)
	c.Println(text)
}
