package main

import (
	"fmt"
	"os/exec"

	"github.com/gdamore/tcell/v2"
	"github.com/itsknob/hawk-tui/envpath"
	"github.com/rivo/tview"
)

type Controller struct {
    app *tview.Application
    form *tview.Form
    list *tview.List
    path *envpath.Path
}

func onSubmit(form *tview.Form, inputLabel string) string {
	inputFormItem := form.GetFormItemByLabel(inputLabel).(*tview.InputField)
	return inputFormItem.GetText()
}

func (pc *Controller) refreshList(entries []string) {
	pc.list.Clear()

	for idx, item := range entries {
		pc.list.AddItem(item, "", int32(idx+97), nil)
	}
}

func (pc *Controller) switchToForm(newFocus *tview.Form) {
	// todo will explode
	pc.app.SetFocus(pc.form)
}

func (pc *Controller) switchToList(newFocus *tview.List) {
	pc.app.SetFocus(pc.list)
}

func (pc *Controller) setDirText(selectedText string) {
	formInput := pc.form.GetFormItemByLabel("Directory").(*tview.InputField)
	formInput.SetText(selectedText)
}

func (pc *Controller) NewList() {
	pc.list = tview.NewList().ShowSecondaryText(false)
	pc.list.SetBorder(true).SetTitle("Path")

	for idx, item := range pc.path.Entries {
		pc.list.AddItem(item, "", int32(idx+97), func() {
			mainText, _ := pc.list.GetItemText(idx)
			pc.setDirText(mainText)
			pc.switchToForm(pc.form)
		})
	}
    pc.list.AddItem("Quit", "Save changes to PATH", 'Q', func() {
        pc.app.Stop()
    })
}

func (pc *Controller) NewForm() {
	pc.form = tview.NewForm().
		AddInputField("Directory", "", 0, nil, nil)
	pc.form.
		AddButton("Add to Front", func() {
			value := pc.form.GetFormItemByLabel("Directory").(*tview.InputField).GetText()
			if err := pc.path.AddToPathFront(value); err != nil {
				panic(err)
			}
			pc.refreshList(pc.path.Entries)
		}).
		AddButton("Add to Back", func() {
			value := pc.form.GetFormItemByLabel("Directory").(*tview.InputField).GetText()
			if err := pc.path.AddToPathBack(value); err != nil {
				panic(err)
			}
			pc.refreshList(pc.path.Entries)
		}).
		AddButton("Remove from Path", func() {
			value := pc.form.GetFormItemByLabel("Directory").(*tview.InputField).GetText()
			if err := pc.path.RemoveFromPath(value); err != nil {
				panic(err)
			}
			pc.refreshList(pc.path.Entries)
		}).
        AddButton("Save and Quit", func() {
            newPath := fmt.Sprintf("PATH=%s", pc.path.GetPathAsString())
            exec.Command("export",  newPath).Run()
        })

	pc.form.SetBorder(true).SetTitle("Operation Go!")
}

func main() {
    pathController := Controller{
        app: tview.NewApplication(),
        path: &envpath.Path{},
        list: &tview.List{},
        form: tview.NewForm(),
    }
    // Initialize
    pathController.path.Init()
    pathController.NewList()
    pathController.NewForm()

    container := tview.NewFlex()
    container.SetDirection(tview.FlexRow)
    container.AddItem(pathController.list, 0, 3, true).AddItem(pathController.form, 0, 2, false)

	container.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
        // Toggle between panes with ESC
		case tcell.KeyESC:
			{
				if pathController.list.HasFocus() {
					pathController.switchToForm(pathController.form)
				} else {
					pathController.switchToList(pathController.list)
				}
			}
		default:
			break
		}
		return event
	})

	if err := pathController.app.SetRoot(container, true).SetFocus(container).Run(); err != nil {
		panic(err)
	}
    
}

