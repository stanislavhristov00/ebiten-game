package space

import (
	"fmt"
)

const (
	NUM_ENEMIES = 55
)

type State struct {
	enemies []Enemy
	player  Player
	input   Input
}

func NewState() *State {
	state := &State{
		enemies: make([]Enemy, NUM_ENEMIES),
		input:   *NewInput(),
	}

	return state
}

func (st *State) LoadEnemies(enemies []Enemy) {
	if len(st.enemies)+len(enemies) < NUM_ENEMIES {
		st.enemies = append(st.enemies, enemies...)
	} else {
		fmt.Printf("Cannot load enemies. Exceeding capacity of NUM_ENEMIES: %d\n", NUM_ENEMIES)
	}
}

func (st *State) LoadPlayer(player *Player) {
	st.player = *player
}

func (st State) UpdateDirMovement(isFullScreen bool) (Dir, bool) {
	return st.input.Dir(isFullScreen)
}

func (st State) UpdateShoot(isFullScreen bool) bool {
	return st.input.Update(isFullScreen)
}
