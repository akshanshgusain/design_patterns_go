package abstractFactory

// Adidas Concrete factory

type adidasFactory struct {
}

func (a *adidasFactory) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidasFactory",
			size: 14,
		},
	}
}

func (a *adidasFactory) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "nikeFactory",
			size: 42,
		},
	}
}
