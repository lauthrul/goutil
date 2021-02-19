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

const (
	size = 20
)

type FundPage struct {
	Page
	btnPrev   *tview.Button
	btnNext   *tview.Button
	edit      *tview.InputField
	total     *tview.TextView
	btnGo     *tview.Button
	curPage   int
	totalPage int
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
	menu := tview.NewTextView()
	fmt.Fprintf(menu, `["%s"][white]%s[white][""]  `, MenuNameRank, lang.Text(common.Lan, "navRank"))
	fmt.Fprintf(menu, `["%s"][white]%s[white][""]  `, MenuNameFav, lang.Text(common.Lan, "navFav"))
	menu.SetRegions(true).
		SetDynamicColors(true).
		Highlight(MenuNameRank).
		SetHighlightedFunc(f.onMenuChange)
	menu.SetBackgroundColor(bgColor).
		SetBorderPadding(0, 0, 2, 2)

	// table
	refs := []model.THReference{ /*{" ", false, ""},*/ {"#", false, ""}}
	refs = append(refs, api.GetTHReference()...)
	table := NewTB()
	table.SetHeaders(refs...).
		SetOrderFunc(f.onTableOrderChange)
	table.SetBackgroundColor(bgColor)

	// page navigator
	box := tview.NewTextView().SetBackgroundColor(bgColor)
	btnPrev := tview.NewButton(lang.Text(common.Lan, "btnPrevPage"))
	btnPrev.SetSelectedFunc(f.onPrevPage)
	btnPrev.SetBackgroundColor(btnColor)
	btnNext := tview.NewButton(lang.Text(common.Lan, "btnNextPage"))
	btnNext.SetSelectedFunc(f.onNextPage)
	btnNext.SetBackgroundColor(btnColor)
	edit := tview.NewInputField()
	edit.SetFieldBackgroundColor(edtColor)
	edit.SetDoneFunc(f.onEditDone)
	total := tview.NewTextView()
	total.SetBackgroundColor(edtColor)
	btnGo := tview.NewButton(lang.Text(common.Lan, "btnGo"))
	btnGo.SetSelectedFunc(f.onGoPage)
	btnGo.SetBackgroundColor(btnColor)
	pager := tview.NewFlex()
	pager.SetDirection(tview.FlexColumn).
		AddItem(box, 0, 8, false).
		AddItem(edit, 8, 1, false).
		AddItem(total, 8, 1, false).
		AddItem(btnGo, 4, 1, false).
		AddItem(btnPrev, 4, 1, false).
		AddItem(btnNext, 4, 1, false)

	// layout
	layout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(menu, 1, 1, false).
		AddItem(table, 0, 5, false).
		AddItem(pager, 1, 1, false)

	f.menu = menu
	f.table = table
	f.btnPrev = btnPrev
	f.btnNext = btnNext
	f.edit = edit
	f.total = total
	f.btnGo = btnGo
	f.Primitive = layout
	f.update()
}

func (f *FundPage) update() {
	f.updatePage(1, size, 1, ASC)
}

func (f *FundPage) onMenuChange(added, removed, remaining []string) {
	menu := added[0]
	log.Debug("switch menu:", menu)
	//if menu == MenuNameFav {
	//	favs, err := model.FundFavList("210008", "007047")
	//	if err != nil {
	//		log.Error(err)
	//	}
	//	log.Debug(favs)
	//}
}

func (f *FundPage) onPrevPage() {
	f.curPage -= 1
	if f.curPage < 1 {
		return
	}
	log.Debug("prev page", f.curPage)
	f.updatePage(f.curPage, size, f.table.orderCol, f.table.orderType)
}

func (f *FundPage) onNextPage() {
	f.curPage += 1
	if f.curPage > f.totalPage {
		return
	}
	log.Debug("next page", f.curPage)
	f.updatePage(f.curPage, size, f.table.orderCol, f.table.orderType)
}

func (f *FundPage) onGoPage() {
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

func (f *FundPage) onEditDone(key tcell.Key) {
	f.onGoPage()
}

func (f *FundPage) updatePage(page, pageSize, orderCol int, orderType Order) {
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

func (f *FundPage) onTableOrderChange(col int, order Order) {
	log.Debug("order by", col, order)
	f.updatePage(1, size, col, order)
}
