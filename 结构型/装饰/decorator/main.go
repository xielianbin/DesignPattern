package main

import (
	"decorator/internal"
	"fmt"
)

func main() {

	pizza := &internal.VeggieMania{}

	//Add cheese topping
	pizzaWithCheese := &internal.CheeseTopping{
		Pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &internal.TomatoTopping{
		Pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.GetPrice())
}
