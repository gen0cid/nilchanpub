package X

import "fmt"

var PublicXGlobal int = 100

func PublicFunc() {
	fmt.Println("я публичная функция")

}

type PublicXStruct struct{}

func (n PublicXStruct) PublicMethod() {
	fmt.Println("Я публичный метод")

	b := privateXStruct{}
	b.privateMethod()

	d := privateXGlobal
	fmt.Println("privateXGlobal", d)
}
