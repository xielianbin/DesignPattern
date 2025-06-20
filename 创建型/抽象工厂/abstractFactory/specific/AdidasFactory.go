package specific

import "abstractFactory/abstract"

type AdidasFactory struct {
}

func (a *AdidasFactory) MakeShoe() abstract.IShoe {
	s := Shoe{}
	s.SetLogo("adidas")
	s.SetSize(32)
	s.SetInfo("Adidas-shoe")
	return &AdidasShoe{
		Shoe: s,
	}
}

func (a *AdidasFactory) MakeShirt() abstract.IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			Logo: "adidas",
			Size: 14,
		},
	}
}
