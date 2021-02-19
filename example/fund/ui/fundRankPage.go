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

type FundRankPage struct {
	tview.Primitive
	table     *TB
	btnPrev   *tview.Button
	btnNext   *tview.Button
	edit      *tview.InputField
	total     *tview.TextView
	btnGo     *tview.Button
	curPage   int
	totalPage int
}

func NewFundRankPage() FundRankPage {
	f := FundRankPage{}

	// table
	refs := []model.THReference{ /*{" ", false, ""},*/ {"#", false, ""}}
	refs = append(refs, api.GetTHReference()...)
	f.table = NewTB()
	f.table.SetHeaders(refs...).
		SetOrderFunc(f.onTableOrderChange)
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
	f.Primitive = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(f.table, 0, 5, false).
		AddItem(pager, 1, 1, false)

	f.update()

	return f
}

func (f *FundRankPage) update() {
	f.updatePage(1, size, 1, ASC)
}

func (f *FundRankPage) onPrevPage() {
	f.curPage -= 1
	if f.curPage < 1 {
		return
	}
	log.Debug("prev page", f.curPage)
	f.updatePage(f.curPage, size, f.table.orderCol, f.table.orderType)
}

func (f *FundRankPage) onNextPage() {
	f.curPage += 1
	if f.curPage > f.totalPage {
		return
	}
	log.Debug("next page", f.curPage)
	f.updatePage(f.curPage, size, f.table.orderCol, f.table.orderType)
}

func (f *FundRankPage) onGoPage() {
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

func (f *FundRankPage) onEditDone(key tcell.Key) {
	f.onGoPage()
}

func (f *FundRankPage) updatePage(page, pageSize, orderCol int, orderType Order) {
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
		result, err := api.GetFundRank(model.FundRankArg{
			FundType:    "",
			FundCompany: "",
			SortCode:    sortCode,
			SortType:    sortType,
			StartDate:   "",
			EndDate:     "",
			PageIndex:   page,
			PageNumber:  pageSize,
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

func (f *FundRankPage) onTableOrderChange(col int, order Order) {
	log.Debug("order by", col, order)
	f.updatePage(1, size, col, order)
}
