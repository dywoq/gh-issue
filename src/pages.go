package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/pkg/browser"
	"github.com/rivo/tview"
)

// pagesInitHome creates Home page to p, its identifier is "home".
func pagesInitHome(p *tview.Pages) {
	form1 := tview.NewForm()
	form1.
		AddTextView("gh-issue-deleter", "", 20, 1, true, false).
		AddButton("Github Repository", func() {
			err := browser.OpenURL("https://github.com/dywoq/gh-issue-deleter")
			if err != nil {
				log.Println(err)
			}
		}).
		SetTitleAlign(tview.AlignCenter).
		SetBackgroundColor(tcell.ColorGray).
		SetBorder(true)
	p.AddPage("home", form1, true, true)
}
