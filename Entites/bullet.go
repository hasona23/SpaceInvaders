package entities

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	components "github.com/hasona23/SpaceInvaders/Components"
)

type Bullet struct {
	components.Transform
	components.Sprite
	components.Vel
	Dead bool
}

func (Bullet) Init(X, Y, Speed, IsInverted int) Bullet {
	bullet := Bullet{components.Transform{X: X, Y: Y, ScaleX: 3, ScaleY: 3}, components.Sprite{Img: nil, IsInverted: IsInverted}, components.Vel{Speed: 8}, false}
	Img, _, _ := ebitenutil.NewImageFromFile("./assets/images/tilemap_packed.png")
	bullet.Img = ebiten.NewImage(16, 16)
	bullet.Img.DrawImage(Img.SubImage(image.Rect(4*16, 4*16, 5*16, 5*16)).(*ebiten.Image), nil)
	return bullet
}
func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	bullet.Origin(&op)

	op.GeoM.Scale(float64(bullet.IsInverted)*bullet.ScaleX, bullet.ScaleY)

	op.GeoM.Translate(float64(bullet.X), float64(bullet.Y))

	screen.DrawImage(bullet.Img, &op)
}

func (bullet *Bullet) Update() {
	bullet.X += bullet.Speed * bullet.IsInverted
	bullet.Dead = (bullet.X > 1300 || bullet.X < -100)
}
