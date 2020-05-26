package cterm

import (
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	pages               *tview.Pages
	indexPage           *tview.Flex
	indexList           *tview.List
	contentPage         *tview.Flex
	contentPageLeft     *tview.Flex
	contentPageRight    *tview.Flex
	contentList         *tview.List
	msgList             *tview.List
	currentContentList  []ContentListSourceDataType
	currentContentData  []string
	indexListSourceData []IndexListSourceDataType
)

func Draw(app *tview.Application) {
	initIndexPage(app)
	initContentPage(app)

	pages = tview.NewPages().
		AddPage("indexPage", indexPage, true, true).
		AddPage("contentPage", contentPage, true, false)

	app.SetRoot(pages, true).EnableMouse(true)
}

func initIndexPage(app *tview.Application) {
	indexList = tview.NewList()
	indexList.
		SetBorder(true).
		SetBorderAttributes(tcell.AttrNone).
		SetBorderPadding(1, 1, 1, 1).
		SetTitle("  Welcome to New World  ").
		SetTitleColor(tcell.ColorPink)

	for _, v := range indexListSourceData {
		indexList.AddItem(v.Name, v.Desc, 0, nil)
	}

	indexList.SetSelectedFunc(func(index int, _ string, _ string, _ rune) {
		contentList.Clear()
		currentContentList = indexListSourceData[index].contentList
		for _, v := range currentContentList {
			contentList.AddItem(v.title, "", 0, nil)
		}
		pages.SwitchToPage("contentPage")

	})

	indexPage = tview.NewFlex()
	indexPage.SetBorder(false)
	indexPage.AddItem(tview.NewBox(), 0, 1, false)
	indexPage.AddItem(
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 2, false).
			AddItem(indexList, 0, 3, true).
			AddItem(tview.NewBox(), 0, 2, false),
		0, 1, true)
	indexPage.AddItem(tview.NewBox(), 0, 1, false)
}

func initContentPage(app *tview.Application) {

	contentList = tview.NewList().SetSelectedBackgroundColor(tcell.ColorPink)
	contentList.SetBorderPadding(1, 1, 1, 1)

	contentList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'i':
			inputContentView := tview.NewInputField()
			inputContentView.SetLabel("请输入类目：")
			inputContentView.SetFieldWidth(0)
			inputContentView.SetDoneFunc(func(key tcell.Key) {
				inputValue := inputContentView.GetText()
				if key == tcell.KeyEnter {
					if 0 != len(inputValue) {
						currentContentList = append(currentContentList, ContentListSourceDataType{title: inputValue, ID: len(currentContentList) + 1})
						contentList.AddItem(inputValue, "", 0, nil)
					}
					contentPageLeft.RemoveItem(inputContentView)
					app.SetFocus(contentList)
				}
			})
			contentPageLeft.AddItem(inputContentView, 10, 1, true)
			app.SetFocus(inputContentView)
		}
		return event
	})

	msgList = tview.NewList().SetSelectedBackgroundColor(tcell.ColorBlue)
	msgList.SetBorderPadding(1, 1, 1, 1)

	contentList.SetSelectedFunc(func(index int, _ string, _ string, _ rune) {
		msgList.Clear()
		app.SetFocus(msgList)
		currentContentData = currentContentList[index].msg
		for _, v := range currentContentData {
			msgList.AddItem(v, "", 0, nil)
			contentPageRight.AddItem(msgList, 0, 1, true)
		}
	})

	msgList.SetSelectedFunc(func(index int, _ string, _ string, _ rune) {
		clipboard.WriteAll(currentContentData[index])
	})

	contentPageLeft = tview.NewFlex()
	contentPageRight = tview.NewFlex()

	contentPageLeft.SetDirection(tview.FlexRow).SetBorder(true).SetTitle(" 类目 [ enter 查看内容， q退出 ]")
	contentPageLeft.AddItem(contentList, 0, 1, true)

	contentPageRight.SetBorder(true).SetTitle(" 内容 [ enter 复制内容, q 退出 ]")

	contentPage = tview.NewFlex()
	contentPage.AddItem(contentPageLeft, 0, 1, true).AddItem(contentPageRight, 0, 1, false)
}
