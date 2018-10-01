package game

import "fmt"

// PlayerColleague
type PlayerColleague interface {
	GetName() string
	GetFleet() []Ship
	GetLastGunShootCoordinate() *Coordinate
	GunShoot(c Coordinate)
}

type Player struct {
	// name of player
	name string
	// playerFleet stores all available fleet for player
	fleet []Ship
	// battleField
	battleField *BattleField
	// room where player joined
	room WarRoomMediator
	// lastFireCoordinate place where was gun shoot
	lastFireCoordinate *Coordinate
}

// NewPlayer
func NewPlayer(name string, fleet []Ship, warRoom WarRoomMediator) (p *Player, err error) {
	p = &Player{
		name:        name,
		battleField: newBattleGround(10, 10),
		room:        warRoom,
	}

	if len(fleet) == 0 {
		return nil, fmt.Errorf("player fleet cannot be empty")
	}

	for _, ship := range fleet {
		if len(ship.Location) == 0 {
			return nil, fmt.Errorf("one of the ship in fleet has empty location")
		}
	}

	if isCorrect := p.battleField.ValidateFleetCollision(fleet); isCorrect != nil {
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
func (p *Player) GetFleet() []Ship {
	return p.fleet
}

// GunShoot
func (p *Player) GunShoot(target Coordinate) {
	p.lastFireCoordinate = &target

	p.room.MakeTurn(p)
}

func (p *Player) GetLastGunShootCoordinate() *Coordinate {
	return p.lastFireCoordinate
}