package game

import "fmt"

// PlayerColleague
type PlayerColleague interface {
	GetName() string
	GetFleet() []*Ship
	GetLastGunShootCoordinate() *Coordinate
	GunShoot(c Coordinate) bool
}

// Player
type Player struct {
	// isActive player
	isActive bool
	// name of player
	name string
	// playerFleet stores all available fleet for player
	fleet []*Ship
	// room where player joined
	room WarRoomMediator
	// lastFireCoordinate place where was gun shoot
	lastFireCoordinate *Coordinate
}

// NewPlayer
func NewPlayer(name string, fleet []*Ship, warRoom WarRoomMediator) (p *Player, err error) {
	p = &Player{
		name: name,
		room: warRoom,
	}

	if len(fleet) == 0 {
		return nil, fmt.Errorf("player fleet cannot be empty")
	}
	if len(fleet) > len(ShipTypes) {
		return nil, fmt.Errorf("player fleet should be not greater than -> %d", len(ShipTypes))
	}

	isExpectedLenOfShip := func(lenOfShip int) bool {
		for _, expected := range ShipTypes {
			if expected == uint8(lenOfShip) {
				return true
			}
		}

		return false
	}

	for _, ship := range fleet {
		lenOfShip := len(ship.Location)
		if lenOfShip == 0 {
			return nil, fmt.Errorf("one of the ship in fleet has empty location")
		}

		if !isExpectedLenOfShip(lenOfShip) {
			return nil,
				fmt.Errorf(
					"not allowed size of the ship (%d), fleet must have each of %v",
					lenOfShip,
					ShipTypes,
				)
		}
	}

	if isCorrect := ValidateFleetCollision(fleet); isCorrect != nil {
		return nil, isCorrect
	}

	p.fleet = fleet
	p.room.JoinPlayer(p)

	return p, nil
}

// GetName of player
func (p *Player) GetName() string {
	return p.name
}

// GetFleet of valid ships
func (p *Player) GetFleet() []*Ship {
	return p.fleet
}

// GunShoot try hit target with given coordinates
func (p *Player) GunShoot(target Coordinate) bool {
	p.lastFireCoordinate = &target

	return p.room.MakeTurn(p)
}

// GetLastGunShootCoordinate return last coordinate of firing
func (p *Player) GetLastGunShootCoordinate() *Coordinate {
	return p.lastFireCoordinate
}
