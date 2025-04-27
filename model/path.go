package model

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Path struct {
    Entries []string
}

func New() (*Path){
    foundPath := os.Getenv("PATH")
    if foundPath == "" {
        panic("PATH was empty")
    }

    entries := strings.Split(foundPath, ":")
    slices.Sort(entries)

    p := &Path{
        Entries: entries,
    }

    return p
}

func (path *Path) GetPathEntries() {

    foundPath := os.Getenv("PATH")
    if foundPath == "" {
        panic("PATH was empty")
    }

    entries := strings.Split(foundPath, ":")
    slices.Sort(entries)
    path.Entries = entries
}

func (path *Path) GetPathAsString() (string) {
    return strings.Join(path.Entries, ":")
}

func (path *Path) AddToPathFront(dir string) ([]string, error) {
    // Check that new path is valid
    // println("Adding ", dir, " to front of PATH")
    _, err := os.Stat(dir)
    if err != nil {
        return path.Entries, err
    }

    path.Entries = append([]string{dir}, path.Entries...)
    
    return path.Entries, nil
}

func (path *Path) AddToPathBack(dir string) ([]string, error) {
    _, err := os.Stat(dir)
    if err != nil {
        return path.Entries, err
    }

    path.Entries = append(path.Entries, dir)

    return path.Entries, nil
}

func (path *Path) RemoveFromPath(dir string) (string, error) {
    foundIdx := slices.Index(path.Entries, dir)
    if foundIdx+1 >= len(path.Entries) {
        log.Default().Output(1, fmt.Sprintf("foundIdx+1 out of bounds - foundIdx+1: %d, len(path): %d", foundIdx+1, len(path.Entries)))
        return path.GetPathAsString(), errors.New("Out of bound during delete")
    }
    path.Entries = slices.Delete(path.Entries, foundIdx, foundIdx+1)

    return path.GetPathAsString(), nil 
}

func UpdateZProfile() {}

func UpdateBashProfile() {}

func UpdateZshRc() {}
