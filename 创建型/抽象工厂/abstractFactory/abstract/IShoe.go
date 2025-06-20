package abstract

type IShoe interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
	SetInfo(info string)
	GetInfo() string
}
