package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/itsknob/hawk-tui/envpath"
	"github.com/rivo/tview"
)

func onSubmit(form *tview.Form, inputLabel string) string {
	inputFormItem := form.GetFormItemByLabel(inputLabel).(*tview.InputField)
	return inputFormItem.GetText()
}

func refreshList(list *tview.List, entries []string) {
	list.Clear()

	for idx, item := range entries {
		list.AddItem(item, "", int32(idx+97), nil)
	}
}

func switchToForm(app *tview.Application, newFocus *tview.Form) {
	// todo will explode
	app.SetFocus(newFocus)
}
func switchToList(app *tview.Application, newFocus *tview.List) {
	app.SetFocus(newFocus)
}
func setDirText(selectedText string) {
	formInput := form.GetFormItemByLabel("Directory").(*tview.InputField)
	formInput.SetText(selectedText)
}

var (
	app  *tview.Application
	form *tview.Form
	list *tview.List
)

func main() {
	app = tview.NewApplication()
	pathObject := envpath.Path{}
	pathObject.Init()

	list = tview.NewList().ShowSecondaryText(false)
	list.SetBorder(true).SetTitle("Path")

	for idx, item := range pathObject.Entries {
		list.AddItem(item, "", int32(idx+97), func() {
			mainText, _ := list.GetItemText(idx)
			setDirText(mainText)
			switchToForm(app, form)
		})
	}

	form = tview.NewForm().
		AddInputField("Directory", "", 0, nil, nil)
	form.
		AddButton("Submit", func() {
			result := onSubmit(form, "Directory")
			currentValueField := form.GetFormItemByLabel("Current Value").(*tview.TextView)
			currentValueField.SetText(result)
		}).
		AddButton("Add to Front", func() {
			value := form.GetFormItemByLabel("Directory").(*tview.InputField).GetText()
			_, err := pathObject.AddToPathFront(value)
			if err != nil {
				panic(err)
			}
			refreshList(list, pathObject.Entries)
		}).
		AddButton("Add to Back", func() {
			value := form.GetFormItemByLabel("Directory").(*tview.InputField).GetText()
			_, err := pathObject.AddToPathBack(value)
			if err != nil {
				panic(err)
			}
			refreshList(list, pathObject.Entries)
		}).
		AddButton("Remove from Path", func() {
			value := form.GetFormItemByLabel("Directory").(*tview.InputField).GetText()
			_, err := pathObject.RemoveFromPath(value)
			if err != nil {
				panic(err)
			}
			refreshList(list, pathObject.Entries)
		}).
		AddTextView("Current Value", "", 0, 0, false, false)

	form.SetBorder(true).SetTitle("Operation Go!")

	container := tview.NewFlex()
	container.SetDirection(tview.FlexRow)
	container.AddItem(list, 0, 3, true).AddItem(form, 0, 2, false)

	container.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyESC:
			{
				if list.HasFocus() {
					switchToForm(app, form)
				} else {
					switchToList(app, list)
				}
			}
		default:
			break
		}
		return event
	})

	if err := app.SetRoot(container, true).SetFocus(container).Run(); err != nil {
		panic(err)
	}
}
