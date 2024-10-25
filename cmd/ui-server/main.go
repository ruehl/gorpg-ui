package main

import (
	"log"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (*Game) Update() error {
	return nil
}

func (*Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (*Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	const width = 320
	const height = 240
	return width, height
}

func main() {
	const windowWidth = 640
	const windowHeight = 480

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
