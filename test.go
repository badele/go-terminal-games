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

	//termx, termy := game.InitialTermSize()
	//termx, termy := 194,61
	termx, termy := 80,25

	boxtextheight := 8
	clockwidth := 30
	clockheight := 12

	mybox := box.NewFrame(termx - clockwidth, 0, clockwidth, clockheight, tl.ColorBlue, tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Inventory ", box.AlignHCenter)
	mybox.LevelFollow(level)
	level.AddEntity(mybox)

	posy := clockheight
	width := 30
	mybox = box.NewFrame(termx - width, posy, width, termy - posy - boxtextheight, tl.ColorBlue, tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Other ", box.AlignHCenter)
	mybox.LevelFollow(level)
	level.AddEntity(mybox)

	mybox = box.NewFrame(0, termy-boxtextheight, termx, boxtextheight, tl.ColorBlue, tl.ColorWhite, box.LineDoubleBorder, false)
	mybox.SetTitle(" Chat ", box.AlignHCenter)
	mybox.LevelFollow(level)
	level.AddEntity(mybox)

	mytext := box.NewTextArea(1, termy-boxtextheight+1,termx-2,termy-2, "",tl.ColorBlue, tl.ColorWhite,box.AlignNone)
	mytext.SetTypewriterDuration(100)
	mytext.SetText("This story takes place in the year 3000, You have a spacecraft for navigate",box.AlignNone)
	mytext.LevelFollow(level)
	level.AddEntity(mytext)

	player := Player{
		Entity: tl.NewEntity(1, 1, 1, 1),
		level:  level,
	}
	// Set the character at position (0, 0) on the entity.
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)

	game.Screen().SetLevel(level)
	game.Start()
}
