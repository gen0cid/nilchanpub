package main

import "fmt"

type Profile struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("==================================")

	sliceInt := []int{}
	sliceInt = append(sliceInt, 12)
	fmt.Println("sliceInt:", sliceInt, len(sliceInt), cap(sliceInt))
	fmt.Println()

	fmt.Println("sliceInt[0]", sliceInt[0])
	fmt.Println()

	fmt.Println("sliceInt[0] до изменения:", sliceInt[0])
	sliceInt[0] = 52
	fmt.Println("sliceInt[0] после изменения:", sliceInt[0])
	fmt.Println()

	fmt.Println("==================================")

	sliceStr := make([]string, 0)
	sliceStr = append(sliceStr, "Привет!")
	fmt.Println("sliceStr:", sliceStr, len(sliceStr), cap(sliceStr))
	fmt.Println()
	fmt.Println("sliceStr[0]", sliceStr[0])

	fmt.Println("sliceStr[0] до изменения:", sliceStr[0])
	sliceStr[0] = "Как дела?"
	fmt.Println("sliceStr[0] после изменения:", sliceStr[0])
	fmt.Println()

	fmt.Println("==================================")

	sliceFloat64 := make([]float64, 2, 10)
	sliceFloat64 = append(sliceFloat64, 3.345)
	fmt.Println("sliceFloat64:", sliceFloat64, len(sliceFloat64), cap(sliceFloat64))
	fmt.Println()
	fmt.Println("sliceFloat64[2]", sliceFloat64[2])

	fmt.Println("sliceFloat64[2] до изменения:", sliceFloat64[2])
	sliceFloat64[2] = 32.2345
	fmt.Println("sliceFloat64[2] после изменения:", sliceFloat64[2])
	fmt.Println()

	fmt.Println("==================================")

	sliceStruct := make([]Profile, 3, 3)
	sliceStruct = append(sliceStruct, Profile{Name: "Петя", Age: 12})
	fmt.Println("sliceStruct:", sliceStruct, len(sliceStruct), cap(sliceStruct))
	fmt.Println()
	fmt.Println("sliceStruct[3]", sliceStruct[3])

	fmt.Println("sliceStruct[3] до изменения:", sliceStruct[3])
	sliceStruct[3] = Profile{Name: "Вася", Age: 8}
	fmt.Println("sliceStruct[3] после изменения:", sliceStruct[3])
	fmt.Println()

	fmt.Println("==================================")

	for k, v := range sliceInt {
		fmt.Println(k, ":", v)
	}
	fmt.Println("==================================")

	for k, v := range sliceStr {
		fmt.Println(k, ":", v)
	}
	fmt.Println("==================================")

	for k, v := range sliceFloat64 {
		fmt.Println(k, ":", v)
	}
	fmt.Println("==================================")

	for k, v := range sliceStruct {
		fmt.Println(k, ":", v)
	}
}
