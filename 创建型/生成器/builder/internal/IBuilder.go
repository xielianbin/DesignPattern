package internal

type IBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloors()
	GetHouse() House
}
