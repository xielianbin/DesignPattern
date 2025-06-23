package internal

import "fmt"

// 具体原型
type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + f.Name)
}

func (f *File) Clone() INode {
	return &File{Name: f.Name + "_clone"}
}
