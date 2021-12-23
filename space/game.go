package space

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	count       int
	state       *State
	screenWidth int
}

func NewGame(numEnemies int, screenWidth int) *Game {
	return &Game{
		count:       0,
		screenWidth: screenWidth,
		state:       NewState(numEnemies),
	}
}

func (g *Game) LoadEnemies(enemies []*Enemy, enemies2 []*Enemy, enemies3 []*Enemy, enemies4 []*Enemy, enemies5 []*Enemy) {
	g.state.LoadEnemies(enemies)
	g.state.LoadEnemies(enemies2)
	g.state.LoadEnemies(enemies3)
	g.state.LoadEnemies(enemies4)
	g.state.LoadEnemies(enemies5)
}

func (g *Game) LoadPlayer(player *Player) {
	g.state.LoadPlayer(player)
}

func (g *Game) Update() error {
	g.count++

	if g.state.input.Update() {
		g.state.PlayerShoot()
		g.state.EnemyShoot(0)
		g.state.EnemyShoot(9)
		g.state.EnemyShoot(12)
	}

	g.state.CheckIfEnemyShotPlayer()

	g.state.CheckIfPlayerShotEnemy()

	dir, ok := g.state.input.Dir()

	if ok {
		g.state.MovePlayer(dir.DirToValue()*3, 0)
	}

	g.state.MoveEnemies(g.screenWidth)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.DrawPlayer(screen)
	g.state.DrawEnemies(screen, g.count)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
