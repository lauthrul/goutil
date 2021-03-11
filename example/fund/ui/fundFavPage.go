package ui

import (
	"fmt"
	"fund/common"
	"fund/lang"
	"fund/model"
	"github.com/gdamore/tcell/v2"
	"github.com/lauthrul/goutil/log"
	"github.com/lauthrul/goutil/util"
	"github.com/rivo/tview"
	"strconv"
)

type FundFavPage struct {
	tview.Primitive
	parent    *tview.Pages
	layout    *tview.Flex
	table     *TB
	btnPrev   *tview.Button
	btnNext   *tview.Button
	edit      *tview.InputField
	total     *tview.TextView
	btnGo     *tview.Button
	curPage   int
	totalPage int
}

func NewFundFavPage(parent *tview.Pages) FundFavPage {
	f := FundFavPage{parent: parent}

	// table
	refs := []model.THMeta{{"#", false, ""}}
	refs = append(refs, api.GetFundFavMeta()...)
	f.table = NewTB()
	f.table.SetHeaders(refs...).
		SetOrderFunc(f.onTableOrderChange).
		SetSelectionChangedFunc(f.onSelectChange)/*.
		SetSelectedFunc(f.onShowDetail).
		SetMouseCapture(f.mouseCapture)*/
	f.table.SetBackgroundColor(bgColor)

	// page navigator
	box := tview.NewTextView().SetBackgroundColor(bgColor)
	f.btnPrev = tview.NewButton(lang.Text(common.Lan, "btnPrevPage"))
	f.btnPrev.SetSelectedFunc(f.onPrevPage)
	f.btnPrev.SetBackgroundColor(btnColor)
	f.btnNext = tview.NewButton(lang.Text(common.Lan, "btnNextPage"))
	f.btnNext.SetSelectedFunc(f.onNextPage)
	f.btnNext.SetBackgroundColor(btnColor)
	f.edit = tview.NewInputField()
	f.edit.SetFieldBackgroundColor(edtColor)
	f.edit.SetDoneFunc(f.onEditDone)
	f.total = tview.NewTextView()
	f.total.SetBackgroundColor(edtColor)
	f.btnGo = tview.NewButton(lang.Text(common.Lan, "btnGo"))
	f.btnGo.SetSelectedFunc(f.onGoPage)
	f.btnGo.SetBackgroundColor(btnColor)
	pager := tview.NewFlex()
	pager.SetDirection(tview.FlexColumn).
		AddItem(box, 0, 8, false).
		AddItem(f.edit, 8, 1, false).
		AddItem(f.total, 8, 1, false).
		AddItem(f.btnGo, 4, 1, false).
		AddItem(f.btnPrev, 4, 1, false).
		AddItem(f.btnNext, 4, 1, false)

	// layout
	f.layout = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(f.table, 0, 5, true).
		AddItem(pager, 1, 1, false)
	f.Primitive = f.layout

	f.update()

	return f
}

func (f *FundFavPage) update() {
	f.updatePage(1, size, 1, ASC)
}

func (f *FundFavPage) onPrevPage() {
	f.curPage -= 1
	if f.curPage < 1 {
		f.curPage = 1
		return
	}
	log.Debug("prev page", f.curPage)
	f.updatePage(f.curPage, size, f.table.orderCol, f.table.orderType)
}

func (f *FundFavPage) onNextPage() {
	f.curPage += 1
	if f.curPage > f.totalPage {
		f.curPage = f.totalPage
		return
	}
	log.Debug("next page", f.curPage)
	f.updatePage(f.curPage, size, f.table.orderCol, f.table.orderType)
}

func (f *FundFavPage) onGoPage() {
	n, err := strconv.Atoi(f.edit.GetText())
	if err != nil {
		log.Error("invalid page:", err.Error())
		return
	}
	if n == f.curPage {
		return
	}
	f.curPage = util.Clamp(f.curPage, 1, f.totalPage)
	log.Debug("go page", n)
	f.updatePage(n, size, f.table.orderCol, f.table.orderType)
}

func (f *FundFavPage) onEditDone(key tcell.Key) {
	f.onGoPage()
}

func (f *FundFavPage) updatePage(page, pageSize, orderCol int, orderType Order) {
	go func() {
		sortCode, sortType := "", ""
		if h := f.table.headers[orderCol]; h != nil {
			if ref, ok := h.Reference.(string); ok {
				sortCode = ref
				if orderType == ASC {
					sortType = "asc"
				} else if orderType == DESC {
					sortType = "desc"
				}
			}
		}
		result, err := api.GetFundFav(model.FundFavArg{
			Group:     "",
			IsFav:     -1,
			SortCode:  sortCode,
			SortType:  sortType,
			PageIndex: page,
			PageSize:  pageSize,
		})
		if err != nil {
			return
		}
		for i, fund := range result.List {
			f.curPage = result.PageIndex
			f.totalPage = result.TotalPage
			values := []string{ /*"☆", */ fmt.Sprintf("%d", (page-1)*pageSize+i+1)}
			values = append(values, fund.GetValues()...)
			f.table.UpdateRow(i, values...)
			f.table.GetCell(i+1, 0).SetReference(fund)
			f.edit.SetText(fmt.Sprintf("%d", f.curPage))
			f.total.SetText(fmt.Sprintf("/%d", result.TotalPage))
		}
		count := len(result.List)
		rows := f.table.GetRowCount() - 1
		if count < rows {
			for i := rows; i > count; i-- {
				f.table.RemoveRow(i)
			}
		}
		app.Draw()
	}()
}

func (f *FundFavPage) onTableOrderChange(col int, order Order) {
	log.Debug("order by", col, order)
	f.updatePage(1, size, col, order)
}

func (f *FundFavPage) onSelectChange(row, column int) {
	log.Debug("select", row, column)
}

//func (f *FundFavPage) onShowDetail(row, column int) {
//	log.Debug("show detail", row, column)
//	ref := f.table.GetCell(row, 0).GetReference()
//	if fund, ok := ref.(model.EastMoneyFund); ok {
//		log.DebugF("%+v", fund)
//		form := tview.NewForm().
//			AddInputField("Fund", fmt.Sprintf("%s %s", fund.Code, fund.Name), 20, nil, nil).
//			AddInputField("NetDate", fund.NetDate, 20, nil, nil).
//			AddInputField("NetValue", fmt.Sprintf("%s|%s", fund.NetValue, fund.TotalNetValue), 20, nil, nil).
//			AddButton("Save", nil).
//			AddButton("Quit", func() {
//				f.parent.RemovePage("form")
//			}).SetFocus(0)
//		title := fmt.Sprintf("%s %s", fund.Code, fund.Name)
//		form.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignLeft).
//			SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
//				if event.Key() == tcell.KeyESC {
//					f.parent.RemovePage("form")
//				}
//				return event
//			})
//		modal := func(p tview.Primitive) tview.Primitive {
//			return tview.NewGrid().SetColumns(-1, -3, -1).SetRows(-1, -3, -1).AddItem(p, 1, 1, 1, 1, 0, 0, true)
//		}
//		f.parent.AddPage("form", modal(form), true, true).ShowPage("form")
//	}
//}
//
//func (f *FundFavPage) mouseCapture(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
//	row, col := f.table.GetSelection()
//	if action == tview.MouseLeftDoubleClick {
//		log.Debug("double click", row, col)
//		f.onShowDetail(row, col)
//	}
//	return action, event
//}
