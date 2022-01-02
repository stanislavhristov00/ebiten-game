package space

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	dpi = 72
)

type Game struct {
	count       int
	state       *State
	screenWidth int
	stateCopy   *State
	font        font.Face
}

func NewGame(numEnemies int, screenWidth int) *Game {
	return &Game{
		count:       0,
		screenWidth: screenWidth,
		state:       NewState(numEnemies),
		stateCopy:   NewState(numEnemies),
	}
}

func (g *Game) InitFont(fileName string) {
	bytes, e := ioutil.ReadFile(fileName)

	if e != nil {
		log.Fatal(e)
	}

	tt, err := opentype.Parse(bytes)

	if err != nil {
		log.Fatal(err)
	}

	g.font, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

}

func (g Game) drawText(screen *ebiten.Image) {
	first := fmt.Sprintf("LIVES  %d", g.state.player.GetLives())

	if g.font != nil {
		text.Draw(screen, first, g.font, 10, 20, color.White)
	}
}

func (g *Game) LoadEnemies(enemies []*Enemy, enemies2 []*Enemy, enemies3 []*Enemy, enemies4 []*Enemy, enemies5 []*Enemy) {
	g.state.LoadEnemies(enemies)
	g.state.LoadEnemies(enemies2)
	g.state.LoadEnemies(enemies3)
	g.state.LoadEnemies(enemies4)
	g.state.LoadEnemies(enemies5)

	enemiesCopy := make([]*Enemy, 0)
	enemies2Copy := make([]*Enemy, 0)
	enemies3Copy := make([]*Enemy, 0)
	enemies4Copy := make([]*Enemy, 0)
	enemies5Copy := make([]*Enemy, 0)

	for i := 0; i < g.state.numEnemies/5; i++ {
		enemiesCopy = append(enemiesCopy, enemies[i].MakeCopy())
	}

	for i := 0; i < g.state.numEnemies/5; i++ {
		enemies2Copy = append(enemies2Copy, enemies2[i].MakeCopy())
	}

	for i := 0; i < g.state.numEnemies/5; i++ {
		enemies3Copy = append(enemies3Copy, enemies3[i].MakeCopy())
	}

	for i := 0; i < g.state.numEnemies/5; i++ {
		enemies4Copy = append(enemies4Copy, enemies4[i].MakeCopy())
	}

	for i := 0; i < g.state.numEnemies/5; i++ {
		enemies5Copy = append(enemies5Copy, enemies5[i].MakeCopy())
	}

	g.stateCopy.LoadEnemies(enemiesCopy)
	g.stateCopy.LoadEnemies(enemies2Copy)
	g.stateCopy.LoadEnemies(enemies3Copy)
	g.stateCopy.LoadEnemies(enemies4Copy)
	g.stateCopy.LoadEnemies(enemies5Copy)

}

func (g *Game) LoadPlayer(player *Player) {
	g.state.LoadPlayer(player)
}

func (g *Game) Update() error {
	if g.state.player.IsAlive() {
		g.count++

		g.state.CheckIfEnemyShotPlayer()

		g.state.CheckIfPlayerShotEnemy()

		if g.state.input.Update() {
			g.state.PlayerShoot()
			g.state.EnemyShoot(0)
			g.state.EnemyShoot(9)
			g.state.EnemyShoot(12)
		}

		dir, ok := g.state.input.Dir()

		if ok {
			g.state.MovePlayer(dir.DirToValue()*3, 0)
		}

		g.state.MoveEnemies(g.screenWidth)
	} else {
		g.Restart()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.state.DrawPlayer(screen)
	g.state.DrawEnemies(screen, g.count)
	g.drawText(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Restart() {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.state.CopyEnemiesIntoState(g.stateCopy)
		g.state.player.Revive()
	}
}
