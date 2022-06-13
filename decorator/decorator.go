package decorator

import "fmt"

// Decorator is a structural pattern that allows adding new behaviors to objects dynamically
// by placing them inside special wrapper objects.

// Using decorators you can wrap objects countless number of times since both target objects
// and decorators follow the same interface. The resulting object will get a stacking behavior
//of all wrappers.

func Decorator() {
	myPizza := &veggieMania{}

	//add salami
	myPizzaWithSalami := &extraSalami{pizza: myPizza}

	//add extra cheese
	myPizzaWithCheese := &extraCheese{
		myPizzaWithSalami,
	}

	fmt.Printf("Price of veggeMania with salami and cheese topping is %d\n", myPizzaWithCheese.getPrice())
}
