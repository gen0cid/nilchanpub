package X

import "fmt"

var privateXGlobal string = "я приватная публичная переменныя"

func privateFunc() {
	fmt.Println("я приватная функция")
}

type privateXStruct struct{}

func (n privateXStruct) privateMethod() {
	fmt.Println("Я приватный метод")
}
