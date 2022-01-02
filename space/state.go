package space

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	NUM_ROWS = 5
)

var (
	NUM_ENEMIES_ON_ROW       = 0
	ENEMY_MOVEMENT_DIRECTION = 1
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
		enemies:    make([]*Enemy, 0),
		numEnemies: numEnemies,
		input:      *NewInput(),
	}

	NUM_ENEMIES_ON_ROW = state.numEnemies / NUM_ROWS
	return state
}

func (st *State) LoadEnemies(enemies []*Enemy) {
	// 	if len(st.enemies)+len(enemies) <= st.numEnemies {
	// 		for i := 0; i < len(enemies); i++ {
	// 			st.enemies = append(st.enemies, enemies[i])
	// 		}
	// 	} else {
	// 		fmt.Printf("Cannot load enemies. Exceeding capacity of NUM_ENEMIES: %d\n", st.numEnemies)
	// 	}
	// }
	for i := 0; i < len(enemies); i++ {
		st.enemies = append(st.enemies, enemies[i])
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
		} else if !st.enemies[i].GetDeathFrameDrawn() {
			st.enemies[i].DieDraw(screen)
			st.enemies[i].SetDeathFrameDrawn(true)
		}
	}
}

func (st State) DrawPlayer(screen *ebiten.Image) {
	if st.player.IsAlive() {
		st.player.Draw(screen)
	} else {
		st.player.DieDraw(screen)
	}
}

func (st State) CheckIfPlayerShotEnemy() {
	for i := 0; i < len(st.enemies); i++ {
		if st.enemies[i].IsAlive() {
			if st.player.BulletCollisionWithEnemy(st.enemies[i]) {
				st.enemies[i].Die()
				st.player.SetBulletInAir(false)
			}
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

func (st State) MoveEnemies(screenWidth int) {
	x0, _ := st.enemies[0].GetEnemyXY()
	scaleX0, _ := st.enemies[0].GetScaleXY()
	xRow, _ := st.enemies[NUM_ENEMIES_ON_ROW-1].GetEnemyXY()
	xRowScaleX, _ := st.enemies[NUM_ENEMIES_ON_ROW-1].GetScaleXY()
	xRowWidth := st.enemies[NUM_ENEMIES_ON_ROW-1].GetFrameWidth()
	if float64(x0)*scaleX0 < 0 || float64(xRow)*xRowScaleX+float64(xRowWidth)*xRowScaleX > float64(screenWidth) {
		ENEMY_MOVEMENT_DIRECTION *= -1
	}

	for i := 0; i < st.numEnemies; i++ {
		st.enemies[i].OffsetXY(ENEMY_MOVEMENT_DIRECTION*2, 0)
	}

}

func (st State) MovePlayer(x, y int) {
	st.player.OffsetXY(x, y)
}

func (st State) PlayerShoot() {
	st.player.Shoot()
}

func (st State) EnemyShoot(index int) {
	if index < len(st.enemies) {
		st.enemies[index].Shoot()
	} else {
		fmt.Printf("Index out of bounds with index %d", index)
	}
}

func (st *State) CopyEnemiesIntoState(state *State) {
	for i := 0; i < len(state.enemies); i++ {
		*st.enemies[i] = *state.enemies[i]
	}
}
