package flyweight

import "fmt"

// Flyweight is a structural design pattern that allows programs to support vast quantities of objects by keeping their
// memory consumption low.
// The pattern achieves it by sharing parts of object state between multiple objects.
// In other words, the Flyweight saves RAM by caching the same data used by different objects.

/// Conceptual Example

/*
In a game of Counter-Strike, Terrorist and Counter-Terrorist have a different type of dress. For simplicity, let’s
assume that both Terrorist and Counter-Terrorists have one dress type each. The dress object is embedded in the player
object as below. Below is the struct for a player. We can see that the dress object is embedded in the player struct:

type player struct {
    dress      dress
    playerType string // Can be T or CT
    lat        int
    long       int
}

Let’s say there are 5 Terrorists and 5 Counter-Terrorist, so a total of 10 players. Now there are two options concerning dress.
Each of the 10 player objects creates a different dress object and embeds them. A total of 10 dress objects will be created.

We create two dress objects:
1. Single Terrorist Dress Object: This will be shared across 5 Terrorists.
2. Single Counter-Terrorist Dress Object: This will be shared across 5 Counter-Terrorist.

As you can see in Approach 1, a total of 10 dress objects are created while in approach 2 only 2 dress objects are
created. The second approach is what we follow in the Flyweight design pattern. The two dress objects which we created
are called the flyweight objects.

The Flyweight pattern takes out the common parts and creates flyweight objects. These flyweight objects (dress)
can then be shared among multiple objects (player). This drastically reduces the number of dress objects, and the
good part is that even if you create more players, only two dress objects will be sufficient.

In the flyweight pattern, we store the flyweight objects in the map field. Whenever the other objects which share the
flyweight objects are created, then flyweight objects are fetched from the map.

*/

func Flyweight() {
	game := newGame()

	//Add Terrorists
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)
	game.addTerrorist(TerroristDressType)

	// Add Counter-Terrorists
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}

}
