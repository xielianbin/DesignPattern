package internal

import "fmt"

// 具体原型
type Folder struct {
	Children []INode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)
	for _, i := range f.Children {
		i.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() INode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []INode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}
	cloneFolder.Children = tempChildren
	return cloneFolder
}
