package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	start = iota
	normal
	paused
	Death
)

func (g *Game) DrawUi(screen *ebiten.Image) {
	op := text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)
	if g.state == start {
		op.GeoM.Translate(640-64*5, 360-200)
		text.Draw(screen, fmt.Sprintf("BestScore : %v", g.highscore), &text.GoTextFace{Source: source, Size: 64}, &op)
		op.GeoM.Translate(0, 260)
		text.Draw(screen, "Press Enter to Start", &text.GoTextFace{Source: source, Size: 32}, &op)

		return
	}
	op.GeoM.Translate(8, 8)
	text.Draw(screen, fmt.Sprintf("FPS : %.2f", ebiten.ActualFPS()), &text.GoTextFace{Source: source, Size: 16}, &op)
	op.GeoM.Translate(0, 21)
	text.Draw(screen, fmt.Sprintf("Hp : %v", g.player.Hp.GetHp()), &text.GoTextFace{Source: source, Size: 16}, &op)
	op.GeoM.Translate(0, 29)
	text.Draw(screen, fmt.Sprintf("Score: %v", g.score), &text.GoTextFace{Source: source, Size: 16}, &op)

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
}
func (g *Game) UiInit() []byte {
	file, err := os.OpenFile("bestscore.txt", os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("Error openning file of score : ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scoretext := ""
	for scanner.Scan() {
		scoretext = scanner.Text()
	}

	if scoretext == "" {
		g.highscore = 0
	} else {
		g.highscore, err = strconv.Atoi(scoretext)
		if err != nil {
			log.Fatal("Error converting score text to int :", err)
		}
	}

	g.paused = false
	g.state = start
	GameStates = map[int]string{
		0: "Start",
		1: "Normal",
		2: "Paused",
		3: "PlayerDied",
	}
	g.screenFilter, _, _ = ebitenutil.NewImageFromFile("D:/Code/Go/Projects/spaceInvaders/assets/images/pause.png")
	font_file, err := os.ReadFile("./assets/Minecraft.ttf")
	if err != nil {
		log.Fatal("Error reading font file : ", err)
	}

	return font_file
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
