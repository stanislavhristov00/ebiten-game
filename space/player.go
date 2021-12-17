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
	img        *ebiten.Image
	imgOnDeath *ebiten.Image
	bullet     *PlayerBullet
	posX       int
	posY       int
	lives      int
	isAlive    bool
}

/*
 *	Constructor for player.
 */

func NewPlayer(img *ebiten.Image, bulletImg *ebiten.Image, posX, posY int) *Player {
	playerBullet := &PlayerBullet{
		img:        bulletImg,
		bulletPosX: posX,
		bulletPosY: posY,
		inAir:      false,
	}

	player := &Player{
		bullet:  playerBullet,
		img:     img,
		posX:    posX,
		posY:    posY,
		lives:   NUM_LIVES,
		isAlive: true,
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

func (p Player) Draw(screen *ebiten.Image, scaleX, scaleY float64) {
	op := &ebiten.DrawImageOptions{}
	if p.isAlive {
		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(float64(p.posX)*scaleX, float64(p.posY))
		if !p.bullet.inAir {
			p.bullet.bulletPosX = p.posX + 30
			p.bullet.bulletPosY = p.posY
		}

		screen.DrawImage(p.img, op)

		op.GeoM.Reset()
		if p.bullet.inAir {
			p.offsetBulletXY(0, -3)
			op.GeoM.Scale(0.25, 0.25)
			op.GeoM.Translate(float64(p.bullet.bulletPosX)*scaleX, float64(p.bullet.bulletPosY))
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

func (p Player) Die(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	if p.imgOnDeath != nil {
		op.GeoM.Translate(float64(p.posX), float64(p.posY))
		screen.DrawImage(p.imgOnDeath, op)
	}
}
