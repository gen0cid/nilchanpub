package main

import "fmt"

type User struct {
	Name     string  //""
	Surname  string  //""
	Age      int     //0 - 150
	Rating   float64 //0.0 - 10.0
	IsClosed bool    //false
	Height   float64 //0.0-3.0
	Weight   float64 //0.0-300.0
}

func NewUser(
	name string,
	surname string,
	age int,
	rating float64,
	isClosed bool,
	height float64,
	weight float64,
) User {
	if name == "" {
		return User{}
	}
	if surname == "" {
		return User{}
	}
	if age < 0 || age > 150 {
		return User{}
	}
	if rating < 0.0 || rating > 10.0 {
		return User{}
	}
	if height < 0.0 || height > 3.0 {
		return User{}
	}
	if weight < 0.0 || weight > 300.0 {
		return User{}
	}
	return User{
		Name:    name,
		Surname: surname,
		Age:     age,
		Rating:  rating,
		Height:  height,
		Weight:  weight,
	}
}

func (n *User) OpenAccount() {
	n.IsClosed = false
}

func (n *User) CloseAccount() {
	n.IsClosed = true
}

func (n *User) ChangeAge(NewAge int) {
	if NewAge >= 0 && NewAge <= 150 {
		n.Age = NewAge
	}

}

func (n *User) ChangeHeight(NewHeight float64) {
	if NewHeight >= 0.0 && NewHeight <= 3.0 {
		n.Height = NewHeight
	}
}

func (n *User) ChangeWeight(NewWeight float64) {
	if NewWeight >= 0.0 && NewWeight <= 300.0 {
		n.Weight = NewWeight
	}
}

func (n User) greeteeng() {
	fmt.Println("Всем привет, меня зовут:", n.Name)
	fmt.Println("Моя фамилия:", n.Surname)
	fmt.Println("Мой возраст:", n.Age)
	fmt.Println("Мой рейтинг:", n.Rating)
	fmt.Println()
}

func (n User) goodbye() {
	fmt.Println("Всем пока, меня звали:", n.Name)
	fmt.Println("Моя фамилия была:", n.Surname)
	fmt.Println("Мой возраст, был:", n.Age)
	fmt.Println("Мой рейтинг был:", n.Rating)
	fmt.Println()
}
func (n *User) ChangeName(newName string) {
	if newName == "" {
		fmt.Println("Имя не может быть пустой строкой")
	} else {
		n.Name = newName
		fmt.Println("Имя уcпешно изменено")
	}
}
func (n *User) ChangeSurname(newSurname string) {
	if newSurname == "" {
		fmt.Println("Фамилия не может быть пустой строкой")
	} else {
		n.Surname = newSurname
		fmt.Println("Фамилия упешно изменена")
	}
}
func (n *User) RatingUp(rating float64) {
	if ((n.Rating + rating) >= 0) && ((n.Rating + rating) <= 10.0) {

		fmt.Println("Рейтинг пользователя:", n.Rating)

		n.Rating += rating

		fmt.Println("Рейтинг не выходит за границы, все в порядке")
		fmt.Println("Увеличиваем значение параметра 'рейтинг' на", rating)
		fmt.Println("Рейтинг пользователя:", n.Rating)
	} else {
		fmt.Println("Рейтинг  выходит за границы ")
		fmt.Println()
	}
}

func (n *User) RatingDown(rating float64) {
	if ((n.Rating - rating) >= 0) && ((n.Rating - rating) <= 10.0) {

		fmt.Println("Рейтинг пользователя:", n.Rating)

		n.Rating -= rating

		fmt.Println("Рейтинг не выходит за границы, все в порядке")
		fmt.Println("Уменьшаем значение параметра 'рейтинг' на", rating)
		fmt.Println("Рейтинг пользователя:", n.Rating)
	} else {
		fmt.Println("Рейтинг  выходит за границы ")
		fmt.Println()
	}
}

func main() {
	user1 := NewUser(
		"Вася",
		"Иванов",
		13,
		5.5,
		false,
		1.75,
		74.3)
	fmt.Println("user1:", user1)
}
