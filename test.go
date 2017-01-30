package main

import (
	tl "github.com/badele/termloop"
	"github.com/badele/termloop/box"
)


type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?
		player.prevX, player.prevY = player.Position()
		switch event.Key { // If so, switch on the pressed key.
		case tl.KeyArrowRight:
			player.SetPosition(player.prevX+1, player.prevY)
		case tl.KeyArrowLeft:
			player.SetPosition(player.prevX-1, player.prevY)
		case tl.KeyArrowUp:
			player.SetPosition(player.prevX, player.prevY-1)
		case tl.KeyArrowDown:
			player.SetPosition(player.prevX, player.prevY+1)
		}
	}
}

func (player *Player) Collide(collision tl.Physical) {
	// Check if it's a Rectangle we're colliding with
	if _, ok := collision.(*tl.Rectangle); ok {
		player.SetPosition(player.prevX, player.prevY)
	}
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	startx := 26
	starty := 2
	mybox := box.NewFrame(startx, starty, 19, 7, tl.ColorRed, tl.ColorWhite, box.LineSingleBorder, true)
	mybox.SetTitle("Sign mode", box.AlignCenter)
	level.AddEntity(mybox)

	starty = starty+8
	mybox = box.NewFrame(startx, starty, 19, 7, tl.ColorMagenta, tl.ColorWhite, box.LineSingleBorder, false)
	mybox.SetTitle(" Single border ", box.AlignHCenter)
	level.AddEntity(mybox)

	starty = starty+8
	mybox = box.NewFrame(startx, starty, 19, 7, tl.ColorBlue, tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Double border ", box.AlignHCenter)
	level.AddEntity(mybox)

	startx = 0
	starty = 2
	mybox = box.NewFrame(startx, starty, 24, 23, tl.RgbTo256Color(120,120,120), tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Double border ", box.AlignHCenter)
	level.AddEntity(mybox)

	myhline := box.NewHLine(startx+1, starty+5,22,tl.RgbTo256Color(120,120,120), tl.ColorWhite,box.LineSingleBorder)
	level.AddEntity(myhline)

	mytext := box.NewTextArea(startx+1, starty,22,5, "This is a first text",tl.RgbTo256Color(120,120,120), tl.ColorWhite,box.AlignCenter)
	level.AddEntity(mytext)
	mytext = box.NewTextArea(startx+1, starty+1,22,5, "This is another text",tl.RgbTo256Color(120,120,120), tl.ColorWhite,box.AlignCenter)
	level.AddEntity(mytext)

	game.Screen().SetLevel(level)
	game.Start()
}