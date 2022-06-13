package bridge

import "fmt"

//Bridge is a structural design pattern that divides business logic or huge class into separate class hierarchies that
//can be developed independently.

// Example
/*
Say, you have two types of computers: Mac and Windows. Also, two types of printers: Epson and HP. Both computers and
printers need to work with each other in any combination. The client doesn’t want to worry about the details of
connecting printers to computers.

If we introduce new printers, we don’t want our code to grow exponentially. Instead of creating four structs for the
2*2 combination, we create two hierarchies:

Abstraction hierarchy: this will be our computers
Implementation hierarchy: this will be our printers
These two hierarchies communicate with each other via a Bridge, where the Abstraction (computer) contains a reference
to the Implementation (printer). Both the abstraction and implementation can be developed independently without
affecting each other.
*/

func Bridge() {
	epsonPrinter := new(epson)
	hpPrinter := new(hp)

	macComputer := new(mac)
	windowsComputer := new(windows)

	//Mac Computer and printers
	macComputer.setPrinter(epsonPrinter)
	macComputer.print()
	fmt.Println()

	macComputer.setPrinter(hpPrinter)
	macComputer.print()
	fmt.Println()

	//Windows Computer and Printers
	windowsComputer.setPrinter(epsonPrinter)
	windowsComputer.print()
	fmt.Println()

	windowsComputer.setPrinter(hpPrinter)
	windowsComputer.print()
	fmt.Println()

}
