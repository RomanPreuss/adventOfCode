package main

import (
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFillTree(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

	t.Run("build correct tree", func(t *testing.T) {

		tree := reconstructFileSystem(strings.NewReader(input))

		assert.Equal(t, "/", tree.Name)
		assert.Len(t, tree.Children, 4)

		a, d, e := getFolders(tree)

		assert.Equal(t, "a", a.Name)
		assert.Len(t, a.Children, 4)

		assert.Equal(t, "d", d.Name)
		assert.Len(t, d.Children, 4)

		assert.Equal(t, "e", e.Name)
		assert.Len(t, e.Children, 1)
	})

	t.Run("find folder with max size", func(t *testing.T) {

		tree := reconstructFileSystem(strings.NewReader(input))

		smallDirectories := FindDirectoriesWithMaxSize(tree, 100000)

		assert.Len(t, smallDirectories, 2)
		sum := 0
		for _, d := range smallDirectories {
			sum = sum + d.Size
		}
		assert.Equal(t, 95437, sum)
	})

	t.Run("find correct folder to delete", func(t *testing.T) {
		log.SetLevel(log.DebugLevel)
		root := reconstructFileSystem(strings.NewReader(input))

		FindDirectoriesWithMaxSize(root, 30000000)
		log.Info("Size: ", root.Size)

		folderToBeDeleted := FindDirToDelete(root, 70000000, 30000000)
		assert.Equal(t, "d", folderToBeDeleted.Name)
	})
}

func getFolders(tree *Node) (*Node, *Node, *Node) {
	var a *Node
	var d *Node
	var e *Node
	allChildren := tree.Children
	for i := 0; i < len(allChildren); i++ {
		n := allChildren[i]
		if n.IsFolder {
			allChildren = append(allChildren, n.Children...)
		}

		if n.Name == "a" {
			a = n
		}
		if n.Name == "d" {
			d = n
		}
		if n.Name == "e" {
			e = n
		}
	}
	return a, d, e
}
