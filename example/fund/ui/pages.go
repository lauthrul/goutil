package ui

import (
	"github.com/rivo/tview"
)

const (
	PageNameFund  = "fund"
	PageNameStock = "stock"
	MenuNameRank  = "rank"
	MenuNameFav   = "fav"
)

type Page struct {
	tview.Primitive
	menu  *tview.TextView
	table *TB
	Name  string
}

func NewPage(name string) Page {
	return Page{
		Name: name,
	}
}

type PageInterface interface {
	create()
	update()
	onMenuChange(added, removed, remaining []string)
	//onPrevPage()
	//onNextPage()
	//onGoPage()
	//onTableOrderChange(col int, order Order)
}
