package factoryMethod

import "fmt"

// Factory Method is a creation design pattern that provides an interface for creating objects in a superclass,
// but allows subclasses to alter the type of objects that will be created.

//https://refactoring.guru/design-patterns/factory-method/go/example#example-0

//It’s impossible to implement the classic Factory Method pattern in Go due to lack of OOP features such as classes and
//inheritance. However, we can still implement the basic version of the pattern, the Simple Factory.

// Client Code

func FactoryMethod() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g iGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
