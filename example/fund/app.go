package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

type App struct {
	app            *tview.Application
	left           *tview.List
	nav            *tview.TextView
	table          *TB
	center         *tview.Flex
	layout         *tview.Flex
	updateInterval time.Duration
	updateFunc     func(app *App)
}

func (a *App) Init() {
	a.app = tview.NewApplication()

	// left part
	a.left = tview.NewList().
		AddItem("基金", "", 0, func() {

		}).
		AddItem("股票", "", 0, func() {

		})
	a.left.SetBackgroundColor(tcell.NewHexColor(0x2b2b2b)).SetBorderPadding(2, 2, 1, 1)

	// center part
	a.nav = tview.NewTextView().SetRegions(true).SetDynamicColors(true).
		SetHighlightedFunc(func(added, removed, remaining []string) {

		})
	fmt.Fprintf(a.nav, `["%d"][white]%s[white][""]  `, 0, "排行")
	fmt.Fprintf(a.nav, `["%d"][white]%s[white][""]  `, 1, "自选")
	a.nav.Highlight("0")
	a.nav.SetBackgroundColor(tcell.NewHexColor(0x313335)).SetBorderPadding(0, 0, 2, 2)

	a.table = NewTB()
	a.table.SetHeaders(Fund{}.GetTitles()...)
	a.table.SetBackgroundColor(tcell.NewHexColor(0x313335))

	a.center = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(a.nav, 1, 1, false).
		AddItem(a.table, 0, 5, false)

	// layout
	a.layout = tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(a.left, 6, 1, false).
		AddItem(a.center, 0, 5, false)

	a.app.SetRoot(a.layout, true).EnableMouse(true)
}

func (a *App) Run() {
	if a.updateInterval > 0 && a.updateFunc != nil {
		go func() {
			ticker := time.NewTicker(a.updateInterval)
			defer ticker.Stop()
			for ; true; <-ticker.C {
				a.updateFunc(a)
			}
		}()
	}
	if err := a.app.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Update(data interface{}) {
	switch data.(type) {
	case []Fund:
		funds := data.([]Fund)
		for i, f := range funds {
			values := f.GetValues()
			a.table.UpdateRow(i, values...)
		}
		a.app.Draw()
	}
}
