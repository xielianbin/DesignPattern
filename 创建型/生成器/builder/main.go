package main

import (
	"builder/internal"
	"fmt"
)

func getBuilder(builderType string) internal.IBuilder {
	// if builderType == "normal" {
	//     return newNormalBuilder()
	// }

	if builderType == "igloo" {
		return internal.NewIglooBuilder()
	}
	return nil
}
func main() {
	// normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := internal.NewDirector(iglooBuilder)
	// normalHouse := director.buildHouse()

	// fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	// fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	// fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.NumFloors)

}
