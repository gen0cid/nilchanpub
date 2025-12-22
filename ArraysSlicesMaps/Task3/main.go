package main

import "fmt"

func main() {
	map1 := map[int]int{}
	map1[1] = -5
	map1[15] = 1
	map1[2] = 2
	map1[4] = -3
	map1[7] = -5

	v1, ok1 := map1[15]

	fmt.Println("========================")
	fmt.Println("v1:", v1, "---", "ok1:", ok1)
	fmt.Println("========================")
	fmt.Println()

	fmt.Println("map1 до удаления значения по ключу '15':", map1)
	delete(map1, 15)
	fmt.Println("map1 после удаления значения по ключу '15':", map1)
	fmt.Println()

	map2 := make(map[string]bool)

	map2["Петя"] = true
	map2["Вася"] = true
	map2["Коля"] = true
	map2["Ваня"] = false
	map2["Степа"] = true

	v2, ok2 := map2["Владимир"]

	fmt.Println("========================")
	fmt.Println("v2:", v2, "---", "ok2:", ok2)
	fmt.Println("========================")
	fmt.Println()

	fmt.Println("map2 до удаления значения по ключу 'Коля':", map2)
	delete(map2, "Коля")
	fmt.Println("map2 после удаления значения по ключу 'Коля':", map2)
	fmt.Println()

}
