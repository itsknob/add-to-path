package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/itsknob/hawk-tui/model"
	"github.com/rivo/tview"
	"github.com/sahilm/fuzzy"
)

type State struct {
    Root *tview.Application
    Pages *tview.Pages
    Focused *tview.Primitive
    List *tview.List
}

func New() (*State) {
    app := tview.NewApplication()
    pages := tview.NewPages()

    list := tview.NewList().ShowSecondaryText(false).SetHighlightFullLine(true)
    list.SetBorder(true).SetTitle("Path").SetTitleAlign(tview.AlignLeft)

    state := &State{
        Root: app,
        Pages: pages,
        Focused: nil,
        List: list,
    }

    return state
}

func createAddToStartMenu() (*tview.Form) {
    ////////////////////
    // Add To Start Menu
    addStartMenu := tview.NewForm()
    addStartMenu.SetBorder(true).SetTitle("Add to Start").SetTitleAlign(tview.AlignLeft)
    addStartMenu.AddInputField("Dir", "", 0, nil, nil)
    // addStartMenu.AddButton("Back", func() {
    //     app.SetRoot(flex, true).SetFocus(fuzzyFindMenu)
    //     app.Draw()
    // })
    
    return addStartMenu
}

func createAddToEndMenu() (*tview.Form) {
    ////////////////////
    // Add To End Menu
    addEndMenu := tview.NewForm()
    addEndMenu.SetBorder(true).SetTitle("Add to End").SetTitleAlign(tview.AlignLeft)
    addEndMenu.AddInputField("Dir", "", 0, nil, nil)
    // addEndMenu.AddButton("Back", func() {
    //     app.SetRoot(flex, true).SetFocus(fuzzyFindMenu)
    //     app.Draw()
    // })

    return addEndMenu
}

func createRemoveMenu() (*tview.Form) {
    ////////////////////
    // Remove Entry Menu
    removeMenu := tview.NewForm()

    return removeMenu
}

func createFuzzyFindMenu(app *tview.Application) (*tview.Form) {
    ////////////////////
    // Fuzzy Find Menu
    fuzzyFindMenu := tview.NewForm()
    fuzzyFindMenu.SetBorder(true).SetTitle("Search").SetTitleAlign(tview.AlignLeft)
    fuzzyFindMenu.AddInputField("Input", "", 0, nil, nil)
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

    return fuzzyFindMenu
}

func (state *State) Start(path *model.Path) {
    // Initialize Path List
    state.AddDataToList(path.Entries)

    // Create Menu Forms
    fuzzyFindMenu := createFuzzyFindMenu(state.Root)
    addToStartMenu := createAddToStartMenu()
    addToEndMenu := createAddToEndMenu()
    removeFromPathMenu := createRemoveMenu()

    //////////////////////
    // Main Flex Container
    mainMenuContainer := tview.NewFlex().SetDirection(tview.FlexRow)
    mainMenuContainer.SetBorder(true).SetTitle("Hawk Tui")
    mainMenuContainer.AddItem(state.List, 0, 3, false)
    mainMenuContainer.AddItem(fuzzyFindMenu, 0, 1, true)
    // Main Menu Page
    state.Pages = state.Pages.AddAndSwitchToPage("Main Menu", mainMenuContainer, true)


    addToStartContainer := tview.NewFlex().SetDirection(tview.FlexRow)
    addToStartContainer.SetBorder(true).SetTitle("Hawk Tui")
    addToStartContainer.AddItem(state.List, 0, 3, false)
    addToStartContainer.AddItem(addToStartMenu, 0, 1, true)
    state.Pages = state.Pages.AddPage("Add to Path Front", addToStartContainer, true, false)

    addToEndContainer := tview.NewFlex().SetDirection(tview.FlexRow)
    addToEndContainer.SetBorder(true).SetTitle("Hawk Tui")
    addToEndContainer.AddItem(state.List, 0, 3, false)
    addToEndContainer.AddItem(addToEndMenu, 0, 1, true)
    state.Pages = state.Pages.AddPage("Add to Path Back", addToEndContainer, true, false)

    removeFromPathContainer := tview.NewFlex().SetDirection(tview.FlexRow)
    removeFromPathContainer.SetBorder(true).SetTitle("Hawk Tui")
    removeFromPathContainer.AddItem(state.List, 0, 3, false)
    removeFromPathContainer.AddItem(removeFromPathMenu, 0, 1, true)
    state.Pages = state.Pages.AddPage("Remove From Path", removeFromPathContainer, true, false)

    if err := state.Root.SetRoot(mainMenuContainer, true).SetFocus(fuzzyFindMenu).Run(); err != nil {
        panic(err)
    }
}

func (state *State) showMainMenu() {
    state.Pages.SwitchToPage("Main Menu")
}

func (state *State) showAddToPathFrontMenu() {
    state.Pages.SwitchToPage("Add to Path Front")
}
func (state *State) showAddToPathBackMenu() {
    state.Pages.SwitchToPage("Add to Path Back")
}
func (state *State) showRemoveFromPathMenu() {
    state.Pages.SwitchToPage("Remove From Path")
}

func (state *State) AddDataToList(data []string) {
    // for each entry, print to TextView
	for idx, path := range []string(data) {
        state.List.AddItem(path, "", int32(idx+95), nil)
	}
}

func (state *State) CreateMainMenu(root *tview.Application, addToPathFrontMenu *tview.Form, addToPathBackMenu *tview.Form) (*tview.List) {
    main := tview.NewList().ShowSecondaryText(false)
    main.SetTitle("What would you like to do?").SetBorder(true)

    main.
    AddItem("Add item to PATH", "Start of Path: $PATH=/new/dir:$PATH", 'a', state.showAddToPathFrontMenu).
    AddItem("Add item to PATH", "End of Path: $PATH=$PATH:/new/dir", 'b', state.showAddToPathBackMenu).
    AddItem("Remove item from PATH", "", 'c', func() {
        // println("Selected (c) - Remove item from PATH")
        // main.SetFocus(list.View)
    }).
    AddItem("Search for item in PATH", "Fuzzy Find", 'd', func() {
        // println("Selected (d) - Search for item in PATH")
        state.Pages.SwitchToPage("Fuzzy Find")
        // root.SetFocus(list.View)
    }).
    AddItem("Quit", "No further changes will be made", 'q', func() {
        root.Stop()
    })

    return main
}

func (state *State) CreateAddToPathMenu(pc *controller.PathController, addToFront bool) (*tview.Form) {
    menu := tview.NewForm()
    menu.SetTitle("Add to Path, e.g. /usr/bin/local")
    menu.SetBorder(true)

    menu.AddInputField("Directory", "", 0, nil, nil)

    inputField := menu.GetFormItemByLabel("Directory").(*tview.InputField)
    inputField.SetDoneFunc(func (key tcell.Key) {

        // println("KEY", key)
        switch key {
            case tcell.KeyEnter: {
                if (addToFront) {
                    path, err := pc.Path.AddToPathFront(inputField.GetText())
                    if err != nil {
                        // println(err.Error())
                        // popup? maybe a warning message somehwhere somehow?
                        // modal.SetText(err.Error())
                        // state.Root.SetRoot(modal, false).SetFocus(modal)
                    }
                    // println("Updated PATH=", path)
                } else {
                    path, _ := pc.Path.AddToPathFront(inputField.GetText())
                    // if err != nil {
                    //     panic(err)
                    // }
                    // println("Updated PATH=", path)
                }
            }
        }

		//       modal := tview.NewModal()
		//       modal.SetTitle("Error").SetBorder(true)
		//       modal.AddButtons([]string{"Quit", "Cancel"})
		// modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		// 	if buttonLabel == "Quit" {
		// 		state.Root.Stop()
		//               state.Pages.SwitchToPage("Main Menu")
		// 	}
		// })

    })

    inputField.SetAutocompleteFunc(func (currentText string) (entries []string) {
            matches := fuzzy.Find(inputField.GetText(), pc.Path.Entries)
            entries = []string{}
            for _, match := range matches {
                entries = append(entries, match.Str)
            }
            return entries
        },
    )

    menu.AddButton("Menu", state.showMainMenu)

    return menu
}

func (state *State) CreateFuzzyFindMenu(pathItems []string) (*tview.Form) {
    menu := *tview.NewForm()
    menu.SetTitle("Fuzzy Find").SetBorder(true)
    menu.AddInputField("Search", "", 0, nil, nil)

    inputField := menu.GetFormItemByLabel("Search")
    in := menu.GetFormItemByLabel("Search").(*tview.InputField)

    inputField.SetFinishedFunc(func(key tcell.Key) {
        switch key {
            case tcell.KeyEsc: {
                in.SetText("")
                break;
            }
            case tcell.KeyEnter: {
                // TODO: Copy to clipboard?
                // println("Found:", in.GetText())
                break;
            }
            default: {
                state.Pages.SwitchToPage("Main Menu")
                break;
            }
        }


    })
    in.SetAutocompleteFunc(func (currentText string) (entries []string) {
            matches := fuzzy.Find(in.GetText(), pathItems)
            entries = []string{}
            for _, match := range matches {
                entries = append(entries, match.Str)
            }
            return entries
        },
    )

    return &menu
    
}

func (state *State) CreateRemoveFromPathMenu(pc *controller.PathController, pathItems []string) (*tview.Form) {
    menu := *tview.NewForm()
    menu.SetTitle("Remove from Path").SetBorder(true)
    menu.AddInputField("Directory", "", 0, nil, nil)

    inputField := menu.GetFormItemByLabel("Directory")
    inputField.SetFinishedFunc(func(key tcell.Key) {
        in := inputField.(*tview.InputField)
        switch key {
            case tcell.KeyEsc: {
                in.SetText("")
                break;
            }
            case tcell.KeyEnter: {
                // TODO: Copy to clipboard?
                println("Found:", in.GetText())
                pc.Path.RemoveFromPath(in.GetText())
                state.showMainMenu()
                break;
            }
            default: {
                state.Pages.SwitchToPage("Main Menu")
                break;
            }
        }

    })


    return &menu
}
