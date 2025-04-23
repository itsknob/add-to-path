package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Init() {

    foundPath := os.Getenv("PATH")
    if foundPath == "" {
        panic("PATH was empty")
    }

    // split
    entries := strings.Split(foundPath, ":")

    // sort
    slices.Sort(entries)
    // save
    path = entries // global

}

func GetPathAsString() (string) {
    return strings.Join(path, ":") // global
}



func AddToPathFront(dir string) (string, error) {
    // Check that new path is valid
    println("Adding ", dir, " to front of PATH")
    _, err := os.Stat(dir)
    if err != nil {
        return GetPathAsString(), err
    }

    path = append([]string{dir}, path...)
    
    return GetPathAsString(), nil
}

// func AddToPathBack(dir string) ([]string, error) {
//     _, err := os.Stat(dir)
//     if err != nil {
//         return path, err
//     }
//
//     path = append(path, dir)
//
//     return path, nil
// }

func RemoveFromPath(dir string) ([]string, error) {
    foundIdx := slices.Index(path, dir)
    if foundIdx+1 >= len(path) {
        log.Default().Output(1, fmt.Sprintf("foundIdx+1 out of bounds - foundIdx+1: %d, len(path): %d", foundIdx+1, len(path)))
        return path, errors.New("Out of bound during delete")
    }
    path = slices.Delete(path, foundIdx, foundIdx+1)

    return path, nil 
}

func UpdateZProfile() {}

func UpdateBashProfile() {}

func UpdateZshRc() {}

