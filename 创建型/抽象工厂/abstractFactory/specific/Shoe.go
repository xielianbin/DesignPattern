package specific

type Shoe struct {
	logo string
	size int
	info string
}

func (s *Shoe) SetLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) GetLogo() string {
	return s.logo
}

func (s *Shoe) SetSize(size int) {
	s.size = size
}

func (s *Shoe) GetSize() int {
	return s.size
}
func (s *Shoe) GetInfo() string {
	return s.info
}
func (s *Shoe) SetInfo(info string) {
	s.info = info
}
