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
	screenWidth  = 640
	screenHeight = 480

	frameOX     = 0
	frameOY     = 0
	frameWidth  = 140
	frameHeight = 120
	frameNum    = 2
)

var (
	runnerImage *ebiten.Image
	enemy       *space.Enemy
)

type Sprite struct {
	frames  []*ebiten.Image
	frameOX uint32
	frameOY uint32
}

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
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		fmt.Println("S")
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// w, y := screen.Size()
	// fmt.Println(w, y)
	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Scale(0.5, 0.5)
	//op.GeoM.Translate(float64(g.count), 0)
	// i := (g.count / 10) % frameNum
	// sx, sy := frameOX+i*frameWidth, frameOY
	// screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

	// if g.count%10 == 0 {
	// 	enemy.Die(screen, op)
	// } else {
	// 	enemy.Draw(screen, op, g.count)
	// 	enemy.OffsetXY(1, 0)
	// }

	//enemy.Die(screen, op)
	//enemy.OffsetXY(1, 0)

	// op2 := &ebiten.DrawImageOptions{}
	// op2.GeoM.Scale(0.5, 0.5)
	// op2.GeoM.Translate(0, 50)
	// i2 := (g.count / 10) % frameNum
	// sx2, sy2 := frameOX+i2*frameWidth, frameOY+frameHeight
	// screen.DrawImage(runnerImage.SubImage(image.Rect(sx2, sy2, sx2+frameWidth, sy2+frameHeight)).(*ebiten.Image), op2)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	src := getImage("resources/1.png")
	runnerImage = ebiten.NewImageFromImage(src)

	fmt.Println(runnerImage)

	//enemy = space.NewEnemy(runnerImage, frameOX, frameOY, frameWidth, frameHeight, frameNum, 0, 0)

	g := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func getImage(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	return img
}
