package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Привет, я отложенная функция defer1 в функции main, написана в самом начале main")
	}()
	fmt.Println("Привет, я main и я начался")
	fmt.Println("Начинается выполнение функции boo")
	boo()
	fmt.Println("Функция boo выполнена, начинается выполнение функции foo")
	foo()
	fmt.Println("Функция foo выполнена")
	fmt.Println("Привет, я main и я закончился")
	defer func() {
		fmt.Println("Привет, я отложенная функция defer1 в функции main, написана в самом конце main")
	}()
}
func boo() {
	fmt.Println("Привет, я boo и я начался")

	defer func() {
		fmt.Println("Привет, я отложенная функция defer1 в функции boo")
	}()
	defer func() {
		fmt.Println("Привет, я отложенная функция defer2 в функции boo")
	}()
	defer func() {
		fmt.Println("Привет, я отложенная функция defer3 в функции boo")
	}()
	fmt.Println("Привет, я boo и я закончился")

}
func foo() {
	fmt.Println("Привет, я foo и я начался")

	defer func() {
		fmt.Println("Привет, я отложенная функция defer в функции foo")
	}()
	fmt.Println("Привет, я foo и я закончился")

}
