package flyweight

// Context

type player struct {
	dress      iDress
	playerType string
	lat        int
	long       int
}

func newPlayer(playerType string, dressType string) *player {
	dress, _ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *player) newLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
