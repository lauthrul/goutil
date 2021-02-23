package ui

import (
	"fund/common"
	"fund/lang"
	"fund/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	api    model.Api
	app    *tview.Application
	nav    *tview.List
	pages  *tview.Pages
	layout *tview.Flex
)

var (
	navBgColor  = tcell.NewHexColor(0x2b2b2b)
	bgColor     = tcell.NewHexColor(0x313335)
	btnColor    = tcell.NewHexColor(0x34424d)
	edtColor    = tcell.NewHexColor(0x3e5c73)
	borderColor = tcell.NewHexColor(0x555555)
)

func Run() {
	app = tview.NewApplication()
	api = model.NewEastMoneyApi()

	// nav part
	nav = tview.NewList().
		AddItem(lang.Text(common.Lan, "menuFund"), "", 0, onNavFund).
		AddItem(lang.Text(common.Lan, "menuStock"), "", 0, onNavStock)
	nav.SetBackgroundColor(navBgColor).SetBorderPadding(2, 2, 1, 1)

	// main part
	pageFund := NewFundPage()
	pageStock := NewStockPage()
	pages = tview.NewPages().
		AddPage(pageFund.Name, pageFund, true, true).
		AddPage(pageStock.Name, pageStock, true, false)

	// layout
	layout = tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(nav, 8, 1, false).
		AddItem(pages, 0, 5, true)

	app.SetRoot(layout, true).EnableMouse(true)

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func onNavFund() {
	pages.SwitchToPage(PageNameFund)
}

func onNavStock() {
	pages.SwitchToPage(PageNameStock)
}
