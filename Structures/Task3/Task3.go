package main

import "fmt"

type Flat struct {
	Address string  //адрес не должен быть пустой строкой ""
	Square  float64 //площадь должна быть > 10
	Floor   int
	Price   int
}

func NewFlat(
	address string,
	square float64,
	floor, price int,
) Flat {
	fmt.Println("Валидирую адрес")
	if address == "" {
		fmt.Println("Адрес не прошел валидацию")
		return Flat{}
	}
	fmt.Println("Валидирую площать")
	if square <= 10 {
		fmt.Println("Площадь не прошла валидацию")
		return Flat{}
	}
	fmt.Println("Валидирую этаж квартиры")
	if floor <= 0 || floor > 100 {
		fmt.Println("Этаж не прошел валидацию")
		return Flat{}
	}
	fmt.Println("Валидирую стоимость квартиры")
	if price <= 0 {
		fmt.Println("Стоимость квартиры не прошла валидацию")
		return Flat{}
	}
	return Flat{
		Address: address,
		Square:  square,
		Floor:   floor,
		Price:   price,
	}
}

func (n *Flat) ChangePrice(NewPrice int) {
	if NewPrice <= 0 {
		fmt.Println("Стоимость квартиры не может быть не положительной")
		return
	}
	n.Price = NewPrice
}
func main() {
	flat1 := NewFlat("город тула, проспект Ленина 88 кв 58", 57.5, 6, 200000)
	fmt.Println(flat1)
	fmt.Println()

	flat1.ChangePrice(150000)
	fmt.Println(flat1)

}
