package main

import "fmt"

type Profile struct {
	Name string
	Age  int
}

func main() {
	arrInt := [5]int{1, 2, 3, 4, 5}
	arrString := [4]string{"hello", "goodbye", "Name", "Age"}
	arrBool := [4]bool{true, false, true, true}
	arrStruct := [3]Profile{
		{
			Name: "Гена",
			Age:  22,
		},
		{
			Name: "Федя",
			Age:  13,
		},
		{
			Name: "Паша",
			Age:  23,
		},
	}
	arrInt[3] += 345
	i := arrInt[3]
	fmt.Println("i:", i)

	arrString[1] += " my friend"
	s := arrString[1]
	fmt.Println("s:", s)

	arrBool[3] = false
	b := arrBool[3]
	fmt.Println("b:", b)

	arrStruct[2].Name = "Петя"
	STRUCT := arrStruct[2]
	fmt.Println("STRUCT:", STRUCT)

}
