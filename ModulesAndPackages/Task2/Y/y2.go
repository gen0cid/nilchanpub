package Y

import "fmt"

var privateYGlobal string = "я приватная публичная переменныя Y и я вывожусь через отдельную функцию в файле Y1)))))"

func privateFunc() {
	fmt.Println("я приватная функция")
}

type privateYStruct struct{}

func (n privateYStruct) privateMethod() {
	fmt.Println("Я приватный метод")
}
