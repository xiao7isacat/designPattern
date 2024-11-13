package factory

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			Name:  "Musket",
			Power: 4,
		},
	}
}
