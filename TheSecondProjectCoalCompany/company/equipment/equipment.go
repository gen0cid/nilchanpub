package equipment

import "fmt"

type Equipment struct {
	pickaxe     bool
	ventilation bool
	trolleys    bool
}

func newEquipment() Equipment {
	return Equipment{}
}

func (e *Equipment) buyPickaxe() {
	e.pickaxe = true
	fmt.Println("pickaxe is purchased")
}

func (e *Equipment) buyVentilation() {
	e.ventilation = true
	fmt.Println("ventilation is purchased")
}

func (e *Equipment) buyTrolleys() {
	e.trolleys = true
	fmt.Println("trolley is purchased")
}
