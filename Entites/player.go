package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	components "github.com/hasona23/SpaceInvaders/Components"
	"github.com/hasona23/vec"
)

type Player struct {
	components.Transform
	//components.Sprite
	components.AnimSprite
	components.Vel
	components.Timer
	components.Rect
	components.Hp
}

func (player Player) Init() Player {

	player = Player{components.Transform{X: 40, Y: 40, ScaleX: 3, ScaleY: 3},
		//components.Sprite{Img: ebiten.NewImage(16, 16), IsInverted: 1},
		components.AnimSprite{}.Init(nil, 16, 16),
		components.Vel{Speed: 10},
		components.Timer{}.Init(20),
		components.Rect{}.Init(40, 40, 48, 48),
		components.Hp{}.Init(5),
	}

	var err error
	player.Atlas, _, err = ebitenutil.NewImageFromFile("D:/Code/Go/Projects/spaceInvaders/assets/images/tilemap_packed.png")
	if err != nil {
		log.Fatalf("Sorry failed import image as %v", err)
	}
	if player.Atlas == nil {
		log.Fatal("player image is null")
	}
	Idle := components.AnimationFrames{}.Init(0, 1, 4, 5, 1, "idle")
	player.AnimSprite.Add(Idle.Name, Idle)
	Run := components.AnimationFrames{}.Init(0, 2, 4, 5, 10, "run")
	player.AnimSprite.Add(Run.Name, Run)
	Shoot := components.AnimationFrames{}.Init(2, 3, 4, 5, 10, "shoot")
	player.AnimSprite.Add(Shoot.Name, Shoot)
	player.AnimSprite.SetDefault(Idle.Name)

	return player
}
func (player Player) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	player.Origin(&op)
	op.GeoM.Scale(float64(player.IsInverted)*player.ScaleX, player.ScaleY)

	op.GeoM.Translate(float64(player.Transform.X), float64(player.Transform.Y))

	screen.DrawImage(player.Img, &op)

}
func (player *Player) Update(bulletManager *vec.Vec[Bullet]) {

	player.Move()
	player.shoot(bulletManager)
	player.AnimSprite.Update()

}
func (player *Player) shoot(bulletManager *vec.Vec[Bullet]) {
	player.UpdateTimer()

	if ebiten.IsKeyPressed(ebiten.KeyE) && player.Ticked() {
		bulletManager.PushBack(Bullet{}.Init(player.Transform.X, 5+player.Transform.Y, 15, player.IsInverted, "player"))
		player.ChangeTo("shoot")
	}

}
func (player *Player) Move() {
	player.Rect.Move(player.Transform.X, player.Transform.Y)

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		player.Transform.Y -= int(player.Speed)

	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		player.Transform.Y += int(player.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.Transform.X += int(player.Speed)
		player.IsInverted = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.Transform.X -= int(player.Speed)
		player.IsInverted = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		player.ChangeTo("run")
	} else {
		if player.Current.Name == "run" {
			player.Current = player.Animations["idle"]
		}
	}
}
