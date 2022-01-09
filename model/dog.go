package model

type Dog struct {
	Name string
}

func NewDog(Name string) *Dog {
	name := Name
	if Name == "" {
		name = "NoName"
	}
	return &Dog{Name: name}
}

func (d Dog) Cry() string {
	return "Bow-wow"
}

func (d Dog) GetName() string {
	return d.Name
}
