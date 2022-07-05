package main

type musket struct {
	gun
}

// musket Builder
func newMusket() iGun {
	return &musket{
		gun: gun{ // Blueprint
			name:  "Musket gun",
			power: 1,
		},
	}
}
