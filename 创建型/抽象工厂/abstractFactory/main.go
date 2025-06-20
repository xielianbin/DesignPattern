package main

import (
	"abstractFactory/abstract"
	"abstractFactory/specific"
	"fmt"
)

func GetSportsFactory(brand string) (abstract.ISportsFactory, error) {
	if brand == "adidas" {
		return &specific.AdidasFactory{}, nil
	}

	if brand == "nike" {
		return &specific.NikeFactory{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
func printShoeDetails(s abstract.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
	fmt.Printf("Info: %s", s.GetInfo())
	fmt.Println()
}

func printShirtDetails(s abstract.IShirt) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

}
