package decorator

import "fmt"

type extraSalami struct {
	pizza iPizza
}

func (e *extraSalami) getPrice() int {
	pizzaPrice := e.pizza.getPrice()
	fmt.Println("Adding extra salami to the Pizza!")
	return pizzaPrice + 100
}
