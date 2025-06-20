package abstract

// 抽象工厂，接口要在名字前面加大写的I。首字母大写是公有变量，小写是私有变量

type ISportsFactory interface {
	MakeShoe() IShoe
	MakeShirt() IShirt
}
