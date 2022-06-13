package factoryMethod

// Concrete Product

type musket struct {
	gun
}

// Concrete product

func newMusket() iGun {
	return &musket{
		gun{
			name:  "Musket",
			power: 1,
		},
	}
}
