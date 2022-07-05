package main

type ak47 struct {
	name  string
	power int
}

// ak47 Builder
func newAk47() iGun {
	return &ak47{
		name:  "AK47 gun",
		power: 4,
	}
}

func (g *ak47) setName(name string) {
	g.name = name
}

func (g *ak47) getName() string {
	return g.name
}

func (g *ak47) setPower(power int) {
	g.power = power
}

func (g *ak47) getPower() int {
	return g.power
}
