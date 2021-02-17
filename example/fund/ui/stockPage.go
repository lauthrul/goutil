package ui

import (
	"fmt"
	"fund/common"
	"fund/lang"
	"github.com/lauthrul/goutil/log"
	"github.com/rivo/tview"
)

type StockPage struct {
	Page
}

func NewStockPage() StockPage {
	page := StockPage{
		Page: NewPage(PageNameStock),
	}
	page.create()
	return page
}

func (s *StockPage) create() {
	menu := tview.NewTextView()
	fmt.Fprintf(menu, `["%s"][white]%s[white][""]  `, MenuNameFav, lang.Text(common.Lan, "navFav"))
	menu.SetRegions(true).
		SetDynamicColors(true).
		Highlight(MenuNameFav).
		SetHighlightedFunc(s.onMenuChange)
	menu.SetBackgroundColor(bgColor).
		SetBorderPadding(0, 0, 2, 2)

	table := NewTB()
	table.SetHeaders()
	table.SetBackgroundColor(bgColor)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(menu, 1, 1, false).
		AddItem(table, 0, 5, false)

	s.menu = menu
	s.table = table
	s.Primitive = layout
	s.update()
}

func (s *StockPage) update() {
	//go func() {
	//	ticker := time.NewTicker(5 * time.Second)
	//	defer ticker.Stop()
	//	for ; true; <-ticker.C {
	//		// update stock
	//	}
	//}()
}

func (s *StockPage) onMenuChange(added, removed, remaining []string) {
	log.Debug("switch menu:", added[0])
}

func (s *StockPage) onPrevPage() {
	log.Debug("prev page")
}

func (s *StockPage) onNextPage() {
	log.Debug("next page")
}

func (s *StockPage) onGoPage() {
	log.Debug("go page")
}
