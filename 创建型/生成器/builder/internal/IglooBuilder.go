package internal

type IglooBuilder struct {
	windowType string
	doorType   string
	numFloors  int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}
func (b *IglooBuilder) SetWindowType() {
	b.windowType = "igloo window"
}
func (b *IglooBuilder) SetDoorType() {
	b.doorType = "igloo door"
}
func (b *IglooBuilder) SetNumFloors() {
	b.numFloors = 4
}
func (b *IglooBuilder) GetHouse() House {
	return House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		NumFloors:  b.numFloors,
	}
}
