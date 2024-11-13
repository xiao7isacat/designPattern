package factory

type Ak47 struct {
	Gun
}

func NewAk47() IGun {
	return &Ak47{
		Gun: Gun{
			Name:  "ak47",
			Power: 7,
		},
	}
}
