package specific

type Gun struct {
	name  string
	power int
	scope int
}

func (g *Gun) SetName(name string) {
	g.name = name
}
func (g *Gun) SetPower(power int) {
	g.power = power
}
func (g *Gun) GetName() string {
	return g.name
}
func (g *Gun) GetPower() int {
	return g.power
}
func (g *Gun) SetRange(scope int) {
	g.scope = scope
}
func (g *Gun) GetRange() int {
	return g.scope
}
