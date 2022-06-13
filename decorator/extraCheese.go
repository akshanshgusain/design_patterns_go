package decorator

import "fmt"

type extraCheese struct {
	pizza iPizza
}

func (e *extraCheese) getPrice() int {
	pizzaPrice := e.pizza.getPrice()
	fmt.Println("Adding extra cheese to the Pizza!")
	return pizzaPrice + 50
}
