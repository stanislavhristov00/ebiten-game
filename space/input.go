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
 *	Holds key state.
 */
type keyState int

const (
	keyStateNone keyState = iota
	keyStatePressing
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
	keyStateMap map[string]keyState
}

/*
 *	Constructor for Input.
 */

func NewInput() *Input {
	return &Input{
		keyStateMap: *(&map[string]keyState{
			"LeftArrow":  keyState(keyStateNone),
			"RightArrow": keyState(keyStateNone),
		}),
	}
}

/*
 *	Updates the current input. Returns true if space is pressed, false otherwise.
 *	If game is in fullscreen, returns true if mouse is pressed, false otherwise.
 */

func (i *Input) Update() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) { // ebiten.IsKeyPressed
		return true
	}

	return false
}

/*
 *	Dir returns currently pressed direction and false if no direction key is pressed.
 *	If game is in fullscreen, checks for A and D keys as well.
 */

func (i *Input) Dir() (Dir, bool) {
	switch i.keyStateMap["LeftArrow"] { // inputUtil.isKeyJustPressed?
	case keyStateNone:
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			i.keyStateMap["LeftArrow"] = keyStatePressing
			return DirLeft, true
		}
	case keyStatePressing:
		if !ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			i.keyStateMap["LeftArrow"] = keyStateNone
			return DirNone, false
		}
	}

	switch i.keyStateMap["RightArrow"] {
	case keyStateNone:
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			i.keyStateMap["RightArrow"] = keyStatePressing
			return DirRight, true
		}
	case keyStatePressing:
		if !ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			i.keyStateMap["RightArrow"] = keyStateNone
			return DirNone, false
		}
	}

	return DirNone, false
}
