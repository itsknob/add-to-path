package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/itsknob/hawk-tui/envpath"
	"github.com/rivo/tview"
	"github.com/sahilm/fuzzy"
)

type State struct {
    Root *tview.Application
    Pages *tview.Pages
    Path *envpath.Path
    Focused tview.Primitive
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

func (state *State) CreateList(root *tview.Application) (*tview.TextView){
    l := tview.NewTextView()
	l.SetBorder(true)
	l.SetTitle("Current Path Entries")
	l.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyESC:
			{
				state.showMainMenu()
				break
			}
		default:
			{
				break
			}
		}
		return event
	})
    return l
}

func AddDataToList(l *tview.TextView, data []string) {
    for each entry, print to TextView
	for _, path := range []string(data) {
		// fmt.Fprintln(l, path)
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

func (state *State) CreateAddToPathMenu(addToFront bool) (*tview.Form) {
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
                    path, err := state.Path.AddToPathFront(inputField.GetText())
                    if err != nil {
                        // println(err.Error())
                        // popup? maybe a warning message somehwhere somehow?
                        // modal.SetText(err.Error())
                        // state.Root.SetRoot(modal, false).SetFocus(modal)
                    }
                    // println("Updated PATH=", path)
                } else {
                    path, _ := state.Path.AddToPathBack(inputField.GetText())
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
            matches := fuzzy.Find(inputField.GetText(), state.Path.Entries)
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

func (state *State) CreateFuzzyFindMenu() (*tview.Form) {
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
            matches := fuzzy.Find(in.GetText(), state.Path.Entries)
            entries = []string{}
            for _, match := range matches {
                entries = append(entries, match.Str)
            }
            return entries
        },
    )

    return &menu
    
}

func (state *State) CreateRemoveFromPathMenu() (*tview.Form) {
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
                // println("Found:", in.GetText())
                state.Path.RemoveFromPath(in.GetText())
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
