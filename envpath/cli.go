package envpath

import (
	"os"
	"slices"
	"strings"
)

type Path struct {
    Entries []string
}
            
func (p *Path) Init() {

    path := os.Getenv("PATH")
    if path == "" {
        panic("PATH was empty")
    }

    // split
    entries := strings.Split(path, ":")
    // sort
    slices.Sort(entries)
    // save
    p.Entries = entries

}

func (p *Path) GetPathAsString() (string) {
    return strings.Join(p.Entries, ":")
}

func (p *Path) AddToPathFront(dir string) (error) {
    // Check that new path is valid
    // println("Adding ", dir, " to front of PATH")
    _, err := os.Stat(dir)
    if err != nil {
        return err
    }

    p.Entries = append([]string{dir}, p.Entries...)
    
    return nil
}

func (p *Path) AddToPathBack(dir string) (error) {
    _, err := os.Stat(dir)
    if err != nil {
        return err
    }

    p.Entries = append(p.Entries, dir)

    return nil
}

func (p *Path) RemoveFromPath(dir string) (error) {
    foundIdx := slices.Index(p.Entries, dir)
    p.Entries = slices.Delete(p.Entries, foundIdx, foundIdx+1)

    return nil 
}

func (p *Path) UpdateZProfile() {}

func (p *Path) UpdateBashProfile() {}

func (p *Path) UpdateZshRc() {}

