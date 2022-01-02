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
	count              int
	state              *State
	screenWidth        int
	stateCopy          *State
	fontSmall          font.Face
	fontBig            font.Face
	fontMedium         font.Face
	isOnStartingScreen bool
	isGameOver         bool
	score              int
	highScore          int
}

func NewGame(numEnemies int, screenWidth int) *Game {
	return &Game{
		count:              0,
		screenWidth:        screenWidth,
		state:              NewState(numEnemies),
		stateCopy:          NewState(numEnemies),
		isOnStartingScreen: true,
		isGameOver:         false,
		score:              0,
		highScore:          0,
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

	g.fontSmall, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	g.fontMedium, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    36,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	g.fontBig, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    72,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

}

func (g Game) drawText(screen *ebiten.Image) {
	if g.isOnStartingScreen {
		if g.fontMedium != nil && g.fontBig != nil {
			text.Draw(screen, "SPACE INVADERS", g.fontBig, 50, 100, color.White)
			text.Draw(screen, "Move  with  arrow keys", g.fontMedium, 130, 150, color.White)
			text.Draw(screen, "Shoot  with  spacebar", g.fontMedium, 135, 200, color.White)
			text.Draw(screen, "PRESS SPACEBAR TO START", g.fontMedium, 115, 350, color.White)
		}
	} else if g.isGameOver {
		if g.fontMedium != nil && g.fontBig != nil {
			text.Draw(screen, "GAME OVER", g.fontBig, 150, 200, color.White)
			text.Draw(screen, "Press  Q  to continue playing", g.fontMedium, 80, 250, color.White)
		}
	} else {
		first := fmt.Sprintf("LIVES  %d", g.state.player.GetLives())
		second := fmt.Sprintf("SCORE  %d", g.score)

		if g.fontSmall != nil {
			text.Draw(screen, first, g.fontSmall, 10, 20, color.White)
			text.Draw(screen, second, g.fontSmall, 120, 20, color.White)
		}
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
		enemies2Copy = append(enemies2Copy, enemies2[i].MakeCopy())
		enemies3Copy = append(enemies3Copy, enemies3[i].MakeCopy())
		enemies4Copy = append(enemies4Copy, enemies4[i].MakeCopy())
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
	if g.isOnStartingScreen {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.isOnStartingScreen = false
		}
	} else if !g.isGameOver {
		g.count++

		g.isGameOver = !g.state.player.IsAlive()
		g.state.CheckIfEnemyShotPlayer()

		if g.state.CheckIfPlayerShotEnemy() {
			g.score += 10
		}

		if g.state.input.Update() {
			g.state.PlayerShoot()
			g.state.EnemyShoot(0)
			g.state.EnemyShoot(9)
			g.state.EnemyShoot(12)
		}

		dir, ok := g.state.input.Dir()

		if ok {
			g.state.MovePlayer(dir.DirToValue()*6, 0)
		}

		g.state.MoveEnemies(g.screenWidth)

		if g.state.CheckIfAllEnemiesAreDead() {
			g.LoadNextWave(1)
		}

	} else {
		g.Restart()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawText(screen)

	if !g.isGameOver && !g.isOnStartingScreen {
		g.state.DrawPlayer(screen)
		g.state.DrawEnemies(screen, g.count)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Restart() {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.state.CopyEnemiesIntoState(g.stateCopy)
		g.state.player.Revive()
		g.state.SetEnemyMovementDirectionRight()
		g.isGameOver = false
	}
}

func (g *Game) LoadNextWave(enemySpeed int) {
	g.state.CopyEnemiesIntoState(g.stateCopy)
	g.state.SetEnemyMovementDirectionRight()
	g.state.IncreaseEnemyMovementSpeed(enemySpeed)
}
