package abstractFactory

import "fmt"

// Abstract Factory is a creational design pattern that lets you produce families of related objects without specifying
// their concrete classes.

//https://refactoring.guru/design-patterns/abstract-factory/go/example#example-0

// Abstract Factory defines an interface for creating all distinct products but leaves the actual product creation to
// concrete factory classes. Each factory type corresponds to a certain product variety.

func AbstractFactoryPattern() {
	adidasFactory, _ := getSportsFactory("adidasFactory")
	nikeFactory, _ := getSportsFactory("nikeFactory")

	adidasShirt := adidasFactory.makeShirt()
	adidasShoe := adidasFactory.makeShoe()

	nikeShirt := nikeFactory.makeShirt()
	nikeShoe := nikeFactory.makeShoe()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShirt) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
