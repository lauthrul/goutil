package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Order = byte

const (
	NONE Order = 0
	ASC        = 1
	DESC       = 2
)

type TH struct {
	*tview.TableCell
	EnableSort bool
	Col        int
	order      Order
	tb         *TB
}

func NewTH(text string, tb *TB) *TH {
	th := &TH{
		TableCell:  tview.NewTableCell(text),
		EnableSort: true,
		tb:         tb,
	}
	th.SetClickedFunc(th.DefaultClickedFunc)
	return th
}

func (th *TH) DefaultClickedFunc() bool {
	prefix := " "
	// remove other col orders
	if th.Col != th.tb.GetOrderCol() {
		th.tb.SetOrderCol(th.Col)
		for _, h := range th.tb.Headers {
			if h != th {
				h.Text = prefix + string([]rune(h.Text)[1:])
				h.SetOrder(NONE)
			}
		}
	}
	// set current col order
	order := (th.GetOrder() + 1) % 3
	switch order {
	case ASC:
		prefix = "↑"
	case DESC:
		prefix = "↓"
	}
	th.Text = prefix + string([]rune(th.Text)[1:])
	th.SetOrder(order)
	return true
}

func (th *TH) SetOrder(order Order) {
	th.order = order
}

func (th *TH) GetOrder() Order {
	return th.order
}

type TB struct {
	*tview.Table
	Headers  []*TH
	orderCol int
}

func NewTB() *TB {
	tb := &TB{
		Table:    tview.NewTable(),
		Headers:  nil,
		orderCol: 0,
	}
	tb.SetSeparator('|').SetSelectable(true, false)
	return tb
}

func (tb *TB) SetHeaders(texts ...string) {
	for i, h := range texts {
		th := NewTH(" "+h, tb)
		th.Col = i
		// default header attributions, can set by call tview.TableCell functions with tb.Header[i]
		th.SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorBlack).
			SetBackgroundColor(tcell.ColorWhite).
			SetSelectable(false)
		tb.SetCell(0, i, th.TableCell)
		tb.Headers = append(tb.Headers, th)
	}
}

func (tb *TB) SetOrderCol(col int) {
	tb.orderCol = col
}

func (tb *TB) GetOrderCol() int {
	return tb.orderCol
}

func (tb *TB) UpdateRow(row int, texts ...string) {
	for i, h := range texts {
		cell := tb.GetCell(row+1, i)
		cell.SetTextColor(tcell.ColorWhite)
		cell.Text = h
		tb.SetCell(row+1, i, cell)
	}
}
