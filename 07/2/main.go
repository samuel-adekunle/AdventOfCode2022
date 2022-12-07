package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Dir struct {
	Name    string
	Path    string
	Files   []*File
	SubDirs []*Dir
	Size    int
}

func (d *Dir) GetAllSubDirs() []*Dir {
	subDirs := []*Dir{}
	subDirs = append(subDirs, d.SubDirs...)
	for _, subDir := range d.SubDirs {
		subDirs = append(subDirs, subDir.GetAllSubDirs()...)
	}
	return subDirs
}

func (d *Dir) TotalSize() int {
	size := d.Size

	subDirs := d.GetAllSubDirs()
	for _, subDir := range subDirs {
		size += subDir.Size
	}

	return size
}

func pathString(path []*Dir) string {
	pathStr := ""
	for _, dir := range path {
		pathStr += dir.Name + "/"
	}
	return pathStr
}

func main() {
	fileBytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(fileBytes), "\n")
	dirs := make(map[string]map[string]*Dir) // path -> dirName -> dir
	path := []*Dir{{Name: "root"}}

	var currentDir *Dir
	currentPath := pathString(path)
	for _, line := range lines {
		if line[0] == '$' {
			command := line[2:4]
			if command == "cd" {
				dirName := line[5:]
				if dirName == ".." {
					path = path[:len(path)-1]
					currentPath = pathString(path)
					currentDir = path[len(path)-1]
				} else {
					if _, ok := dirs[currentPath]; !ok {
						dirs[currentPath] = make(map[string]*Dir)
					}
					if _, ok := dirs[currentPath][dirName]; !ok {
						dirs[currentPath][dirName] = &Dir{Name: dirName}
					}
					currentDir = dirs[currentPath][dirName]
					path = append(path, currentDir)
					currentPath = pathString(path)
				}
			}
			continue
		}
		if line[0:3] == "dir" {
			dirName := line[4:]
			if _, ok := dirs[currentPath]; !ok {
				dirs[currentPath] = make(map[string]*Dir)
			}
			if _, ok := dirs[currentPath][dirName]; !ok {
				dirs[currentPath][dirName] = &Dir{Name: dirName}
			}
			currentDir.SubDirs = append(currentDir.SubDirs, dirs[currentPath][dirName])
			continue
		}
		// file
		fileDetails := strings.Split(line, " ")
		fileSize, err := strconv.ParseInt(fileDetails[0], 10, 64)
		if err != nil {
			panic(err)
		}
		fileName := fileDetails[1]
		file := &File{Name: fileName, Size: int(fileSize)}
		currentDir.Files = append(currentDir.Files, file)
		currentDir.Size += int(fileSize)
	}

	targetSize := 30000000
	currentSize := dirs["root/"]["/"].TotalSize()
	fmt.Println("current size:", currentSize)

	for pathName, path := range dirs {
		for _, dir := range path {
			fmt.Println(pathName, dir.Name, dir.TotalSize())
		}
	}

	var dirToDelete *Dir
	var dirToDeleteSize int
	// find smallest dir to delete to get under target size
	for _, path := range dirs {
		for _, dir := range path {
			dirSize := dir.TotalSize()
			if currentSize-dirSize <= targetSize {
				if dirToDelete == nil || dirSize < dirToDeleteSize {
					dirToDelete = dir
					dirToDeleteSize = dirSize
				}
			}
		}
	}

	fmt.Println("dir to delete:", dirToDelete.Name, dirToDeleteSize)
}
