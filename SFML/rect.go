package main

import sf "github.com/manyminds/gosfml"

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {
	window := sf.NewRenderWindow(sf.VideoMode{screenWidth, screenHeight, 32}, "Rectangle", sf.StyleDefault, sf.DefaultContextSettings())
	window.SetVSyncEnabled(true)
	window.SetFramerateLimit(60)

	rect, _ := sf.NewRectangleShape()
	rect.SetSize(sf.Vector2f{30, 50})
	rect.SetFillColor(sf.ColorBlue())
	rect.SetOrigin(sf.Vector2f{15, 25})
	rect.SetPosition(sf.Vector2f{screenWidth / 2, screenHeight / 2})

	for window.IsOpen() {
		for event := window.PollEvent(); event != nil; event = window.PollEvent() {
			switch ev := event.(type) {
			case sf.EventKeyReleased:
				if ev.Code == sf.KeyEscape {
					window.Close()
				}
			case sf.EventClosed:
				window.Close()
			}
		}

		window.Clear(sf.ColorWhite())
		window.Draw(rect, sf.DefaultRenderStates())
		window.Display()
	}
}
