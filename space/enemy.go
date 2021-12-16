package space

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

/*
 *	Represents a single enemy bullet.
 */
type EnemyBullet struct {
	img        *ebiten.Image
	bulletPosX int
	bulletPosY int
	inAir      bool
}

/*
 *	Represents a single enemy.
 */
type Enemy struct {
	img         []*ebiten.Image
	imgOnDeath  *ebiten.Image
	bullet      *EnemyBullet
	frameWidth  int
	frameHeight int
	numFrames   int
	posX        int
	posY        int
	isAlive     bool
}

/*
 *	Constructor for enemy.
 */

func NewEnemy(img *ebiten.Image, bulletImg *ebiten.Image, frameX, frameY, frameWidth, frameHeigth, numFrames, posX, posY int) *Enemy {
	enemy := &Enemy{}
	enemyBullet := &EnemyBullet{
		img:        bulletImg,
		bulletPosX: posX + frameWidth/2,
		bulletPosY: posY + frameHeigth/2,
		inAir:      false,
	}
	enemy.bullet = enemyBullet
	enemy.img = make([]*ebiten.Image, numFrames)
	enemy.numFrames = numFrames

	for i := 0; i < numFrames; i++ {
		frameX = frameX + i*frameWidth
		enemy.img[i] = img.SubImage(image.Rect(frameX, frameY, frameX+frameWidth, frameY+frameHeigth)).(*ebiten.Image)
	}

	enemy.posX = posX
	enemy.posY = posY
	enemy.isAlive = true
	enemy.frameHeight = frameHeigth
	enemy.frameWidth = frameWidth

	return enemy
}

/*
 *	Load an image for the dying frame.
 */

func (en *Enemy) LoadDeathFrame(img *ebiten.Image) {
	en.imgOnDeath = img
}

/*
 *	SHOOT!
 */

func (en Enemy) Shoot(screen *ebiten.Image) {
	en.bullet.inAir = true
}

/*
 *	Draw the enemy animation on the context screen.
 */

func (en Enemy) Draw(screen *ebiten.Image, count int, scaleX, scaleY float64) {
	op := &ebiten.DrawImageOptions{}

	if en.isAlive {
		i := (count / 20) % en.numFrames

		if !en.bullet.inAir {
			en.bullet.bulletPosX = en.posX + en.frameWidth/2
			en.bullet.bulletPosY = en.posY + en.frameHeight/2
		}

		op.GeoM.Translate(float64(en.posX), float64(en.posY))
		op.GeoM.Scale(scaleX, scaleY)
		screen.DrawImage(en.img[i], op)

		if en.bullet.inAir {
			en.bulletOffsetXY(0, 1)
			op.GeoM.Reset()
			op.GeoM.Scale(0.25, 0.25)
			op.GeoM.Translate(float64(en.bullet.bulletPosX), float64(en.bullet.bulletPosY))
			screen.DrawImage(en.bullet.img, op)
		}
	}
}

/*
 *	Offset X, Y for the enemy bullet.
 */

func (en *Enemy) bulletOffsetXY(x, y int) {
	en.bullet.bulletPosX += x
	en.bullet.bulletPosY += y
}

/*
 *	Offset X, Y of the enemy.
 */

func (en *Enemy) OffsetXY(x, y int) {
	en.posX += x
	en.posY += y
}

/*
 *	Displays the death animation if there is any.
 */

func (en Enemy) Die(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	if en.imgOnDeath != nil {
		op.GeoM.Translate(float64(en.posX), float64(en.posY))
		screen.DrawImage(en.imgOnDeath, op)
	}
}

func (en Enemy) GetX() int {
	return en.posX
}

func (en Enemy) GetY() int {
	return en.posY
}

func (en Enemy) GetFrameWidth() int {
	return en.frameWidth
}

func (en Enemy) GetFrameHeight() int {
	return en.frameHeight
}

func (en Enemy) MakeCopy() *Enemy {
	return &Enemy{
		img:         en.img,
		imgOnDeath:  en.imgOnDeath,
		bullet:      en.bullet,
		numFrames:   en.numFrames,
		frameWidth:  en.frameWidth,
		frameHeight: en.frameHeight,
		posX:        en.posX,
		posY:        en.posY,
		isAlive:     en.isAlive,
	}
}
