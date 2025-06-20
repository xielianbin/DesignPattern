package abstract

type IGun interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
	SetRange(scope int)
	GetRange() int
}
