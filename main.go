package main

import (
	"fmt"
	"slices"
	// "strconv"

	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func setupList(app *tview.Application) (*tview.TextView) {
    list := tview.NewTextView()
    list.SetBorder(true)
    list.SetTitle("Current Path Entries")
    list.SetInputCapture(func(event *tcell.EventKey) (*tcell.EventKey) {
        switch event.Key() {
            case tcell.KeyESC: {
                focusList(app, list)
                break;
            }
            default: {
                break;
            }
        }
        return event
    })

    // list.SetSelectedFunc(func(i int, s1, s2 string, r rune) {
    //     *selected = s1
    //     // fmt.Println("\n\nSelected updated to: ", *selected)
    // })

    return list
}

func focusList(app *tview.Application, list *tview.TextView) {
    app.SetFocus(list)
}

func main() {
    pathDirectories := getPathAsList()
    slices.Sort(pathDirectories)

    // Container with Builtins
    app := tview.NewApplication()

    // List of every item on the path
    // Pass by reference
    list := setupList(app)

    // Shows the currently selected option
    selectedTextArea := tview.NewTextView().
        SetDynamicColors(true).
        SetRegions(true).
        SetWordWrap(true)

    // A function to clear and replace selected text
    // refreshSelectedText := func(selected string) {
    //     selectedTextArea.Clear()
    //     // Print to the selectedTextArea directly
    //     fmt.Fprintln(selectedTextArea, selected)
    // }

    /////////////////////
    /// Add items to List
    /////////////////////
    // Add individual path items to list
    for _, path := range []string(pathDirectories) {
        // start runes at lowercase 'a' offset by idx
        // ascii := strconv.Itoa(idx+97) // offset and convert to string
        // r, _ := strconv.Atoi(ascii) // convert string back to int

        
        fmt.Fprintln(list, path)
        // list.AddItem(path, "", rune(r), func() {
        //     selected := pathDirectories[list.GetCurrentItem()]
        //     refreshSelectedText(selected)
        // })
    }

    // Add option to quit after list of path directories
    // list.AddItem("Quit", "Press to exit", 'Q', func() {
    //     fmt.Println("Selected Q, Stopping App")
    //     app.Stop()
    // })

    ///////////////
    // Menu Form //
    ///////////////
    createMenu := func () (*tview.List) {
        menu := tview.NewList()
        menu.SetTitle("What would you like to do?")
        menu.SetBorder(true)

        menu.
        AddItem("Add item to PATH", "Add to start of PATH", 'a', func () {
            fmt.Println("Selected (a) - Add item to PATH (start)")
            app.Draw()
        }).
        AddItem("Add item to PATH", "Add to end of PATH", 'b', func () {
            fmt.Println("Selected (b) - Add item to PATH (end)")
        }).
        AddItem("Remove item from PATH", "", 'c', func(){
            fmt.Println("Selected (c) - Remove item from PATH")
            app.SetFocus(list)
        }).
        AddItem("Search for item in PATH", "Fuzzy find in path", 'd', func(){
            fmt.Println("Selected (d) - Search for item in PATH")
            app.SetFocus(list)
        }).
        AddItem("Quit", "No further changes will be made", 'q', func(){
            app.Stop()
        })
        return menu
    }

    mainMenu := createMenu()
    ///////////////////////////
    // Add Panes to Flex Layout
    ///////////////////////////
    // Container for List and Selected Text Area
    pathFlex := tview.NewFlex().SetDirection(tview.FlexRow).
        AddItem(list, 0, 1, true).
        AddItem(selectedTextArea, 0, 2, false)


    // Container for Path Entries and Menus
    flex := tview.NewFlex().
        AddItem(pathFlex, 0, 2, false).
        AddItem(mainMenu, 0, 1, true)
        // AddItem(tview.NewBox().SetBorder(true).SetTitle(string(selected)), 0, 1, false).

    
    // Render App, with Flex as Root, and Focus on List
    if err := app.SetRoot(flex, true).SetFocus(mainMenu).Run(); err != nil {
        panic(err)
    }
}

func getPathAsList() ([]string) {
    pathString := getPath();
    // Split path on ":"
    pathList := strings.Split(pathString, ":")
    return pathList
}

// Get $PATH as a string
func getPath() (string) {
    path := os.Getenv("PATH")
    cmd := exec.Command("echo", path)
    stdout, err := cmd.Output()

    if err != nil {
        panic(err)
    }
    return string(stdout)

}
