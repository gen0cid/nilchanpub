package car

type Car struct {
	name           string
	model          string
	hp             int
	engineCapacity float64
}

func NewCar(
	name string,
	model string,
	hp int,
	engineCapacity float64,
) Car {
	if name == "" {
		return Car{}
	}
	if model == "" {
		return Car{}
	}
	if hp < 50 || hp > 1500 {
		return Car{}
	}
	if engineCapacity < 0.5 || engineCapacity > 10.0 {
		return Car{}
	}
	return Car{
		name:           name,
		model:          model,
		hp:             hp,
		engineCapacity: engineCapacity,
	}
}
