package entities

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	components "github.com/hasona23/SpaceInvaders/Components"
	"github.com/hasona23/vec"
)

type Enemy struct {
	components.Transform
	components.AnimSprite
	components.Vel
	components.Timer
	components.Rect
	Dead bool
}

func (enemy Enemy) Init() Enemy {

	enemy = Enemy{components.Transform{X: 1300, Y: int(rand.Intn(680)) + 32, ScaleX: 5, ScaleY: 5},
		components.AnimSprite{}.Init(nil, 16, 16),
		components.Vel{Speed: 2},
		components.Timer{}.Init(90),
		components.Rect{}.Init(1300, enemy.Transform.Y-20, 70, 70),
		false,
	}
	var err error
	enemy.Atlas, _, err = ebitenutil.NewImageFromFile("./assets/images/tilemap_packed.png")
	if err != nil {
		log.Fatalf("Sorry failed import image as %v", err)
	}
	if enemy.Img == nil {
		log.Fatal("enemy image is null")
	}
	enemy.Add("move", components.AnimationFrames{}.Init(1, 3, 5, 6, 20, "move"))
	enemy.SetDefault("move")
	enemy.AnimSprite.IsInverted = -1
	return enemy
}

func (enemy *Enemy) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	enemy.Origin(&op)
	op.GeoM.Scale(float64(enemy.IsInverted)*enemy.ScaleX, enemy.ScaleY)

	op.GeoM.Translate(float64(enemy.Transform.X), float64(enemy.Transform.Y))

	screen.DrawImage(enemy.Img, &op)

}

func (enemy *Enemy) Update(bulletManager *vec.Vec[Bullet], player *Player, score *int) {
	if enemy.Dead {
		return
	}
	enemy.Transform.X -= enemy.Speed
	enemy.Rect.Move(enemy.Transform.X, enemy.Transform.Y)
	enemy.Timer.UpdateTimer()
	if enemy.Ticked() {
		bulletManager.PushBack(Bullet{}.Init(enemy.Transform.X, enemy.Transform.Y+2, 5, enemy.IsInverted, "enemy"))
	}
	enemy.AnimSprite.Update()
	if enemy.Rect.Intersect(player.Rect) {
		enemy.Dead = true
		player.Hp.Dec(1)
		*score += 100
	}
}

type EnemySpawner struct {
	vec.Vec[Enemy]
	components.Timer
}

func (spawner EnemySpawner) Init() EnemySpawner {
	spawner.Timer = components.Timer{}.Init(120)
	spawner.Vec.Init()

	return spawner
}
func (spawner *EnemySpawner) Update(bulletManager *vec.Vec[Bullet], player *Player, score *int) {
	spawner.Timer.UpdateTimer()
	if spawner.Ticked() {
		spawner.PushBack(Enemy{}.Init())
	}

	for i := range spawner.Arr {
		if !(spawner.At(i).Transform.X < 0) {
			spawner.At(i).Update(bulletManager, player, score)
		}

	}
	for i, e := range spawner.Arr {
		if e.Transform.X < 0 {
			err := spawner.PopIndex(i)
			if err != nil {
				log.Fatal("Error Deleting Enemy:", err)
			}
		}
		if e.Dead {
			err := spawner.PopIndex(i)
			if err != nil {
				log.Fatal("Error Deleting Enemy:", err)

			}
		}

	}

}
func (spawner *EnemySpawner) Draw(screen *ebiten.Image) {
	for i, e := range spawner.Arr {
		if !(e.Transform.X < 0) {
			spawner.At(i).Draw(screen)
		}
	}
}
