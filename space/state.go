package space

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

/*
 *	Represents current game state
 */
type State struct {
	enemies    []*Enemy
	player     *Player
	input      Input
	numEnemies int
}

func NewState(numEnemies int) *State {
	state := &State{
		enemies:    make([]*Enemy, numEnemies),
		numEnemies: numEnemies,
		input:      *NewInput(),
	}

	return state
}

func (st *State) LoadEnemies(enemies []*Enemy) {
	if len(st.enemies)+len(enemies) < st.numEnemies {
		st.enemies = append(st.enemies, enemies...)
	} else {
		fmt.Printf("Cannot load enemies. Exceeding capacity of NUM_ENEMIES: %d\n", st.numEnemies)
	}
}

func (st *State) LoadPlayer(player *Player) {
	st.player = player
}

func (st State) UpdateDirMovement() (Dir, bool) {
	return st.input.Dir()
}

func (st State) UpdateShoot() bool {
	return st.input.Update()
}

func (st State) DrawEnemies(screen *ebiten.Image, count int) {
	for i := 0; i < len(st.enemies); i++ {
		if st.enemies[i].IsAlive() {
			st.enemies[i].Draw(screen, count)
		} else {
			st.enemies[i].DieDraw(screen)
			removeElement(st.enemies, i)
		}
	}
}

func (st State) DrawPlayer(screen *ebiten.Image) {
	if st.player.IsAlive() {
		st.player.Draw(screen)
	}
}

func (st State) CheckIfPlayerShotEnemy() {
	for i := 0; i < len(st.enemies); i++ {
		if st.player.BulletCollisionWithEnemy(st.enemies[i]) {
			st.enemies[i].Die()
			st.player.SetBulletInAir(false)
		}
	}
}

func (st State) CheckIfEnemyShotPlayer() {
	for i := 0; i < len(st.enemies); i++ {
		if st.enemies[i].BulletCollisionWithPlayer(st.player) {
			st.player.LoseLife()
			st.enemies[i].SetBulletInAir(false)
		}
	}
}

func removeElement(enemies []*Enemy, index int) {
	enemies[index] = enemies[len(enemies)-1]
	enemies[len(enemies)-1] = &Enemy{}
	enemies = enemies[:len(enemies)-1]
}
