package bridge

import "fmt"

type mac struct {
	printer iPrinter
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	m.printer.printFile()
}

func (m *mac) setPrinter(p iPrinter) {
	m.printer = p
}
