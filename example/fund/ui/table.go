package ui

import (
	"fund/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Order = byte

const (
	NONE Order = 0
	DESC       = 1
	ASC        = 2
)

type TH struct {
	*tview.TableCell
	EnableOrder bool
	col         int
	order       Order
	tb          *TB
}

func NewTH(text string, tb *TB) *TH {
	th := &TH{
		TableCell:   tview.NewTableCell(text),
		EnableOrder: true,
		tb:          tb,
	}
	th.SetClickedFunc(th.defaultClickedFunc)
	return th
}

func (th *TH) defaultClickedFunc() bool {
	if th.EnableOrder {
		prefix := " "
		// remove other col orders
		for _, h := range th.tb.headers {
			if h != th {
				h.Text = prefix + string([]rune(h.Text)[1:])
				h.order = NONE
			}
		}
		// set current col order
		order := (th.order + 1) % 3
		switch order {
		case ASC:
			prefix = "↑"
		case DESC:
			prefix = "↓"
		}
		th.Text = prefix + string([]rune(th.Text)[1:])
		th.order = order
		th.tb.orderCol = th.col
		th.tb.orderType = order
		if th.tb.orderFunc != nil {
			th.tb.orderFunc(th.col, order)
		}
	}
	return true
}

type TB struct {
	*tview.Table
	headers   []*TH
	orderCol  int
	orderType Order
	orderFunc func(col int, order Order)
}

func NewTB() *TB {
	tb := &TB{
		Table:   tview.NewTable(),
		headers: nil,
	}
	tb.SetSeparator(tview.Borders.Vertical).SetSelectable(true, false)
	return tb
}

func (tb *TB) SetHeaders(refs ...model.THReference) *TB {
	for i, r := range refs {
		th := NewTH(" "+r.Text, tb)
		th.col = i
		th.EnableOrder = r.EnableOrder
		th.SetReference(r.OrderFiled)
		// default header attributions, can set by call tview.TableCell functions with tb.Header[i]
		th.SetAlign(tview.AlignCenter).
			SetTextColor(tcell.ColorBlack).
			SetBackgroundColor(tcell.ColorGray).
			SetSelectable(false)
		tb.SetCell(0, i, th.TableCell)
		tb.headers = append(tb.headers, th)
	}
	return tb
}

func (tb *TB) SetOrderFunc(f func(col int, order Order)) *TB {
	tb.orderFunc = f
	return tb
}

func (tb *TB) UpdateRow(row int, texts ...string) {
	for i, h := range texts {
		cell := tb.GetCell(row+1, i)
		cell.SetTextColor(tcell.ColorWhite)
		cell.Text = h
		tb.SetCell(row+1, i, cell)
	}
}
