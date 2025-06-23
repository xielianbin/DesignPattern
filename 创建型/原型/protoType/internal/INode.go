package internal

// 原型接口
type INode interface {
	Print(string)
	Clone() INode
}
