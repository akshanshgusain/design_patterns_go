package bridge

import "fmt"

type windows struct {
	printer iPrinter
}

func (w *windows) print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *windows) setPrinter(p iPrinter) {
	w.printer = p
}
