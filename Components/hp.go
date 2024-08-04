package Components

type Hp struct {
	max int
	hp  int
}

func (Hp) Init(hp int) Hp {
	return Hp{max: hp, hp: hp}
}

func (hp *Hp) Inc(val int) {
	if hp.hp+val > hp.max {
		hp.hp = hp.max
	} else {
		hp.hp += val
	}
}
func (hp *Hp) Dec(val int) {
	if hp.hp-val <= 0 {
		hp.hp = 0
	} else {
		hp.hp -= val
	}
}

func (hp Hp) GetMax() int {
	return hp.max
}
func (hp Hp) GetHp() int {
	return hp.hp
}

func (hp Hp) OnDeath(action func()) {
	if hp.hp == 0 {
		action()
	}
}
