package Components

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Img        *ebiten.Image
	IsInverted int
}

func (sprite *Sprite) Origin(op *ebiten.DrawImageOptions) {
	s := sprite.Img.Bounds()
	op.GeoM.Translate(-float64(s.Dx())/2, -float64(s.Dy())/2)

}
