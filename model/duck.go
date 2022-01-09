package model

type Duck struct {
	Name string
}

func NewDuck(Name string) *Duck {
	name := Name
	if Name == "" {
		name = "NoName"
	}
	return &Duck{Name: name}
}

func (d Duck) Cry() string {
	return "Quack"
}

func (d Duck) GetName() string {
	return d.Name
}
