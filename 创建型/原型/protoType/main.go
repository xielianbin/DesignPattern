package main

import (
	"fmt"
	"protoType/internal"
)

func main() {
	file1 := &internal.File{Name: "File1"}
	file2 := &internal.File{Name: "File2"}
	file3 := &internal.File{Name: "File3"}

	folder1 := &internal.Folder{
		Children: []internal.INode{file1},
		Name:     "Folder1",
	}

	folder2 := &internal.Folder{
		Children: []internal.INode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
