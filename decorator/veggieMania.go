package decorator

import "fmt"

// Concrete Component

type veggieMania struct {
}

func (v *veggieMania) getPrice() int {
	fmt.Println("Creating Base Pizza!")
	return 300
}
