package main

import (
	"composite/internal"
)

func main() {
	file1 := &internal.File{Name: "File1"}
	file2 := &internal.File{Name: "File2"}
	file3 := &internal.File{Name: "File3"}

	folder1 := &internal.Folder{
		Name: "Folder1",
	}

	folder1.Add(file1)

	folder2 := &internal.Folder{
		Name: "Folder2",
	}
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
