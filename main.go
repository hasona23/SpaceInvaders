package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"os"

	entities "github.com/hasona23/SpaceInvaders/Entites"
	vec "github.com/hasona23/vec"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	player        entities.Player
	bulletManager vec.Vec[entities.Bullet]
	spawner       entities.EnemySpawner
	screenFilter  *ebiten.Image
	paused        bool
	state         int
}

const (
	normal = iota
	paused
	Death
)

var GameStates map[int]string
var source *text.GoTextFaceSource

func (g *Game) Init() {
	g.player = entities.Player{}.Init()
	g.bulletManager.Init()
	g.spawner = entities.EnemySpawner{}.Init()
	g.paused = false
	g.state = normal
	GameStates = map[int]string{
		0: "Normal",
		1: "Paused",
		2: "PlayerDied",
	}
	g.screenFilter, _, _ = ebitenutil.NewImageFromFile("D:/Code/Go/Projects/spaceInvaders/assets/images/pause.png")
	font_file, err := os.ReadFile("./assets/Minecraft.ttf")
	if err != nil {
		log.Fatal("Error reading font file : ", err)
	}
	source, err = text.NewGoTextFaceSource(bytes.NewReader(font_file))
	if err != nil {
		log.Fatal("Error loading font :", err)
	}
}
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
func (g *Game) Death() {
	g.state = Death
	if g.state == Death && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.Init()
	}
}
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) && g.player.Hp.GetHp() != 0 {
		if g.state == paused {
			g.state = normal
		}
		if g.state == normal {
			g.state = paused
		}
	}
	g.player.OnDeath(g.Death)

	if g.state == normal {
		g.player.Update(&g.bulletManager)
		g.spawner.Update(&g.bulletManager, &g.player)
		entities.UpdateBulletManager(&g.bulletManager, &g.spawner, &g.player)
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	op := text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)

	g.player.Draw(screen)
	g.spawner.Draw(screen)
	for i := range g.bulletManager.Arr {
		if !g.bulletManager.At(i).Dead {
			g.bulletManager.At(i).Draw(screen)
		}
	}

	op.GeoM.Translate(8, 8)
	text.Draw(screen, fmt.Sprintf("FPS : %.2f", ebiten.ActualFPS()), &text.GoTextFace{Source: source, Size: 16}, &op)
	op.GeoM.Translate(0, 21)
	text.Draw(screen, fmt.Sprintf("Hp : %v", g.player.Hp.GetHp()), &text.GoTextFace{Source: source, Size: 16}, &op)
	op.GeoM.Translate(0, 29)
	//text.Draw(screen, fmt.Sprintf("Score:ComingSoon"), &text.GoTextFace{Source: source, Size: 16}, &op)

	if g.state != normal {
		vector.DrawFilledRect(screen, 0, 0, 1280, 720, color.RGBA{60, 60, 60, 100}, true)
		textOptions := text.GoTextFace{Source: source, Size: 64}
		if g.state == paused {
			op.GeoM.Translate(640-96, 360-64)
			text.Draw(screen, "Paused", &textOptions, &op)
		} else {
			op.GeoM.Translate(640-64, 360-64)
			text.Draw(screen, "Died", &textOptions, &op)
			op.GeoM.Translate(-5*64, 64)
			text.Draw(screen, "Press Enter to Restart", &textOptions, &op)
		}
	}
	//ebitenutil.DebugPrint(screen, "Hello, World!")
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("FPS : %.4v", ebiten.ActualFPS()), 0, 15)
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Bullets : %v", g.bulletManager.Size), 0, 25)
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Bullets : %v", len(g.bulletManager.Arr)), 0, 35)
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Enemies : %v", g.spawner.Vec.Size), 0, 45)
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Enemies : %v / %v", g.player.Hp.GetHp(), g.player.Hp.GetMax()), 0, 60)
	//ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Game State : %v", GameStates[g.state]), 0, 70)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 720
}

func main() {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{}
	game.Init()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
