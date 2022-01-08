package space

import (
	"fmt"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	NUM_ROWS = 5
)

var (
	NUM_ENEMIES_ON_ROW       = 0
	ENEMY_MOVEMENT_DIRECTION = 1
	ENEMY_SPEED              = 2
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

/*
 *	Creates a state object.
 */

func NewState(numEnemies int) *State {
	state := &State{
		enemies:    make([]*Enemy, 0),
		numEnemies: numEnemies,
		input:      *NewInput(),
	}

	NUM_ENEMIES_ON_ROW = state.numEnemies / NUM_ROWS
	return state
}

/*
 *	Load enemies into state.
 */
func (st *State) LoadEnemies(enemies []*Enemy) {
	for i := 0; i < len(enemies); i++ {
		st.enemies = append(st.enemies, enemies[i])
	}
}

/*
 *	Load player into state.
 */

func (st *State) LoadPlayer(player *Player) {
	st.player = player
}

/*
 *	Used to check for player input about direction.
 */

func (st State) UpdateDirMovement() (Dir, bool) {
	return st.input.Dir()
}

/*
 *	Used to check for player input about shooting.
 */

func (st State) UpdateShoot() bool {
	return st.input.Update()
}

/*
 *	Draw the enemies onto the screen passed. Count argument is used
 *	for animation.
 */

func (st State) DrawEnemies(screen *ebiten.Image, count int) {
	for i := 0; i < len(st.enemies); i++ {
		st.enemies[i].Draw(screen, count)

		if !st.enemies[i].IsAlive() && !st.enemies[i].GetDeathFrameDrawn() {
			st.enemies[i].DieDraw(screen)
			st.enemies[i].SetDeathFrameDrawn(true)
		}
	}
}

/*
 *	Returns true if all current enemies are dead, false otherwise.
 */

func (st State) CheckIfAllEnemiesAreDead() bool {
	for i := 0; i < len(st.enemies); i++ {
		if st.enemies[i].IsAlive() {
			return false
		}
	}

	return true
}

/*
 *	Draw the player on the screen passed.
 */

func (st State) DrawPlayer(screen *ebiten.Image) {
	if st.player.IsAlive() {
		st.player.Draw(screen)
	} else {
		st.player.DieDraw(screen)
	}
}

/*
 *	Returns true if the player's bullet has collided with any of the enemies, false otherwise.
 */

func (st State) CheckIfPlayerShotEnemy() bool {
	for i := 0; i < len(st.enemies); i++ {
		if st.enemies[i].IsAlive() {
			if st.player.BulletCollisionWithEnemy(st.enemies[i]) {
				st.player.SetBulletInAir(false)
				st.enemies[i].Die()
				return true
			}
		}
	}

	return false
}

/*
 *	Returns true if any of the enemies' bullets has collided with the player, false otherwise.
 */

func (st State) CheckIfEnemyShotPlayer() {
	for i := 0; i < len(st.enemies); i++ {
		if st.enemies[i].BulletCollisionWithPlayer(st.player) {
			st.player.LoseLife()
			st.enemies[i].SetBulletInAir(false)
		}
	}
}

/*
 *	Moves the enemies horizontally and vertically. Also checks if any of the enemies
 *	collide with the player. If so, the player's state is set to dead.
 */

func (st State) MoveEnemies(screenWidth int) {
	x0, _ := st.enemies[0].GetEnemyXY()
	scaleX0, _ := st.enemies[0].GetScaleXY()
	xRow, _ := st.enemies[NUM_ENEMIES_ON_ROW-1].GetEnemyXY()
	xRowScaleX, _ := st.enemies[NUM_ENEMIES_ON_ROW-1].GetScaleXY()
	xRowWidth := st.enemies[NUM_ENEMIES_ON_ROW-1].GetFrameWidth()

	/*
	 *	Change direction if any of the enemies have reached the end.
	 *	Also move each enemy one row down.
	 */
	if float64(x0)*scaleX0 < 0 || float64(xRow)*xRowScaleX+float64(xRowWidth)*xRowScaleX > float64(screenWidth) {
		ENEMY_MOVEMENT_DIRECTION *= -1
		for i := 0; i < st.numEnemies; i++ {
			st.enemies[i].OffsetXY(0, st.enemies[i].GetFrameHeight())
		}
	}

	for i := 0; i < st.numEnemies; i++ {
		st.enemies[i].OffsetXY(ENEMY_MOVEMENT_DIRECTION*ENEMY_SPEED, 0)

		_, scaleY := st.enemies[i].GetScaleXY()

		if st.enemies[i].IsAlive() && int(float64(st.enemies[i].posY)*scaleY) > st.player.posY {
			st.player.Die()
		}
	}
}

/*
 *	Makes random enemies shoot. Number of shooting enemies is dependent on
 *	amount of currently alive enemies.
 */

func (st *State) EnemiesShoot() {

	for i := 0; i < len(st.enemies); i++ {
		if st.enemies[i].IsAlive() && st.enemies[i].IsBulletInAir() {
			return
		}
	}

	var numberOfEnemiesToShoot int
	alive := make([]int, 0)

	for i := 0; i < len(st.enemies); i++ {
		if st.enemies[i].IsAlive() {
			alive = append(alive, i)
		}
	}

	if len(alive) == 0 {
		return
	}

	if len(alive) <= 5 {
		for i := 0; i < len(alive); i++ {
			st.EnemyShoot(alive[i])
		}
		return
	}

	if len(alive) > 20 {
		numberOfEnemiesToShoot = 5
	} else {
		numberOfEnemiesToShoot = 3
	}

	for i := 0; i < numberOfEnemiesToShoot; i++ {
		index := pickARandomNumberInRange(0, len(alive)-1)

		if st.enemies[index].IsBulletInAir() {
			for st.enemies[index].IsBulletInAir() {
				index = pickARandomNumberInRange(0, len(alive)-1)
			}
		}

		st.EnemyShoot(index)
	}
}

/*
 *	Move player horizontally.
 */

func (st State) MovePlayer(x, y, limit int) {
	x1, _ := st.player.GetPlayerXY()
	scaleX, _ := st.player.GetScaleXY()

	if x1+x < 0 || int(float64(x1)*scaleX)+int(float64(st.player.GetFrameWidth())*scaleX)+x-5 > limit {
		return
	}
	st.player.OffsetXY(x, y)
}

/*
 *	Copy passed state's enemies slice into current state's enemies slice.
 */

func (st *State) CopyEnemiesIntoState(state *State) {
	for i := 0; i < len(state.enemies); i++ {
		*st.enemies[i] = *state.enemies[i]
	}
}

func (st State) PlayerShoot() {
	st.player.Shoot()
}

func (st State) EnemyShoot(index int) {
	if index < len(st.enemies) {
		if st.enemies[index].IsAlive() {
			st.enemies[index].Shoot()
		}
	} else {
		fmt.Printf("Index out of bounds with index %d", index)
	}
}

func (st State) SetEnemyMovementDirectionRight() {
	ENEMY_MOVEMENT_DIRECTION *= -1
}

func (st State) ResetEnemyMovementSpeed() {
	ENEMY_SPEED = 2
}

func (st *State) IncreaseEnemyMovementSpeed(amount int) {
	ENEMY_SPEED += amount
}

func pickARandomNumberInRange(min, max int) int {
	return rand.Intn(max-min) + min
}
