package specific

import "abstractFactory/abstract"

type NikeFactory struct {
}

func (n *NikeFactory) MakeShoe() abstract.IShoe {
	s := Shoe{}
	s.SetLogo("nike")
	s.SetSize(42)
	s.SetInfo("Nike-shoe")
	return &NikeShoe{
		Shoe: s,
	}
}

func (n *NikeFactory) MakeShirt() abstract.IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
