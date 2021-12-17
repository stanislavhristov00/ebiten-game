package space

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

/*
 *	Represent direction.
 */
type Dir int

const (
	DirRight Dir = iota
	DirLeft
	DirNone
)

/*
 *	Returns a string representing direction.
 */

func (dir Dir) String() string {
	switch dir {
	case DirRight:
		return "Right"
	case DirLeft:
		return "Left"
	case DirNone:
		return "None"
	}

	panic("not reachable")
}

/*
 *	Returns a value from -1,0,1 corresponding to the X axis.
 */

func (dir Dir) DirToValue() (x int) {
	switch dir {
	case DirRight:
		return 1
	case DirLeft:
		return -1
	case DirNone:
		return 0
	}
	panic("not reach")
}

// Input represents the current key states.
type Input struct {
}

/*
 *	Constructor for Input.
 */

func NewInput() *Input {
	return &Input{}
}

/*
 *	Updates the current input. Returns true if space is pressed, false otherwise.
 */

func (i *Input) Update() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

/*
 *	Dir returns currently pressed direction and false if no direction key is pressed.
 */

func (i *Input) Dir() (Dir, bool) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		return DirRight, true
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		return DirLeft, true
	}

	return DirNone, false
}
