package entities

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	components "github.com/hasona23/SpaceInvaders/Components"
)

type Bullet struct {
	components.Transform
	components.Sprite
	components.Vel
	components.Rect
	Dead    bool
	Shooter string
}

func (Bullet) Init(X, Y, Speed, IsInverted int, shooter string) Bullet {
	bullet := Bullet{components.Transform{X: X, Y: Y, ScaleX: 3, ScaleY: 3},
		components.Sprite{Img: nil, IsInverted: IsInverted}, components.Vel{Speed: Speed},
		components.Rect{}.Init(X, Y, 3*16, 3*16),
		false,
		shooter}
	Img, _, _ := ebitenutil.NewImageFromFile("./assets/images/tilemap_packed.png")
	bullet.Img = ebiten.NewImage(16, 16)
	bullet.Img.DrawImage(Img.SubImage(image.Rect(4*16, 4*16, 5*16, 5*16)).(*ebiten.Image), nil)
	return bullet
}
func (bullet *Bullet) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	bullet.Origin(&op)

	op.GeoM.Scale(float64(bullet.IsInverted)*bullet.ScaleX, bullet.ScaleY)

	op.GeoM.Translate(float64(bullet.Transform.X), float64(bullet.Transform.Y))

	screen.DrawImage(bullet.Img, &op)
}

func (bullet *Bullet) Update() {
	bullet.Transform.X += bullet.Speed * bullet.IsInverted
	bullet.Dead = (bullet.Transform.X > 1300 || bullet.Transform.X < -100)
	bullet.Rect.Move(bullet.Transform.X, bullet.Transform.Y)
}
func UpdateBulletManager(bulletManager *Vec[Bullet], enemyspawner *EnemySpawner, player *Player, score *int) {
	for i, b := range bulletManager.Arr {
		for j, b2 := range bulletManager.Arr {
			if b.Rect.Intersect(b2.Rect) && !b.Dead && !b2.Dead && b2 != b {
				bulletManager.At(j).Dead = true
				bulletManager.At(i).Dead = true
				*score += 10
			}

		}
		if b.Intersect(player.Rect) && b.Shooter != "player" && !b.Dead {
			bulletManager.At(i).Dead = true
			player.Dec(1)
			*score -= 30
		}
		for j, e := range enemyspawner.Arr {
			if b.Intersect(e.Rect) && !b.Dead && b.Shooter != "enemy" {
				enemyspawner.At(j).Dead = true
				bulletManager.At(i).Dead = true
				*score += 20

			}
		}
		if !bulletManager.At(i).Dead {
			bulletManager.At(i).Update()

		}

	}
	for k, b := range bulletManager.Arr {
		if b.Dead {
			err := bulletManager.PopIndex(k)
			if err != nil && err.Error() != "index out of range" {
				log.Fatal("Eror delte bullets", "  ", err)
			}

		}
	}

}
