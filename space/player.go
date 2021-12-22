package space

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	NUM_LIVES = 3
)

/*
 *	Represents a player's bullet.
 */
type PlayerBullet struct {
	img        *ebiten.Image
	bulletPosX int
	bulletPosY int
	inAir      bool
}

/*
 *	Represents a player.
 */
type Player struct {
	img         *ebiten.Image
	imgOnDeath  *ebiten.Image
	bullet      *PlayerBullet
	posX        int
	posY        int
	frameWidth  int
	frameHeigth int
	scaleX      float64
	scaleY      float64
	lives       int
	isAlive     bool
}

/*
 *	Constructor for player.
 */

func NewPlayer(img *ebiten.Image, bulletImg *ebiten.Image,
	posX, posY, frameWidth, frameHeigth int, scaleX, scaleY float64) *Player {
	playerBullet := &PlayerBullet{
		img:        bulletImg,
		bulletPosX: posX,
		bulletPosY: posY,
		inAir:      false,
	}

	player := &Player{
		bullet:      playerBullet,
		img:         img,
		posX:        posX,
		posY:        posY,
		frameWidth:  frameWidth,
		frameHeigth: frameHeigth,
		scaleX:      scaleX,
		scaleY:      scaleY,
		lives:       NUM_LIVES,
		isAlive:     true,
	}

	return player
}

/*
 *	Load a death frame.
 */

func (p *Player) LoadDeathFrame(img *ebiten.Image) {
	p.imgOnDeath = img
}

/*
 *	Offsets a bullet's coordinates.
 */

func (p *Player) offsetBulletXY(x, y int) {
	p.bullet.bulletPosX += x
	p.bullet.bulletPosY += y
}

/*
 *	Represents a player's bullet.
 */

func (p *Player) OffsetXY(x, y int) {
	p.posX += x
	p.posY += y
}

/*
 *	Draw the enemy animation on the context screen.
 */

func (p Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if p.isAlive {
		op.GeoM.Scale(p.scaleX, p.scaleY)
		op.GeoM.Translate(float64(p.posX)*p.scaleX, float64(p.posY))
		if !p.bullet.inAir {
			p.bullet.bulletPosX = p.posX + 30
			p.bullet.bulletPosY = p.posY
		}

		screen.DrawImage(p.img, op)

		op.GeoM.Reset()
		if p.bullet.inAir {
			p.offsetBulletXY(0, -3)
			op.GeoM.Scale(0.25, 0.25)
			op.GeoM.Translate(float64(p.bullet.bulletPosX)*p.scaleX, float64(p.bullet.bulletPosY))
			screen.DrawImage(p.bullet.img, op)

			if p.bullet.bulletPosY < 0 {
				p.bullet.inAir = false
			}
		}
	}
}

/*
 *	Shoot a bullet from the player.
 */

func (p *Player) Shoot() {
	p.bullet.inAir = true
}

/*
 *	Displays the death animation if there is any.
 */

func (p *Player) DieDraw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(p.scaleX, p.scaleY)
	op.GeoM.Translate(float64(p.posX)*p.scaleX, float64(p.posY))
	if p.imgOnDeath != nil {
		op.GeoM.Translate(float64(p.posX)*p.scaleX, float64(p.posY))
		screen.DrawImage(p.imgOnDeath, op)
	}
}

/*
 * Change the state of a plyaer to dead.
 */

func (p *Player) Die() {
	p.isAlive = false
}

func (p Player) BulletCollisionWithEnemy(en *Enemy) bool {
	if p.bullet.inAir {
		bulletX, bulletY := p.bullet.bulletPosX, p.bullet.bulletPosY
		_, bulletScaleX := p.scaleX, p.scaleY
		enemyX, enemyY := en.GetEnemyXY()
		enemyScaleX, enemyScaleY := en.GetScaleXY()
		enemyWidth := en.GetFrameWidth()

		if float64(bulletY) < float64(enemyY)*enemyScaleY {
			if float64(bulletX)*bulletScaleX > float64(enemyX)*enemyScaleX-10 &&
				float64(bulletX)*bulletScaleX < float64(enemyX)*enemyScaleX+float64(enemyWidth)*enemyScaleX {
				return true
			}
		}
	}

	return false
}

func (p *Player) LoseLife() {
	p.lives -= 1

	if p.lives == 0 {
		p.Die()
	}
}

func (p *Player) SetBulletInAir(inAir bool) {
	p.bullet.inAir = inAir
}

func (p Player) IsAlive() bool {
	return p.isAlive
}

func (p Player) GetPlayerXY() (int, int) {
	return p.posX, p.posY
}

func (p Player) GetScaleXY() (float64, float64) {
	return p.scaleX, p.scaleY
}

func (p Player) GetFrameWidth() int {
	return p.frameWidth
}
