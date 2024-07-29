package main

import (
	"fmt"
	"log"

	entities "github.com/hasona23/SpaceInvaders/Entites"
	vec "github.com/hasona23/vec"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

var player entities.Player
var bulletManager vec.Vec[entities.Bullet]

func Init() {
	player = player.Init()
	bulletManager.Init()
}
func (g *Game) Update() error {

	player.Move(&bulletManager)

	for i := range bulletManager.Arr {
		if !bulletManager.At(i).Dead {
			bulletManager.At(i).Update()
		}
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, "Hello, World!")
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS : %.4v", ebiten.ActualFPS()), 0, 10)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Bullets : %v", bulletManager.Size), 0, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Bullets : %v", len(bulletManager.Arr)), 0, 30)
	player.Draw(screen)

	for i := range bulletManager.Arr {
		if !bulletManager.At(i).Dead {
			bulletManager.At(i).Draw(screen)
		}
	}
	for _, b := range bulletManager.Arr {
		if b.Dead {
			err := bulletManager.PopElemnt(b)
			if err != nil {
				log.Fatal("Eror delte bullets", "  ", err)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Hello, World!")

	Init()

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
