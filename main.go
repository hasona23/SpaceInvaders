package main

import (
	"bytes"
	"log"
	"os"
	"strconv"

	entities "github.com/hasona23/SpaceInvaders/Entites"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	player        entities.Player
	bulletManager entities.Vec[entities.Bullet]
	spawner       entities.EnemySpawner
	screenFilter  *ebiten.Image
	score         int
	highscore     int
	paused        bool
	state         int
}

var GameStates map[int]string
var source *text.GoTextFaceSource

func (g *Game) Init() {
	var err error
	g.player = entities.Player{}.Init()
	g.bulletManager.Init()
	g.spawner = entities.EnemySpawner{}.Init()

	source, err = text.NewGoTextFaceSource(bytes.NewReader(g.UiInit()))
	if err != nil {
		log.Fatal("Error loading font :", err)
	}
}

func (g *Game) Death() {
	g.state = Death
	if g.state == Death && inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.Init()
	}
}
func (g *Game) Update() error {

	if g.state == start {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.state = normal
		}
		return nil
	}

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
		g.spawner.Update(&g.bulletManager, &g.player, &g.score)
		entities.UpdateBulletManager(&g.bulletManager, &g.spawner, &g.player, &g.score)
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
	g.spawner.Draw(screen)
	for i := range g.bulletManager.Arr {
		if !g.bulletManager.At(i).Dead {
			g.bulletManager.At(i).Draw(screen)
		}
	}

	g.DrawUi(screen)

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
	if game.highscore < game.score {
		file, err := os.OpenFile("bestscore.txt", os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal("Error saving bestscore openign file :", err)
		}
		_, err = file.WriteString(strconv.Itoa(game.score))
		if err != nil {
			log.Fatal("Error wrting best score : ", err)
		}
		defer file.Close()
	}

}
