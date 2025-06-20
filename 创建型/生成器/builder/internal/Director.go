package internal

type Director struct {
	builder IBuilder
}

func NewDirector(builder IBuilder) *Director {
	return &Director{builder: builder}
}

func (d *Director) SetBuilder(builder IBuilder) {
	d.builder = builder
}
func (d *Director) BuildHouse() House {
	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloors()
	return d.builder.GetHouse()
}
