package main

import (
	"fmt"
	_ "fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"syscall"

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
	SCREEN_SIZE_X  = 640
	SCREEN_SIZE_Y  = 480
)

var (
	enemy                    *space.Enemy
	ENEMY_MOVEMENT_DIRECTION = 1
	spriteSheet              *ebiten.Image
	player                   *space.Player
	bulletImage              *ebiten.Image
	enemies                  []*space.Enemy
	user32                   = syscall.NewLazyDLL("User32.dll")
	getSystemMetrics         = user32.NewProc("GetSystemMetrics")
	screenWidth              = SCREEN_SIZE_X
	screenHeight             = SCREEN_SIZE_Y
	SM_CXSCREEN              = 0
	SM_CYSCREEN              = 1
)

type Game struct {
	count        int
	isFullScreen bool
}

func (g *Game) Update() error {
	g.count++
	//fmt.Println(g.isFullScreen)
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		fmt.Println("Q is pressed")
		g.isFullScreen = !g.isFullScreen
		ebiten.SetFullscreen(g.isFullScreen)
		if g.isFullScreen {
			screenWidth = GetSystemMetrics(SM_CXSCREEN)
			screenWidth = GetSystemMetrics(SM_CYSCREEN)
		} else {
			screenWidth = SCREEN_SIZE_X
			screenHeight = SCREEN_SIZE_Y
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		player.Shoot()
		enemies[0].Shoot()
		enemies[9].Shoot()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.OffsetXY(3, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.OffsetXY(-3, 0)
	}

	if float64(enemies[0].GetX())/4 < 0 || float64(enemies[ENEMIES_ON_ROW-1].GetX())/4+float64(enemies[ENEMIES_ON_ROW-1].GetFrameWidth())/4 > float64(screenWidth) {
		ENEMY_MOVEMENT_DIRECTION *= -1
	}

	// if enemies[0].GetX() < 0 || enemies[ENEMIES_ON_ROW-1].GetX()+enemies[ENEMIES_ON_ROW-1].GetFrameWidth()/4 > screenWidth {
	// 	ENEMY_MOVEMENT_DIRECTION *= -1
	// }

	for i := 0; i < ENEMIES_ON_ROW; i++ {
		enemies[i].OffsetXY(ENEMY_MOVEMENT_DIRECTION*2, 0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Scale(0.5, 0.5)
	//op.GeoM.Translate(float64(g.count), 0)
	// i := (g.count / 10) % frameNum
	// sx, sy := frameOX+i*frameWidth, frameOY
	// screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
	//op := &ebiten.DrawImageOptions{}
	player.Draw(screen, 0.35, 0.35)
	//enemy.Draw(screen, g.count, 0.25, 0.25)
	for i := 0; i < ENEMIES_ON_ROW; i++ {
		enemies[i].Draw(screen, g.count, 0.25, 0.25)
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
	player = space.NewPlayer(heroImage, bulletImage, 0, screenHeight-30)
	enemy = space.NewEnemy(spriteSheet, bulletImage, 0, 0, 140, 120, 2, 0, 0)
	enemies = Load10Enemies(enemy, 1)

	// for _, k := range enemies {
	// 	fmt.Printf("OFFSET X: %d", k.GetX())
	// }

	g := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
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
	//enemy.OffsetXY(0, row*int(float64(enemy.GetFrameHeight())*0.25))
	enemy.OffsetXY(0, row*enemy.GetFrameHeight())
	for i := 0; i < ENEMIES_ON_ROW; i++ {
		en := enemy.MakeCopy()
		slice = append(slice, en)
		//enemy.OffsetXY(int(float64(enemy.GetFrameWidth())*0.25), 0)
		enemy.OffsetXY(enemy.GetFrameWidth(), 0)
	}

	return slice
}

func GetSystemMetrics(nIndex int) int {
	index := uintptr(nIndex)
	ret, _, _ := getSystemMetrics.Call(index)
	return int(ret)
}
