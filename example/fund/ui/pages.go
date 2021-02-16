package ui

import (
	"fmt"
	"fund/common"
	"fund/lang"
	"fund/model"
	"github.com/gdamore/tcell/v2"
	"github.com/lauthrul/goutil/log"
	"github.com/rivo/tview"
)

type Page struct {
	tview.Primitive
	menu  *tview.TextView
	table *TB
	Name  string
}

const (
	PageNameFund  = "fund"
	PageNameStock = "stock"
	MenuNameRank  = "rank"
	MenuNameFav   = "fav"
)

func NewFundPage() Page {
	menu := tview.NewTextView()
	fmt.Fprintf(menu, `["%s"][white]%s[white][""]  `, MenuNameRank, lang.Text(common.Lan, "navRank"))
	fmt.Fprintf(menu, `["%s"][white]%s[white][""]  `, MenuNameFav, lang.Text(common.Lan, "navFav"))
	menu.SetRegions(true).
		SetDynamicColors(true).
		Highlight(MenuNameRank).
		SetHighlightedFunc(FundPageHighlightedFunc)
	menu.SetBackgroundColor(tcell.NewHexColor(0x313335)).
		SetBorderPadding(0, 0, 2, 2)

	table := NewTB()
	table.SetHeaders(model.Fund{}.GetTitles()...)
	table.SetBackgroundColor(tcell.NewHexColor(0x313335))

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(menu, 1, 1, false).
		AddItem(table, 0, 5, false)

	page := Page{
		Primitive: layout,
		menu:      menu,
		table:     table,
		Name:      PageNameFund,
	}

	FundPageUpdate(page)

	return page
}

func FundPageUpdate(page Page) {
	go func() {
		funds, err := model.FundMarketList()
		if err != nil {
			return
		}
		for i, f := range funds {
			values := f.GetValues()
			page.table.UpdateRow(i, values...)
		}
		app.Draw()
	}()
}

func FundPageHighlightedFunc(added, removed, remaining []string) {
	log.Debug(added)
}

func NewStockPage() Page {
	menu := tview.NewTextView()
	fmt.Fprintf(menu, `["%s"][white]%s[white][""]  `, MenuNameFav, lang.Text(common.Lan, "navFav"))
	menu.SetRegions(true).
		SetDynamicColors(true).
		Highlight(MenuNameFav).
		SetHighlightedFunc(StockPageHighlightedFunc)
	menu.SetBackgroundColor(tcell.NewHexColor(0x313335)).
		SetBorderPadding(0, 0, 2, 2)

	table := NewTB()
	table.SetHeaders()
	table.SetBackgroundColor(tcell.NewHexColor(0x313335))

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(menu, 1, 1, false).
		AddItem(table, 0, 5, false)

	page := Page{
		Primitive: layout,
		menu:      menu,
		table:     table,
		Name:      PageNameStock,
	}

	StockPageUpdate(page)

	return page
}

func StockPageHighlightedFunc(added, removed, remaining []string) {
	log.Debug(added)
}

func StockPageUpdate(page Page) {
	//go func() {
	//	ticker := time.NewTicker(5 * time.Second)
	//	defer ticker.Stop()
	//	for ; true; <-ticker.C {
	//		// update stock
	//	}
	//}()
}
