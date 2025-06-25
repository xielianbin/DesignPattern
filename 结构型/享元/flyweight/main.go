package main

import (
	"flyweight/internal"
	"fmt"
)

func main() {
	game := internal.NewGame()

	//Add Terrorist
	game.AddTerrorist(internal.TerroristDressType)
	game.AddTerrorist(internal.TerroristDressType)
	game.AddTerrorist(internal.TerroristDressType)
	game.AddTerrorist(internal.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(internal.CounterTerrroristDressType)
	game.AddCounterTerrorist(internal.CounterTerrroristDressType)
	game.AddCounterTerrorist(internal.CounterTerrroristDressType)

	dressFactoryInstance := internal.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
