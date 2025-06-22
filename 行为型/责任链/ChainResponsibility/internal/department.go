package internal
type Department interface{
	Execute(* Patient)
	SetNext(Department)
}