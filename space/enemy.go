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
	img        []*ebiten.Image
	imgOnDeath *ebiten.Image
	shotImg    *ebiten.Image
	bullet     *EnemyBullet
	numFrames  int
	posX       int
	posY       int
	isAlive    bool
}

/*
 *	Constructor for enemy.
 */

func NewEnemy(img *ebiten.Image, bulletImg *ebiten.Image, frameX, frameY, frameWidth, frameHeigth, numFrames, posX, posY int) *Enemy {
	enemy := &Enemy{}
	enemyBullet := &EnemyBullet{
		img:        bulletImg,
		bulletPosX: 0,
		bulletPosY: 0,
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

func (en Enemy) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, count int) {
	if en.isAlive {
		i := (count / 20) % en.numFrames

		if !en.bullet.inAir {
			en.bullet.bulletPosX = en.posX
			en.bullet.bulletPosY = en.posY
		}

		op.GeoM.Translate(float64(en.posX), float64(en.posY))
		screen.DrawImage(en.img[i], op)

		if en.bullet.inAir {
			en.bulletOffsetXY(0, 1)
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
