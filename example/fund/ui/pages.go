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
	pages *tview.Pages
	Name  string
}

func NewPage(name string) Page {
	return Page{
		Name: name,
	}
}
