package cterm

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func Keybinds(app *tview.Application) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		curFocus := app.GetFocus()

		switch event.Rune() {
		case 'j':
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModMask(0))
		case 'k':
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModMask(0))
		case 'q':
			if curFocus == indexList {
				app.Stop()
			}
			if curFocus == contentList {
				pages.SwitchToPage("indexPage")
			}
			if curFocus == msgList {
				msgList.Clear()
				app.SetFocus(contentList)
				contentPageRight.RemoveItem(msgList)
			}
		}
		return event
	})
}
