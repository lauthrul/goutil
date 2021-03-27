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
	parent      *tview.Pages
	layout      *tview.Flex
	table       *TB
	groups      *tview.DropDown
	btnAddGroup *tview.Button
	btnUpdate   *tview.Button
	btnPrev     *tview.Button
	btnNext     *tview.Button
	edit        *tview.InputField
	total       *tview.TextView
	btnGo       *tview.Button
	curGroup    string
	curPage     int
	totalPage   int
}

func NewFundFavPage(parent *tview.Pages) FundFavPage {
	f := FundFavPage{parent: parent}

	// table
	refs := []model.THMeta{{"#", false, ""}}
	refs = append(refs, api.GetFundFavMeta()...)
	f.table = NewTB()
	f.table.SetHeaders(refs...).
		SetOrderFunc(f.onTableOrderChange).
		SetSelectionChangedFunc(f.onSelectChange) /*.
	SetSelectedFunc(f.onShowDetail).
	SetMouseCapture(f.mouseCapture)*/
	f.table.SetBackgroundColor(bgColor)

	//
	f.groups = tview.NewDropDown()
	groups := []string{"--分组--"}
	g, _ := model.ListGroup()
	groups = append(groups, g...)
	f.groups.SetOptions(groups, f.onGroupChange).SetCurrentOption(0).
		SetFieldBackgroundColor(btnColor).SetBackgroundColor(btnColor)

	f.btnAddGroup = tview.NewButton("+").SetSelectedFunc(f.onAddGroup)
	f.btnAddGroup.SetBackgroundColor(btnColor)

	f.btnUpdate = tview.NewButton("∽").SetSelectedFunc(f.onUpdateEstimate)
	f.btnUpdate.SetBackgroundColor(btnColor)

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
		AddItem(f.groups, 8, 1, false).
		AddItem(f.btnAddGroup, 4, 1, false).
		AddItem(f.btnUpdate, 4, 1, false).
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
	f.layout.SetMouseCapture(f.onMouseCapture)
	f.Primitive = f.layout

	f.update()

	return f
}

func (f *FundFavPage) update() {
	f.updatePage("", 1, size, 1, ASC)
}

func (f *FundFavPage) onPrevPage() {
	f.curPage -= 1
	if f.curPage < 1 {
		f.curPage = 1
		return
	}
	log.Debug("prev page", f.curPage)
	f.updatePage(f.curGroup, f.curPage, size, f.table.orderCol, f.table.orderType)
}

func (f *FundFavPage) onNextPage() {
	f.curPage += 1
	if f.curPage > f.totalPage {
		f.curPage = f.totalPage
		return
	}
	log.Debug("next page", f.curPage)
	f.updatePage(f.curGroup, f.curPage, size, f.table.orderCol, f.table.orderType)
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
	f.updatePage(f.curGroup, n, size, f.table.orderCol, f.table.orderType)
}

func (f *FundFavPage) onEditDone(key tcell.Key) {
	f.onGoPage()
}

func (f *FundFavPage) updatePage(group string, page, pageSize, orderCol int, orderType Order) {
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
			Group:     group,
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
	f.updatePage(f.curGroup, 1, size, col, order)
}

func (f *FundFavPage) onSelectChange(row, column int) {
	log.Debug("select", row, column)
}

func (f *FundFavPage) onGroupChange(text string, index int) {
	log.Debug("onGroupChange: ", text, index)
	if index == 0 {
		text = ""
	}
	f.curGroup = text
	f.updatePage(f.curGroup, f.curPage, size, f.table.orderCol, f.table.orderType)
}

func (f *FundFavPage) onAddGroup() {
	const formName = "addGroupForm"
	var addGroupForm *tview.Form
	fnCancel := func() {
		f.parent.RemovePage(formName)
	}
	fnOK := func() {
		groupCtrl := addGroupForm.GetFormItemByLabel("Group").(*tview.InputField)
		group := groupCtrl.GetText()
		log.Debug("add group: ", group)
		if err := model.AddGroup(group); err == nil {
			f.groups.AddOption(group, nil)
		}
		fnCancel()
	}
	addGroupForm = tview.NewForm().
		AddInputField("Group", "", 0, nil, nil).
		AddButton(lang.Text(common.Lan, "btnOK"), fnOK).
		AddButton(lang.Text(common.Lan, "btnCancel"), fnCancel).
		SetFocus(0)
	addGroupForm.SetBorder(true).SetTitle(lang.Text(common.Lan, "NewGroup")).SetTitleAlign(tview.AlignLeft).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyESC {
				fnCancel()
			}
			return event
		})
	modal := func(p tview.Primitive) tview.Primitive {
		return tview.NewGrid().SetColumns(-1, 30, -1).SetRows(-1, 8, -1).AddItem(p, 1, 1, 1, 1, 0, 0, true)
	}
	f.parent.AddPage(formName, modal(addGroupForm), true, true).ShowPage(formName)
}

func (f *FundFavPage) onMouseCapture(action tview.MouseAction, event *tcell.EventMouse) (tview.MouseAction, *tcell.EventMouse) {
	if action == tview.MouseLeftDown {
		x, y := event.Position()
		inRect := f.groups.GetList().InRect(x, y)
		if inRect {
			f.groups.MouseHandler()(action, event, func(p tview.Primitive) {
				app.SetFocus(p)
			})
		}
	}
	return action, event
}

func (f *FundFavPage) onUpdateEstimate() {

}
