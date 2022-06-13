package builder

// normalBuilder.go: Concrete builder

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumberOfFloors() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
