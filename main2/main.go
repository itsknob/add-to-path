package main2

/**

import (
	"github.com/itsknob/hawk-tui/envpath"
	"github.com/itsknob/hawk-tui/ui"
	"github.com/rivo/tview"
)

func main() {

    state := ui.State{}
	// Container with Builtins
	root := tview.NewApplication()

    // Get p from env and set up entries
    // Path Controller
    p := envpath.Path{}
    p.Init() // get PATH from env

    state.Path = &p

    // List of items in path
    list := state.CreateList(root)
    // Add entries from path to list
    ui.AddDataToList(list, p.Entries)

    // Fuzzy Find Menu
    fuzzyFindMenu := state.CreateFuzzyFindMenu()

    // Remove menus
    // removeFromPath := state.CreateRemoveFromPathMenu(root, p)

    // Add to Path Menus
    addToPathFrontMenu := state.CreateAddToPathMenu(true)
    addToPathBackMenu := state.CreateAddToPathMenu(false)

    // Main Menu
    mainMenu := state.CreateMainMenu(root, addToPathFrontMenu, addToPathBackMenu)

	// Container for List and Selected Text Area
	pathFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(list, 0, 1, true)

	// Container for Path Entries and Menus
	entriesAndMainMenu := tview.NewFlex().
		AddItem(pathFlex, 0, 1, false).
		AddItem(mainMenu, 0, 1, true)

    entriesAndAddFormFront := tview.NewFlex().
        AddItem(pathFlex, 0, 1, false).
        AddItem(addToPathFrontMenu, 0, 1, true)

    entriesAndAddFormBack := tview.NewFlex().
        AddItem(pathFlex, 0, 1, false).
        AddItem(addToPathBackMenu, 0, 1, true)

    entriesAndFuzzyFind := tview.NewFlex().
        AddItem(pathFlex, 0, 1, false).
        AddItem(fuzzyFindMenu, 0, 1, true)


    pages := tview.NewPages()
    pages.AddPage("Main Menu", entriesAndMainMenu, true, true)
    pages.AddPage("Add to Path Front", entriesAndAddFormFront, true, false)
    pages.AddPage("Add to Path Back", entriesAndAddFormBack, true, false)
    pages.AddPage("Fuzzy Find", entriesAndFuzzyFind, true, false)
    // pages.AddPage("Delete from Path", entriesAndDeleteForm, true, false)
    state.Pages = pages

    state.Focused = root.GetFocus()

	// Render App, with Flex as Root, and Focus on List
	if err := root.SetRoot(pages, true).SetFocus(mainMenu).Run(); err != nil {
		panic(err)
	}
}
**/
