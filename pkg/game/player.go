package game

import "fmt"

// PlayerColleague
type PlayerColleague interface {
	GetName() string
	GetFleet() []Ship
	GetBattleField() *BattleField
}

// Player represents size of playing field in x,y coordinate axis
type Player struct {
	// name of player
	name string
	// playerFleet stores all available fleet for player
	playerFleet []Ship
	// battleField
	battleField *BattleField
}

// NewPlayer
func NewPlayer(name string, fleet []Ship) (p *Player, err error) {
	p = &Player{name: name}

	if len(fleet) == 0 {
		return nil, fmt.Errorf("player fleet cannot be empty")
	}

	// TODO Validate fleet position, if they can be placed by game rules
	p.playerFleet = fleet

	return p, nil
}

// GetName of player
func (p *Player) GetName() string {
	return p.name
}

// GetFleet of valid ships
func (p *Player) GetFleet() []Ship {
	return p.playerFleet
}

// GetBattleField
func (p *Player) GetBattleField() *BattleField {
	return p.battleField
}
