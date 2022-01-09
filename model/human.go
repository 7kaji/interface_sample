package model

import (
	"fmt"

	Crier "github.com/7kaji/interface_sample/interface"
)

type Human struct {
	Name string
}

func NewHuman(Name string) *Human {
	name := Name
	if Name == "" {
		name = "NoName"
	}
	return &Human{Name: name}
}

func (h Human) Touch(a Crier.Animal) string {
	animalName := a.GetName()
	return fmt.Sprintf("%s touched the %s then %s cried %s.", h.Name, animalName, animalName, a.Cry())
}
