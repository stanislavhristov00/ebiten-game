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
	enemy.bullet = MakeBullet(bulletImg, posX, posY, false)
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

func (en Enemy) Shoot() {
	en.bullet.inAir = true
}

/*
 *	Draw the enemy animation on the context screen.
 */

func (en Enemy) Draw(screen *ebiten.Image, count int, scaleX, scaleY float64) {
	op := &ebiten.DrawImageOptions{}

	if en.isAlive {
		i := (count / 20) % en.numFrames

		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(float64(en.posX)*scaleX, float64(en.posY)*scaleY)

		if !en.bullet.inAir {
			en.bullet.bulletPosX = en.posX + int((float64(en.frameWidth)/2)*scaleX)
			en.bullet.bulletPosY = en.posY + int((float64(en.frameHeight) * scaleY))
		}

		_, h := screen.Size()

		screen.DrawImage(en.img[i], op)

		if en.bullet.inAir {
			op.GeoM.Reset()
			en.bulletOffsetXY(0, 3)
			op.GeoM.Scale(scaleX, scaleY)
			op.GeoM.Translate(float64(en.bullet.bulletPosX)*scaleX, float64(en.bullet.bulletPosY)*scaleY)
			screen.DrawImage(en.bullet.img, op)

			if float64(en.bullet.bulletPosY)*scaleY > float64(h) {
				en.bullet.inAir = false
			}
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
		bullet:      MakeBullet(en.bullet.img, en.bullet.bulletPosX, en.bullet.bulletPosY, en.bullet.inAir),
		numFrames:   en.numFrames,
		frameWidth:  en.frameWidth,
		frameHeight: en.frameHeight,
		posX:        en.posX,
		posY:        en.posY,
		isAlive:     en.isAlive,
	}
}

func MakeBullet(bulletImg *ebiten.Image, posX, posY int, inAir bool) *EnemyBullet {
	return &EnemyBullet{
		img:        bulletImg,
		bulletPosX: posX,
		bulletPosY: posY,
		inAir:      false,
	}
}
