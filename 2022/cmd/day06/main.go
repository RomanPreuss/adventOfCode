package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("Day 06")

	file, err := os.Open("lvl1.txt")
	defer func() { _ = file.Close() }()
	if err != nil {
		log.Fatal(err)
	}

	root := reconstructFileSystem(file)
	smallDirectories := FindDirectoriesWithMaxSize(root, 100000)

	smallDirectoriesSizeSum := 0
	for _, d := range smallDirectories {
		smallDirectoriesSizeSum = smallDirectoriesSizeSum + d.Size
	}

	folderToBeDeleted := FindDirToDelete(root, 70000000, 30000000)

	log.Info("1. sum of small directories: ", smallDirectoriesSizeSum)
	log.Infof("2. delete folder: %v with size: %v", folderToBeDeleted.Name, folderToBeDeleted.Size)
}

type Node struct {
	IsFolder bool
	Name     string
	Size     int
	Children []*Node
	Parent   *Node
}

func (n *Node) Add(node *Node) {
	n.Children = append(n.Children, node)
	node.Parent = n
}

func FindDirToDelete(root *Node, maxFileSystemSize, neededSpace int) *Node {
	freeSpace := maxFileSystemSize - root.Size
	spaceToBeFreed := neededSpace - freeSpace
	log.Info("free space: ", freeSpace, " space to be freed: ", spaceToBeFreed)

	folderToBeDeleted := root
	allSubFolders := root.Children
	for i := 0; i < len(allSubFolders); i++ {
		subFolder := allSubFolders[i]
		if !subFolder.IsFolder {
			continue
		}

		allSubFolders = append(allSubFolders, subFolder.Children...)
		if subFolder.Size < folderToBeDeleted.Size && subFolder.Size >= spaceToBeFreed {
			folderToBeDeleted = subFolder
		}
	}

	return folderToBeDeleted
}

func FindDirectoriesWithMaxSize(node *Node, maxSize int) []*Node {
	if !node.IsFolder {
		node.Parent.Size = node.Parent.Size + node.Size
		return []*Node{}
	}

	subDirectories := []*Node{}
	for _, c := range node.Children {
		subDirectories = append(subDirectories, FindDirectoriesWithMaxSize(c, maxSize)...)
	}

	if node.Parent != nil {
		node.Parent.Size = node.Parent.Size + node.Size
	}

	if node.Size <= maxSize {
		subDirectories = append(subDirectories, node)
	}

	return subDirectories
}

func reconstructFileSystem(data io.Reader) *Node {
	scanner := bufio.NewScanner(data)
	root := Node{
		IsFolder: true,
		Name:     "/",
	}
	currentDir := &root
	lsMode := false

outer:
	for scanner.Scan() {
		line := scanner.Text()
		log.Debugf("current dir: %v process: [%v] lsMode: [%v]\n", currentDir.Name, line, lsMode)

		// switch to ls mode
		if line == "$ ls" {
			log.Debugf("\tenable ls mode\n")
			lsMode = true
			continue
		}

		// navigate to root folder
		if line == "$ cd /" {
			log.Debugf("\tnavigate to root\n")
			currentDir = &root
			lsMode = false
			continue
		}

		// navigate to sub folder (create if not existing)
		if subFolderName, found := strings.CutPrefix(line, "$ cd "); found {
			log.Debugf("\tnavigate to: '%v'\n", subFolderName)
			lsMode = false

			// navigate to parent
			if subFolderName == ".." {
				if currentDir.Parent == nil {
					log.Panic("cannot navigate to empty parent")
				}
				currentDir = currentDir.Parent
				continue
			}

			// check if sub folder already exists
			for _, c := range currentDir.Children {
				if c.IsFolder && c.Name == subFolderName {
					currentDir = c
					continue outer
				}
			}
			log.Panic("folder ", subFolderName, " doesn't exist")
		}

		if lsMode {
			// sub folder doesn't exist yet - create it
			if subFolderName, found := strings.CutPrefix(line, "dir "); found {
				log.Debugf("\tcreating sub folder '%v'\n", subFolderName)
				subFolder := Node{
					IsFolder: true,
					Name:     subFolderName,
				}
				currentDir.Add(&subFolder)
				continue
			}

			fileInfo := strings.Split(line, " ")
			fileSize, err := strconv.Atoi(fileInfo[0])
			if err != nil {
				log.Panic(err)
			}
			file := Node{
				IsFolder: false,
				Name:     fileInfo[1],
				Size:     fileSize,
			}
			log.Debugf("\tcreating file '%v'\n", file.Name)
			currentDir.Add(&file)
		}
	}
	return &root
}
