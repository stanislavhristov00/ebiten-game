package space

// import (
// 	"github.com/hajimehoshi/ebiten/v2"
// )

// type Game struct {
// 	count 			int
// 	isFullScreen 	bool
// 	state			State
// }

// func NewGame() *Game {
// 	return &Game{
// 		count : 0,
// 		isFullScreen: false,
// 		state: *NewState(),
// 	}
// }

// func (g *Game) Init(enemies [][]Enemy, player *Player) {
// 	for index, _ := range enemies {
// 		g.state.LoadEnemies(enemies[index])
// 	}
// 	g.state.LoadPlayer(player)
// }

// func (g *Game) Update() error {
// 	g.count++
// 	//fmt.Println(g.isFullScreen)
// 	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
// 		fmt.Println("Q is pressed")
// 		g.isFullScreen = !g.isFullScreen
// 		ebiten.SetFullscreen(g.isFullScreen)
// 	}

// 	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
// 		player.Shoot()
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
// 		player.OffsetXY(1, 0)
// 	}

// 	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
// 		player.OffsetXY(-1, 0)
// 	}
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	//op := &ebiten.DrawImageOptions{}
// 	//op.GeoM.Scale(0.5, 0.5)
// 	//op.GeoM.Translate(float64(g.count), 0)
// 	// i := (g.count / 10) % frameNum
// 	// sx, sy := frameOX+i*frameWidth, frameOY
// 	// screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
// 	//op := &ebiten.DrawImageOptions{}
// 	player.Draw(screen)
// 	//op.GeoM.Translate(0, 0)
// 	//screen.DrawImage(bulletImage, op)
// 	//op.GeoM.Reset()

// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
// 	return outsideWidth, outsideHeight
// }
