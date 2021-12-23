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

// type Game struct {
// 	count int
// 	input *space.Input
// }

// func (g *Game) Update() error {
// 	g.count++

// 	// if inpututil.IsKeyJustPressed(ebiten.KeyS) {
// 	// 	player.Shoot()
// 	// 	enemies[0].Shoot()
// 	// 	enemies[9].Shoot()
// 	// 	enemies2[2].Shoot()
// 	// }

// 	if g.input.Update() {
// 		player.Shoot()
// 		enemies[0].Shoot()
// 		enemies[9].Shoot()
// 		enemies2[2].Shoot()
// 	}

// 	// bulletX, bulletY := enemies[0].GetBulletXY()
// 	// playerX, playerY := player.GetPlayerXY()
// 	// playerScaleX, _ := player.GetScaleXY()
// 	// bulletScaleX, bulletScaleY := enemies[0].GetScaleXY()

// 	// if float64(bulletY)*bulletScaleY > float64(playerY)-10 {
// 	// 	if float64(bulletX)*bulletScaleX > float64(playerX)*playerScaleX-10 &&
// 	// 		float64(bulletX)*bulletScaleX < float64(playerX)*playerScaleX+90*playerScaleX {
// 	// 		enemies[0].SetBulletInAir(false)
// 	// 		player.Die()
// 	// 	}
// 	// }

// 	// if enemies[0].BulletCollisionWithPlayer(player) {
// 	// 	player.Die()
// 	// 	enemies[0].SetBulletInAir(false)
// 	// }

// 	// if player.BulletCollisionWithEnemy(enemies[0]) {
// 	// 	enemies[0].Die()
// 	// 	player.SetBulletInAir(false)
// 	// }

// 	// if enemies2[0].IsAlive() {
// 	// 	if player.BulletCollisionWithEnemy(enemies2[0]) {
// 	// 		enemies2[0].Die()
// 	// 		player.SetBulletInAir(false)
// 	// 	}
// 	// }

// 	for i := 0; i < ENEMIES_ON_ROW; i++ {
// 		if enemies[i].IsAlive() {
// 			if player.BulletCollisionWithEnemy(enemies[i]) {
// 				enemies[i].Die()
// 				player.SetBulletInAir(false)
// 			}
// 		}

// 		if enemies2[i].IsAlive() {
// 			if player.BulletCollisionWithEnemy(enemies2[i]) {
// 				enemies2[i].Die()
// 				player.SetBulletInAir(false)
// 			}
// 		}
// 	}

// 	// if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
// 	// 	player.OffsetXY(3, 0)
// 	// }

// 	// if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
// 	// 	player.OffsetXY(-3, 0)
// 	// }

// 	dir, ok := g.input.Dir()

// 	if ok {
// 		player.OffsetXY(dir.DirToValue()*3, 0)
// 	}

// 	// x0, _ := enemies[0].GetEnemyXY()
// 	// x9, _ := enemies[ENEMIES_ON_ROW-1].GetEnemyXY()
// 	// if float64(x0)/4 < 0 || float64(x9)*0.25+float64(enemies[ENEMIES_ON_ROW-1].GetFrameWidth())*0.25 > float64(screenWidth) {
// 	// 	ENEMY_MOVEMENT_DIRECTION *= -1
// 	// }

// 	for i := 0; i < ENEMIES_ON_ROW; i++ {
// 		if enemies[i].IsAlive() {
// 			xi, _ := enemies[i].GetEnemyXY()
// 			scaleX, _ := enemies[i].GetScaleXY()
// 			xFrameWidth := enemies[i].GetFrameWidth()
// 			if float64(xi)*scaleX < 0 || float64(xi)*scaleX+float64(xFrameWidth)*scaleX > float64(screenWidth) {
// 				ENEMY_MOVEMENT_DIRECTION_1 *= -1
// 			}
// 		}

// 		if enemies2[i].IsAlive() {
// 			xi, _ := enemies2[i].GetEnemyXY()
// 			scaleX, _ := enemies2[i].GetScaleXY()
// 			xFrameWidth := enemies2[i].GetFrameWidth()
// 			if float64(xi)*scaleX < 0 || float64(xi)*scaleX+float64(xFrameWidth)*scaleX > float64(screenWidth) {
// 				ENEMY_MOVEMENT_DIRECTION_2 *= -1
// 			}
// 		}
// 	}

// 	// for i := 0; i < ENEMIES_ON_ROW; i++ {
// 	// 	enemies[i].OffsetXY(ENEMY_MOVEMENT_DIRECTION*2, 0)
// 	// 	enemies2[i].OffsetXY(ENEMY_MOVEMENT_DIRECTION*2, 0)
// 	// }

// 	for i := 0; i < ENEMIES_ON_ROW; i++ {
// 		enemies[i].OffsetXY(ENEMY_MOVEMENT_DIRECTION_1*2, 0)
// 		enemies2[i].OffsetXY(ENEMY_MOVEMENT_DIRECTION_2*2, 0)
// 	}
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	if player.IsAlive() {
// 		player.Draw(screen)
// 	} else {
// 		player.DieDraw(screen)
// 	}
// 	for i := 0; i < ENEMIES_ON_ROW; i++ {
// 		enemies[i].Draw(screen, g.count)
// 		enemies2[i].Draw(screen, g.count)
// 	}

// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return outsideWidth, outsideHeight
// }

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

	// for _, k := range enemies {
	// 	fmt.Printf("OFFSET X: %d", k.GetX())
	// }

	// g := &Game{
	// 	input: space.NewInput(),
	// }

	g := space.NewGame(55, screenWidth)
	// en := make([][]*space.Enemy, 5)

	// en[0] = enemies
	// en[1] = enemies2
	// en[2] = enemies3
	// en[3] = enemies4
	// en[4] = enemies5

	// // for i := 0; i < 11; i++ {
	// // 	en = append(en, enemies[i])
	// // }

	// // for i := 0; i < 11; i++ {
	// // 	en = append(en, enemies2[i])
	// // }

	// // for i := 0; i < 11; i++ {
	// // 	en = append(en, enemies3[i])
	// // }

	// // for i := 0; i < 11; i++ {
	// // 	en = append(en, enemies4[i])
	// // }

	// // for i := 0; i < 11; i++ {
	// // 	en = append(en, enemies5[i])
	// // }

	g.LoadEnemies(enemies, enemies2, enemies3, enemies4, enemies5)
	g.LoadPlayer(player)

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
	enemy.OffsetXY(0, row*enemy.GetFrameHeight()+100)
	for i := 0; i < ENEMIES_ON_ROW; i++ {
		en := enemy.MakeCopy()
		slice = append(slice, en)
		enemy.OffsetXY(enemy.GetFrameWidth(), 0)
	}

	return slice
}
