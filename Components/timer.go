package Components

type Timer struct {
	Time         float64
	current_time float64
}

func (Timer) Init(time float64) Timer {
	return Timer{Time: time, current_time: 0}
}

func (t Timer) GetCurrentTime() float64 {
	return t.current_time
}
func (timer *Timer) Ticked() bool {
	if timer.current_time >= timer.Time {
		timer.Reset()
		return true
	}
	return false
}
func (timer *Timer) Reset() {
	timer.current_time = 0
}
func (timer *Timer) UpdateTimer() {
	//timer.current_time += 1 / ebiten.ActualFPS()
	timer.current_time += 1
}
func (timer Timer) UpdateTime() Timer {
	timer.current_time++
	return timer
}
