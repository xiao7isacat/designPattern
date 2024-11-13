package factory

type Gun struct {
	Name  string
	Power int
}

func (this *Gun) SetName(name string) {
	this.Name = name
}

func (this *Gun) SetPower(power int) {
	this.Power = power
}

func (this *Gun) GetName() string {
	return this.Name
}

func (this *Gun) GetPower() int {
	return this.Power
}
