package game

// PlayerColleague
type PlayerColleague interface {
	GetName() string
	GetShips() []Ship
	// TODO load correct ships in map to player
}

// Player represents size of playing field in x,y coordinate axis
type Player struct {
	// playerShips stores all available fleet for player
	playerShips []Ship
	// name of player
	name string
}

// NewPlayer
func NewPlayer(name string) *Player {
	return &Player{
		name: name,
	}
}

// GetName of player
func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetShips() []Ship {
	return p.playerShips
}
