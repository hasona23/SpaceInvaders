package Components

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimSprite struct {
	Atlas         *ebiten.Image
	Img           *ebiten.Image
	width, height int
	Animations    map[string]AnimationFrames
	Current       AnimationFrames
	Default       AnimationFrames
	IsInverted    int
}

func (sprite *AnimSprite) Origin(op *ebiten.DrawImageOptions) {
	s := sprite.Img.Bounds()
	op.GeoM.Translate(-float64(s.Dx())/2, -float64(s.Dy())/2)

}

func (AnimSprite) Init(img *ebiten.Image, Width, Height int) AnimSprite {
	return AnimSprite{img, ebiten.NewImage(Width, Height), Width, Height, make(map[string]AnimationFrames), AnimationFrames{}, AnimationFrames{}, 1}
}

func (a *AnimSprite) Add(name string, animtion AnimationFrames) {
	a.Animations[name] = animtion
}

func (a *AnimSprite) SetDefault(name string) {
	a.Default = a.Animations[name]

	//fmt.Println(a.Current.Name, " ", a.Default.Name)
	a.Play(a.Default.Name)
}

func (a *AnimSprite) Play(name string) {
	if a.Current.Name != name {
		a.Current = a.Animations[name]
		//fmt.Println("Chnaged")
	}
	a.Current.Update()
	if a.Current.IsEnd {
		a.ChangeTo(a.Default.Name)
	}

	a.Img = a.Atlas.SubImage(image.Rect(a.Current.row_current*a.width, a.Current.col_current*a.height, (a.Current.row_current+1)*a.width, (a.Current.col_current+1)*a.height)).(*ebiten.Image)
}
func (a *AnimSprite) ChangeTo(name string) {
	if !a.Current.IsEnd && a.Current.Name != a.Default.Name {
		return
	}
	a.Current = a.Animations[name]
}
func (a *AnimSprite) Update() {
	a.Play(a.Current.Name)
}

/*func (a *AnimSprite) PlayTillEnd(name string) {
	if a.Current.Name != name {
		a.Current = a.Animations[name]
		fmt.Println("Chnaged")
	}

	a.Current.Update()
	} else {
		a.ChangeTo(a.Default.Name)
	}
	a.Img = a.Atlas.SubImage(image.Rect(a.Current.row_current*a.width, a.Current.col_current*a.height, (a.Current.row_current+1)*a.width, (a.Current.col_current+1)*a.height)).(*ebiten.Image)
}*/

type AnimationFrames struct {
	row_min     int
	row_max     int
	row_current int
	col_min     int
	col_max     int
	col_current int
	IsEnd       bool
	Name        string
	Timer
}

func (a AnimationFrames) IsEmpty() bool {
	return AnimationFrames{} != a
}

func (AnimationFrames) Init(row_min, row_max, col_min, col_max, duration int, name string) AnimationFrames {
	return AnimationFrames{row_min: row_min, row_max: row_max, row_current: row_min, col_min: col_min, col_max: col_max, col_current: col_min, IsEnd: false, Name: name, Timer: Timer{}.Init(float64(duration))}
}

func (a *AnimationFrames) Update() {
	a.Timer.UpdateTimer()
	//fmt.Println(a.Timer.GetCurrentTime())
	if a.Ticked() {

		a.IsEnd = false
		a.row_current++
		//fmt.Println(" Row", a.row_max, " ", a.row_current)
		if a.row_current >= a.row_max {

			a.row_current = a.row_min
			a.col_current++

		}
		//fmt.Println(" Col ", a.col_max, " ", a.col_current)
		if a.col_current >= a.col_max {

			a.col_current = a.col_min
			a.IsEnd = true

		}

	}
}
