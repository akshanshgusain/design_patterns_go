package flyweight

import "fmt"

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerroristDressType terrorist dress type
	CounterTerroristDressType = "ctDress"
)

var dressFactorySingleInstance = &dressFactory{
	dressMap: make(map[string]iDress),
}

type dressFactory struct {
	dressMap map[string]iDress
}

func (d *dressFactory) getDressByType(dressType string) (iDress, error) {
	if d.dressMap[dressType] != nil {
		return d.dressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}

	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
