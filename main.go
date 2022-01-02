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
	"github.com/stanislavhristov00/Ebitentestrun/space"
)

const (
	frameOX     = 0
	frameOY     = 0
	frameWidth  = 140
	frameHeight = 120
	frameNum    = 2

	// There will be five rows in total but you can control the number of enemies.
	ENEMIES_ON_ROW = 11
	screenWidth    = 640
	screenHeigth   = 480
)

var (
	enemy                      *space.Enemy
	ENEMY_MOVEMENT_DIRECTION_1 = 1
	ENEMY_MOVEMENT_DIRECTION_2 = 1
	spriteSheet                *ebiten.Image
	player                     *space.Player
	bulletImage                *ebiten.Image
	enemies                    []*space.Enemy
	enemies2                   []*space.Enemy
)

func main() {
	src := getImage("resources/1.png")
	spriteSheet = ebiten.NewImageFromImage(src)

	heroImage := spriteSheet.SubImage(image.Rect(130, 600, 220, 720)).(*ebiten.Image)
	bulletImage = spriteSheet.SubImage(image.Rect(450, 360, 500, 480)).(*ebiten.Image)
	deathImage := spriteSheet.SubImage(image.Rect(340, 600, 430, 720)).(*ebiten.Image)
	player = space.NewPlayer(heroImage, bulletImage, 0, screenHeigth-90/3, 90, 90, 0.35, 0.35)
	enemy = space.NewEnemy(spriteSheet, bulletImage, 0, 0, 135, 120, 2, 0, 0, 0.25, 0.25)
	enemy.LoadDeathFrame(deathImage)
	enemy2 := space.NewEnemy(spriteSheet, bulletImage, 0, 120, 135, 120, 2, 0, 0, 0.25, 0.25)
	enemy2.LoadDeathFrame(deathImage)
	enemies = LoadRowEnemies(enemy, 1)
	enemies2 = LoadRowEnemies(enemy2, 2)
	enemy3 := space.NewEnemy(spriteSheet, bulletImage, 0, 120, 135, 120, 2, 0, 0, 0.25, 0.25)
	enemy3.LoadDeathFrame(deathImage)
	enemies3 := LoadRowEnemies(enemy3, 3)
	enemy4 := space.NewEnemy(spriteSheet, bulletImage, 0, 120, 135, 120, 2, 0, 0, 0.25, 0.25)
	enemy4.LoadDeathFrame(deathImage)
	enemies4 := LoadRowEnemies(enemy4, 4)
	enemy5 := space.NewEnemy(spriteSheet, bulletImage, 0, 120, 135, 120, 2, 0, 0, 0.25, 0.25)
	enemy5.LoadDeathFrame(deathImage)
	enemies5 := LoadRowEnemies(enemy5, 5)

	g := space.NewGame(55, screenWidth)

	g.LoadEnemies(enemies, enemies2, enemies3, enemies4, enemies5)
	g.LoadPlayer(player)
	g.InitFont("resources/font/font.ttf")

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

func LoadRowEnemies(enemy *space.Enemy, row int) []*space.Enemy {
	slice := make([]*space.Enemy, 0)
	enemy.OffsetXY(0, row*enemy.GetFrameHeight()+50)
	for i := 0; i < ENEMIES_ON_ROW; i++ {
		en := enemy.MakeCopy()
		slice = append(slice, en)
		enemy.OffsetXY(enemy.GetFrameWidth(), 0)
	}

	return slice
}
