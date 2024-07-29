package entities

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	components "github.com/hasona23/SpaceInvaders/Components"
	"github.com/hasona23/vec"
)

type Player struct {
	components.Transform
	components.Sprite
	components.Vel
}

func (player Player) Init() Player {

	player = Player{components.Transform{X: 40, Y: 40, ScaleX: 3, ScaleY: 3}, components.Sprite{Img: ebiten.NewImage(16, 16), IsInverted: 1}, components.Vel{Speed: 8}}
	Img, _, err := ebitenutil.NewImageFromFile("D:/Code/Go/Projects/spaceInvaders/assets/images/tilemap_packed.png")
	if err != nil {
		log.Fatalf("Sorry failed import image as %v", err)
	}
	player.Img.DrawImage(Img.SubImage(image.Rect(0*16, 4*16, 1*16, 5*16)).(*ebiten.Image), nil)
	if player.Img == nil {
		log.Fatal("player image is null")
	}
	return player
}
func (player Player) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	player.Origin(&op)
	op.GeoM.Scale(float64(player.IsInverted)*player.ScaleX, player.ScaleY)

	op.GeoM.Translate(float64(player.X), float64(player.Y))

	screen.DrawImage(player.Img, &op)

}
func (player *Player) Move(bulletManager *vec.Vec[Bullet]) {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		player.Y -= int(player.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		player.Y += int(player.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.X += int(player.Speed)
		player.IsInverted = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.X -= int(player.Speed)
		player.IsInverted = -1
	}

	if ebiten.IsKeyPressed(ebiten.KeyE) {
		bulletManager.PushBack(Bullet{}.Init(player.X, player.Y, 3, player.IsInverted))
	}

}
