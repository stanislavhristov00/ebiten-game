package main

import (
	"fmt"
	_ "fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/stanislavhristov00/Ebitentestrun/space"
)

const (
	frameOX     = 0
	frameOY     = 0
	frameWidth  = 140
	frameHeight = 120
	frameNum    = 2

	ENEMIES_ON_ROW = 10
	screenWidth    = 640
	screenHeigth   = 480
)

var (
	enemy                    *space.Enemy
	ENEMY_MOVEMENT_DIRECTION = 1
	spriteSheet              *ebiten.Image
	player                   *space.Player
	bulletImage              *ebiten.Image
	enemies                  []*space.Enemy
	enemies2                 []*space.Enemy
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		player.Shoot()
		enemies[0].Shoot()
		enemies[9].Shoot()
		enemies2[2].Shoot()
	}

	bulletX, bulletY := enemies[0].GetBulletXY()
	playerX, playerY := player.GetPlayerXY()
	playerScaleX, _ := player.GetScaleXY()
	bulletScaleX, _ := enemies[0].GetScaleXY()

	if float64(bulletY)*bulletScaleX > float64(playerY) {
		if float64(bulletX)*bulletScaleX > float64(playerX)*playerScaleX-10 &&
			float64(bulletX)*bulletScaleX < float64(playerX)*playerScaleX+90*playerScaleX {
			player.Die()
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.OffsetXY(3, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.OffsetXY(-3, 0)
	}

	x0, _ := enemies[0].GetEnemyXY()
	x9, _ := enemies[ENEMIES_ON_ROW-1].GetEnemyXY()
	if float64(x0)/4 < 0 || float64(x9)*0.25+float64(enemies[ENEMIES_ON_ROW-1].GetFrameWidth())*0.25 > float64(screenWidth) {
		ENEMY_MOVEMENT_DIRECTION *= -1
	}

	for i := 0; i < ENEMIES_ON_ROW; i++ {
		enemies[i].OffsetXY(ENEMY_MOVEMENT_DIRECTION*2, 0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if player.IsAlive() {
		player.Draw(screen)
	} else {
		player.DieDraw(screen)
	}
	for i := 0; i < ENEMIES_ON_ROW; i++ {
		enemies[i].Draw(screen, g.count)
		enemies2[i].Draw(screen, g.count)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	src := getImage("resources/1.png")
	spriteSheet = ebiten.NewImageFromImage(src)

	heroImage := spriteSheet.SubImage(image.Rect(130, 600, 220, 720)).(*ebiten.Image)
	bulletImage = spriteSheet.SubImage(image.Rect(450, 360, 500, 480)).(*ebiten.Image)
	player = space.NewPlayer(heroImage, bulletImage, 0, screenHeigth-90/3, 90, 90, 0.35, 0.35)
	enemy = space.NewEnemy(spriteSheet, bulletImage, 0, 0, 135, 120, 2, 0, 0, 0.25, 0.25)
	enemy2 := space.NewEnemy(spriteSheet, bulletImage, 0, 120, 135, 120, 2, 0, 0, 0.25, 0.25)
	enemies = Load10Enemies(enemy, 1)
	enemies2 = Load10Enemies(enemy2, 2)

	// for _, k := range enemies {
	// 	fmt.Printf("OFFSET X: %d", k.GetX())
	// }

	g := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeigth)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func getImage(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	defer imgFile.Close()
	return img
}

func Load10Enemies(enemy *space.Enemy, row int) []*space.Enemy {
	slice := make([]*space.Enemy, 0)
	enemy.OffsetXY(0, row*enemy.GetFrameHeight())
	for i := 0; i < ENEMIES_ON_ROW; i++ {
		en := enemy.MakeCopy()
		slice = append(slice, en)
		enemy.OffsetXY(enemy.GetFrameWidth(), 0)
	}

	return slice
}
