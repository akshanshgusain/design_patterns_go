package abstractFactory

import "fmt"

type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidasFactory" {
		return &adidasFactory{}, nil
	}
	if brand == "nikeFactory" {
		return &nikeFactory{}, nil
	}
	return nil, fmt.Errorf("Invalid Brand Type")
}
