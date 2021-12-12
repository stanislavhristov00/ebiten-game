package space

import "github.com/hajimehoshi/ebiten/v2"

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
		bulletPosX: 0,
		bulletPosY: 0,
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

func (p Player) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, count int) {
	if p.isAlive {
		op.GeoM.Translate(float64(p.posX), float64(p.posY))

		if !p.bullet.inAir {
			p.bullet.bulletPosX = p.posX
			p.bullet.bulletPosY = p.posY
		}

		screen.DrawImage(p.img, op)

		if p.bullet.inAir {
			p.offsetBulletXY(0, -1)
			op.GeoM.Translate(float64(p.bullet.bulletPosX), float64(p.bullet.bulletPosY))
			screen.DrawImage(p.bullet.img, op)
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
