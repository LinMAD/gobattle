package generator

import (
	"github.com/LinMAD/gobattle/pkg/game"
	"sync"
)

// NewFleet generate whole fleet with random locations
func NewFleet() []*game.Ship {
	fleet := make([]*game.Ship, 0)
	mux := &sync.Mutex{}

	shipCount := len(game.ShipTypes)

	var wg sync.WaitGroup
	wg.Add(shipCount)

	for worker := 1; worker <= shipCount; worker++ {
		go func(worker int) {
			defer wg.Done()

			for {
				newShip := createShip(uint8(worker))

				// Check if ship not collided
				mux.Lock()
				tmpFleet := make([]*game.Ship, len(fleet))
				copy(tmpFleet, fleet)
				mux.Unlock()
				tmpFleet = append(tmpFleet, newShip)
				err := game.ValidateFleetCollision(tmpFleet)
				if err == nil {
					mux.Lock()
					fleet = append(fleet, newShip)
					mux.Unlock()
					break
				}
			}
		}(worker)
	}

	wg.Wait()

	return fleet
}

// createShip randomly generate coordinates of ship
func createShip(shipSize uint8) *game.Ship {
	shipCoordinate := make([]game.Coordinate, 0)

	if shipSize == 1 {
		shipCoordinate = append(
			shipCoordinate,
			game.Coordinate{
				AxisX: RandomNum(0, int(game.FSize-1)),
				AxisY: RandomNum(0, int(game.FSize-1)),
			},
		)

		return &game.Ship{IsAlive: true, Location: shipCoordinate}
	}

	// create horizontal or vertical ship
	isShipVertical := RandomBool()
	// to keep connected coordinates of ship just add to each rand num size
	x := RandomNum(0, int(game.FSize)-int(shipSize))
	y := RandomNum(0, int(game.FSize)-int(shipSize))
	var size uint8
	for size = 0; size < shipSize; size++ {
		if isShipVertical {
			shipCoordinate = append(
				shipCoordinate,
				game.Coordinate{AxisX: x + int8(size), AxisY: y},
			)
			continue
		}

		shipCoordinate = append(
			shipCoordinate,
			game.Coordinate{AxisX: x, AxisY: y + int8(size)},
		)
	}

	return &game.Ship{IsAlive: true, Location: shipCoordinate}
}
