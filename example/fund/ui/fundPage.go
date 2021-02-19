package ui

import (
	"fmt"
	"fund/common"
	"fund/lang"
	"github.com/lauthrul/goutil/log"
	"github.com/rivo/tview"
)

const (
	size = 20
)

type FundPage struct {
	Page
}

func NewFundPage() FundPage {
	page := FundPage{
		Page: NewPage(PageNameFund),
	}
	page.create()
	return page
}

func (f *FundPage) create() {
	// menu
	f.menu = tview.NewTextView()
	fmt.Fprintf(f.menu, `["%s"][white]%s[white][""]  `, MenuNameRank, lang.Text(common.Lan, "navRank"))
	fmt.Fprintf(f.menu, `["%s"][white]%s[white][""]  `, MenuNameFav, lang.Text(common.Lan, "navFav"))
	f.menu.SetRegions(true).
		SetDynamicColors(true).
		Highlight(MenuNameRank).
		SetHighlightedFunc(f.onMenuChange)
	f.menu.SetBackgroundColor(bgColor).
		SetBorderPadding(0, 0, 2, 2)

	// pages
	f.pages = tview.NewPages()
	fundRankPage := NewFundRankPage()
	f.pages.AddPage(MenuNameRank, fundRankPage, true, true)

	// layout
	f.Primitive = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(f.menu, 1, 1, false).
		AddItem(f.pages, 0, 5, false)
}

func (f *FundPage) onMenuChange(added, removed, remaining []string) {
	menu := added[0]
	log.Debug("switch menu:", menu)
	switch menu {
	case MenuNameRank:
		f.pages.SwitchToPage(MenuNameRank)
	case MenuNameFav:
		f.pages.SwitchToPage(MenuNameFav)
	}
}
