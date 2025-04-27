package controller

import (
	"github.com/itsknob/hawk-tui/model"
	"github.com/itsknob/hawk-tui/ui"
)

type PathController struct {
    Path *model.Path
    View *ui.State
}

func New() (*PathController) {
    path := model.New()
    ui := ui.State{}
    app = tview.NewApplication()
    pc := &PathController{
        Path: path,
        View: ui,
    }
    return pc
}

func (pc *PathController) Start() (error) {
    // Setup View?
    pc.View.Start()

    // Show Main Menu Page
    return nil
}

func (pc *PathController) ListPathEntries() {
    // model.GetResource
    entries := 
    // view.ShowResource
    pc.View.AddDataToList(entries)
}
