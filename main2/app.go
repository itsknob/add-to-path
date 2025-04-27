package main

import (
	"errors"
	"os"
	"slices"
	"strings"

	"github.com/itsknob/hawk-tui/controller"
	"github.com/rivo/tview"
	"github.com/sahilm/fuzzy"
)

var (
	app           *tview.Application
	pages         *tview.Pages
	list          *tview.List
	selectedItem  *string
	addStartMenu  *tview.Form
	addEndMenu    *tview.Form
	removeMenu    *tview.Form
	fuzzyFindMenu *tview.Form
	appFlex       *tview.Flex
	path          []string
	filteredPath  []string
)

func main() {

	controller.New().ListPathEntries()
	pc.Start()

	//////////////////////
	// Main Flex Container
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.SetBorder(true).SetTitle("Hawk Tui")

	/////////////
	// Entry List
	list = tview.NewList().
		// List Props
		ShowSecondaryText(false).
		SetHighlightFullLine(true)

	list.
		// Box Props
		SetBorder(true).
		SetTitle("Path").
		SetTitleAlign(tview.AlignLeft)

	updateList := func(path []string) {
		if path == nil {
			path = strings.Split(os.Getenv("PATH"), ":")
		}
		for idx, item := range path {
			list.AddItem(item, "", int32(idx), func() {
				selectedItem = &item
			})
		}
	}

	// ////////////////////
	// // Add To Start Menu
	// addStartMenu = tview.NewForm()
	// addStartMenu.SetBorder(true).SetTitle("Add to Start").SetTitleAlign(tview.AlignLeft)
	// addStartMenu.AddInputField("Dir", "", 0, nil, nil)
	// addStartMenu.AddButton("Back", func() {
	//     app.SetRoot(flex, true).SetFocus(fuzzyFindMenu)
	//     app.Draw()
	// })

	// ////////////////////
	// // Add To End Menu
	// addEndMenu = tview.NewForm()
	// addEndMenu.SetBorder(true).SetTitle("Add to End").SetTitleAlign(tview.AlignLeft)
	// addEndMenu.AddInputField("Dir", "", 0, nil, nil)
	// addEndMenu.AddButton("Back", func() {
	//     app.SetRoot(flex, true).SetFocus(fuzzyFindMenu)
	//     app.Draw()
	// })

	// ////////////////////
	// // Remove Entry Menu
	// removeMenu = tview.NewForm()

	// ////////////////////
	// // Fuzzy Find Menu
	// fuzzyFindMenu = tview.NewForm()
	// fuzzyFindMenu.SetBorder(true).SetTitle("Search").SetTitleAlign(tview.AlignLeft)
	// fuzzyFindMenu.AddInputField("Input", "", 0, nil, nil)
	// fuzzyFindMenu.AddButton(
	// 	"Add to Path",
	// 	func() {
	//     input := fuzzyFindMenu.GetFormItemByLabel("Input").(*tview.InputField)
	//     text := input.GetText()
	//     path, err := AddToPathBack(text)
	//     if err != nil {
	//         app.SetFocus(
	//             tview.NewModal().
	//                 AddButtons([]string{"OK"}).
	//                 SetText(err.Error()).
	//                 SetBorder(true).SetTitle("Error"))
	//             }
	//     println(path)
	// })

	// fuzzyFindMenu.AddButton("Remove from Path", func() {
	//     input := fuzzyFindMenu.GetFormItemByLabel("Input").(*tview.InputField)
	//     text := input.GetText()
	//     path, err := RemoveFromPath(text)
	//     if err != nil {
	//         panic(err)
	//     }
	//     println(path)
	// })

	// fuzzyFindMenu.AddButton("Quit", func() {
	//     app.Stop()
	// })

	// flex.AddItem(list, 0, 3, false)
	// flex.AddItem(fuzzyFindMenu, 0, 1, true)

	// pc := controller.PathController{
	//     Path: model.Path{},
	//     View: ui.State{
	//         Root: app,
	//         Pages: pages,
	//         Focused: fuzzyFindMenu,
	//     },
	// }

	pc.View.Start()

	// if err := app.SetRoot(flex, true).SetFocus(fuzzyFindMenu).Run(); err != nil {
	//     panic(err)
	// }

}

func AddToPathBack(dir string) ([]string, error) {
	_, err := os.Stat(dir)
	if err != nil {
		return path, err
	}

	path = append(path, dir)
	err = os.Setenv("PATH", GetPathAsString())

	// os.Execute("export PATH=%s", GetPathAsString())

	if err != nil {
		// println(err)
		return path, err
	}

	return path, nil
}

func RemoveFromPath(dir string) ([]string, error) {
	foundIdx := slices.Index(path, dir)
	if foundIdx+1 >= len(path) {
		// log.Default().Output(1, fmt.Sprintf("foundIdx+1 out of bounds - foundIdx+1: %d, len(path): %d", foundIdx+1, len(path)))
		return path, errors.New("Out of bound during delete")
	}
	path = slices.Delete(path, foundIdx, foundIdx+1)

	return path, nil
}

func GetPathAsString() string {
	return strings.Join(path, ":") // global
}
