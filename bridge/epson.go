package bridge

import "fmt"

type epson struct {
}

func (e *epson) printFile() {
	fmt.Println("Printing by a EPSON Printer")
}
