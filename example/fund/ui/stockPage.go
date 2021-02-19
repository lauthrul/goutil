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
	s.menu = tview.NewTextView()
	fmt.Fprintf(s.menu, `["%s"][white]%s[white][""]  `, MenuNameFav, lang.Text(common.Lan, "navFav"))
	s.menu.SetRegions(true).
		SetDynamicColors(true).
		Highlight(MenuNameFav).
		SetHighlightedFunc(s.onMenuChange)
	s.menu.SetBackgroundColor(bgColor).
		SetBorderPadding(0, 0, 2, 2)

	// pages
	s.pages = tview.NewPages()

	// layout
	s.Primitive = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(s.menu, 1, 1, false).
		AddItem(s.pages, 0, 5, false)
}

func (s *StockPage) onMenuChange(added, removed, remaining []string) {
	log.Debug("switch menu:", added[0])
}
