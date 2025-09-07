package main

import (
	"strconv"

	"github.com/fatih/color"
)

func colorizeStringHex(hex, str string) string {
	if len(hex) != 6 {
		return ""
	}

	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return ""
	}

	b, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		return ""
	}

	g, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		return ""
	}

	c := color.RGB(int(r), int(g), int(b))
	return c.SprintFunc()(str)
}

func colorizeStringItalic(str string) string {
	return color.New(color.Italic).SprintFunc()(str)
}

func colorizeStringBold(str string) string {
	return color.New(color.Bold).SprintFunc()(str)
}
