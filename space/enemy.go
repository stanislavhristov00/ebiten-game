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
	img             []*ebiten.Image
	imgOnDeath      *ebiten.Image
	bullet          *EnemyBullet
	frameWidth      int
	frameHeight     int
	numFrames       int
	posX            int
	posY            int
	scaleX          float64
	scaleY          float64
	deathFrameDrawn bool
	isAlive         bool
}

/*
 *	Constructor for enemy.
 */

func NewEnemy(img *ebiten.Image, bulletImg *ebiten.Image,
	frameX, frameY, frameWidth, frameHeigth,
	numFrames, posX, posY int, scaleX, scaleY float64) *Enemy {
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
	enemy.scaleX = scaleX
	enemy.scaleY = scaleY
	enemy.deathFrameDrawn = false

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

func (en Enemy) Draw(screen *ebiten.Image, count int) {
	op := &ebiten.DrawImageOptions{}

	if en.bullet.inAir {
		_, h := screen.Size()

		op.GeoM.Reset()
		en.bulletOffsetXY(0, 3)
		op.GeoM.Scale(en.scaleX, en.scaleY)
		/*
		 * Taking into consideration the scale of the bullet, when translating on offsets
		 */
		op.GeoM.Translate(float64(en.bullet.bulletPosX)*en.scaleX, float64(en.bullet.bulletPosY)*en.scaleY)
		screen.DrawImage(en.bullet.img, op)

		if float64(en.bullet.bulletPosY)*en.scaleY > float64(h) {
			en.bullet.inAir = false
		}
	}

	if en.isAlive {
		i := (count / 20) % en.numFrames

		op.GeoM.Reset()

		op.GeoM.Scale(en.scaleX, en.scaleY)
		/*
		 * We have to scale the offsets of the enemy according to its scale (size).
		 */
		op.GeoM.Translate(float64(en.posX)*en.scaleX, float64(en.posY)*en.scaleY)

		if !en.bullet.inAir {
			en.bullet.bulletPosX = en.posX + int((float64(en.frameWidth)/2)*en.scaleX)
			en.bullet.bulletPosY = en.posY + int((float64(en.frameHeight) * en.scaleY))
		}

		screen.DrawImage(en.img[i], op)
	}
}

/*
 *	Offsets x, y for the enemy bullet.
 */

func (en *Enemy) bulletOffsetXY(x, y int) {
	en.bullet.bulletPosX += x
	en.bullet.bulletPosY += y
}

/*
 *	Offsets x, y of the enemy.
 */

func (en *Enemy) OffsetXY(x, y int) {
	en.posX += x
	en.posY += y
}

/*
 *	Displays the death animation if there is any.
 */

func (en Enemy) DieDraw(screen *ebiten.Image) {
	if en.imgOnDeath != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(en.scaleX, en.scaleY)
		/*
		 * We have to scale the offsets of the enemy according to its scale (size).
		 */
		op.GeoM.Translate(float64(en.posX)*en.scaleX, float64(en.posY)*en.scaleY)
		screen.DrawImage(en.imgOnDeath, op)
	}
}

/*
 * Set Enemy state as dead.
 */
func (en *Enemy) Die() {
	en.isAlive = false
}

/*
 *	Makes a deep copy of an Enemy object
 */

func (en Enemy) MakeCopy() *Enemy {
	return &Enemy{
		img:         en.img,
		imgOnDeath:  en.imgOnDeath,
		bullet:      MakeBullet(en.bullet.img, en.bullet.bulletPosX, en.bullet.bulletPosY, en.bullet.inAir),
		numFrames:   en.numFrames,
		frameWidth:  en.frameWidth,
		frameHeight: en.frameHeight,
		scaleX:      en.scaleX,
		scaleY:      en.scaleY,
		posX:        en.posX,
		posY:        en.posY,
		isAlive:     en.isAlive,
	}
}

/*
 *	Makes a deep copy of an EnemyBullet object
 */

func MakeBullet(bulletImg *ebiten.Image, posX, posY int, inAir bool) *EnemyBullet {
	return &EnemyBullet{
		img:        bulletImg,
		bulletPosX: posX,
		bulletPosY: posY,
		inAir:      false,
	}
}

/*
 *	Check if bullet collides with player. Return true if so, false otherwise.
 */

func (en Enemy) BulletCollisionWithPlayer(player *Player) bool {
	if en.bullet.inAir {
		bulletX, bulletY := en.bullet.bulletPosX, en.bullet.bulletPosY
		bulletScaleY, bulletScaleX := en.scaleX, en.scaleY
		playerX, playerY := player.GetPlayerXY()
		playerScaleX, _ := player.GetScaleXY()

		if float64(bulletY)*bulletScaleY > float64(playerY)-10 {
			if float64(bulletX)*bulletScaleX > float64(playerX)*playerScaleX-10 &&
				float64(bulletX)*bulletScaleX < float64(playerX)*playerScaleX+90*playerScaleX {
				return true
			}
		}
	}

	return false
}

func (en Enemy) GetBulletXY() (int, int) {
	return en.bullet.bulletPosX, en.bullet.bulletPosY
}

func (en Enemy) GetScaleXY() (float64, float64) {
	return en.scaleX, en.scaleY
}

func (en *Enemy) SetBulletInAir(inAir bool) {
	en.bullet.inAir = inAir
}

func (en Enemy) GetEnemyXY() (int, int) {
	return en.posX, en.posY
}

func (en Enemy) GetFrameWidth() int {
	return en.frameWidth
}

func (en Enemy) GetFrameHeight() int {
	return en.frameHeight
}

func (en Enemy) IsAlive() bool {
	return en.isAlive
}

func (en *Enemy) SetDeathFrameDrawn(isSet bool) {
	en.deathFrameDrawn = isSet
}

func (en Enemy) GetDeathFrameDrawn() bool {
	return en.deathFrameDrawn
}

func (en Enemy) IsBulletInAir() bool {
	return en.bullet.inAir
}
