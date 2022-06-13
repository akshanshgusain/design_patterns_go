package builder

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumberOfFloors()
	getHouse() House
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
