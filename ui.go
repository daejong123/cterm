package cterm

import (
	"strconv"

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
	currentContentData  *[]string
	currentContentList  *[]ContentListSourceDataType
	indexListSourceData []IndexListSourceDataType
)

func Draw(app *tview.Application) {
	initIndexPage(app)
	initContentPage(app)

	pages = tview.NewPages().
		AddPage("indexPage", indexPage, true, true).
		AddPage("contentPage", contentPage, true, false)

	app.SetRoot(pages, true).EnableMouse(false)
}

func initIndexPage(app *tview.Application) {
	indexList = tview.NewList().
		SetSelectedBackgroundColor(tcell.ColorDeepSkyBlue).
		SetSecondaryTextColor(tcell.ColorGray)
	indexList.
		SetBorder(true).
		SetBorderColor(tcell.Color20).
		SetBorderPadding(5, 5, 5, 5).
		SetBorderAttributes(tcell.AttrNone).
		SetTitle("  Welcome to New World  ").
		SetTitleColor(tcell.ColorPink)

	for _, v := range indexListSourceData {
		indexList.AddItem(strconv.Itoa(v.ID)+"、"+v.Name, v.Desc, 0, nil)
	}

	indexList.SetSelectedFunc(func(index int, _ string, _ string, _ rune) {
		contentList.Clear()
		selectedIndexSourceData := indexListSourceData[index]
		currentContentList = selectedIndexSourceData.ContentList
		for _, v := range *currentContentList {
			contentList.AddItem(strconv.Itoa(v.ID)+"、"+v.Title, "", 0, nil)
		}
		pages.SwitchToPage("contentPage")
	})

	indexPage = tview.NewFlex()
	indexPage.SetBorder(false)
	indexPage.AddItem(tview.NewBox(), 0, 1, false)
	indexPage.AddItem(
		tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tview.NewBox(), 0, 1, false).
			AddItem(indexList, 0, 3, true).
			AddItem(tview.NewBox(), 0, 1, false),
		0, 1, true)
	indexPage.AddItem(tview.NewBox(), 0, 1, false)
}

func initContentPage(app *tview.Application) {

	contentList = tview.NewList().
		SetSelectedBackgroundColor(tcell.ColorDeepSkyBlue)

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
						*currentContentList = append(*currentContentList, ContentListSourceDataType{Title: inputValue, ID: len(*currentContentList) + 1, Msg: &[]string{}})
						contentList.AddItem(inputValue, "", 0, nil)
						writeToDatafile(indexListSourceData)
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

	contentList.SetSelectedFunc(func(index int, _ string, _ string, _ rune) {
		msgList.Clear()
		app.SetFocus(msgList)
		currentContentData = (*currentContentList)[index].Msg
		for _, v := range *currentContentData {
			msgList.AddItem(v, "", 0, nil)
		}
		contentPageRight.AddItem(msgList, 0, 1, true)
	})

	msgList = tview.NewList().
		SetSelectedBackgroundColor(tcell.ColorDeepSkyBlue)
	msgList.SetSelectedFunc(func(index int, _ string, _ string, _ rune) {
		clipboard.WriteAll((*currentContentData)[index])
	})
	msgList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'i':
			inputContentView := tview.NewInputField()
			inputContentView.SetLabel("请输入内容：")
			inputContentView.SetFieldWidth(0)
			inputContentView.SetDoneFunc(func(key tcell.Key) {
				inputValue := inputContentView.GetText()
				if key == tcell.KeyEnter {
					if 0 != len(inputValue) {
						*currentContentData = append(*currentContentData, inputValue)
						msgList.AddItem(inputValue, "", 0, nil)
						writeToDatafile(indexListSourceData)
					}
					contentPageRight.RemoveItem(inputContentView)
					app.SetFocus(msgList)
				}
			})
			contentPageRight.AddItem(inputContentView, 10, 1, true)
			app.SetFocus(inputContentView)
		}
		return event
	})

	contentPageLeft = tview.NewFlex()
	contentPageRight = tview.NewFlex()

	contentPageLeft.
		SetDirection(tview.FlexRow).
		SetBorder(true).
		SetBorderPadding(0, 0, 2, 2).
		SetTitle(" 类目 [ enter 查看内容， q退出 ]")
	contentPageLeft.AddItem(contentList, 0, 1, true)

	contentPageRight.
		SetDirection(tview.FlexRow).
		SetBorderPadding(0, 0, 2, 2).
		SetBorder(true).
		SetTitle(" 内容 [ enter 复制内容, q 退出 ]")

	contentPage = tview.NewFlex()
	contentPage.AddItem(contentPageLeft, 0, 1, true).AddItem(contentPageRight, 0, 1, false)
}
