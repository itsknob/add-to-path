package main

import (
	"errors"
	"os"
	"slices"
	"strings"
)

func addToPathFront(dir string) (string, error) {
    path := os.Getenv("PATH")
    if path == "" {
        return path, errors.New("PATH not found")
    }

    _, err := os.Stat(dir)
    if err != nil {
        return path, err
    }

    path = strings.Join([]string {dir, path}, ":")
    return path, nil
}

func addToPathBack(dir string) (string, error) {
    path := os.Getenv("PATH")
    if path == "" {
        return path, errors.New("PATH not found")
    }

    _, err := os.Stat(dir)
    if err != nil {
        return path, err
    }

    path = strings.Join([]string {path, dir}, ":")
    return path, nil
}

func removeFromPath(dir string) (string, error) {
    path := os.Getenv("PATH")
    if path == "" {
        return path, errors.New("PATH not found")
    }

    pathEntries := strings.Split(path, ":")
    foundIdx := slices.Index(pathEntries, dir)

    updatedEntries := slices.Delete(pathEntries, foundIdx, foundIdx+1)

    newPath := strings.Join(updatedEntries, ":")

    return newPath, nil 
}

func updateZProfile() {}

func updateBashProfile() {}

func updateZshRc() {}
