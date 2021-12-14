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
	spriteSheet *ebiten.Image
	player      *space.Player
	bulletImage *ebiten.Image
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
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		player.Shoot()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.OffsetXY(1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.OffsetXY(-1, 0)
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
	player.Draw(screen)
	//op.GeoM.Translate(0, 0)
	//screen.DrawImage(bulletImage, op)
	//op.GeoM.Reset()

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	src := getImage("resources/1.png")
	spriteSheet = ebiten.NewImageFromImage(src)

	heroImage := spriteSheet.SubImage(image.Rect(130, 600, 220, 720)).(*ebiten.Image)
	bulletImage = spriteSheet.SubImage(image.Rect(450, 360, 500, 480)).(*ebiten.Image)
	player = space.NewPlayer(heroImage, bulletImage, 0, 480-90)

	g := &Game{}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Animation (Ebiten Demo)")
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
