package abstractFactory

type nikeFactory struct {
}

func (n *nikeFactory) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nikeFactory",
			size: 8,
		},
	}
}

func (n *nikeFactory) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nikeFactory",
			size: 42,
		},
	}
}
