package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

type bankAccount struct {
	balance float64
}

func newbankAccount(bal float64) *bankAccount {
	return &bankAccount{
		balance: bal,
	}
}

func (b *bankAccount) withdraw(sum float64) error {
	if sum > b.balance {
		return errors.New("Вы не можете вывести суммму больше, чем у вас есть на балансе")
	}

	if rand.Float64() < 0.3 {
		return errors.New("ошибка при снятии наличных: технический сбой")
	}

	b.balance -= sum
	fmt.Println("Операция успешно совершена! Ващ баланс: ", b.balance)
	return nil
}

func (b *bankAccount) showBalance() error {

	fmt.Println("Ваш баланс:", b.balance)
	return nil
}

func (b *bankAccount) OnlinePayment(sum float64) error {
	if sum > b.balance {
		return errors.New("Вы не можете потратить больше, чем у вас есть на балансе")
	}
	if rand.Float64() < 0.3 {
		return errors.New("ошибка при онлайн-оплате: технический сбой")
	}

	b.balance -= sum
	fmt.Println("Вы произвели онлайн покупку на сумму:", sum)
	fmt.Println("Ваш текущий баланс:", b.balance)
	return nil
}
func simulateOperations(account *bankAccount) {
	for {
		fmt.Println("\n--- Выберите операцию ---")
		fmt.Println("1. Показать баланс")
		fmt.Println("2. Снять наличные")
		fmt.Println("3. Онлайн-оплата")
		fmt.Println("4. Выход")

		choise := 0
		fmt.Print("Ваш выбор:")
		fmt.Scan(&choise)

		switch choise {
		case 1:
			err := account.showBalance()
			if err != nil {
				fmt.Printf("❌ Ошибка: %s\n", err.Error())
			}
		case 2:
			sum := 0.0

			fmt.Print("Введите суммму для снятия:")
			fmt.Scan(&sum)
			err := account.withdraw(sum)

			if err != nil {
				fmt.Printf("❌ Ошибка: %s\n", err.Error())
			}
		case 3:
			sum := 0.0

			fmt.Print("Введите сумму:")
			fmt.Scan(&sum)
			err := account.OnlinePayment(sum)

			if err != nil {
				fmt.Printf("❌ Ошибка: %s\n", err.Error())
			}
		case 4:
			fmt.Println("Завершение работы...")
			fmt.Println("До свидания! :)")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
func main() {
	b := newbankAccount(1000)
	fmt.Println("🎉 Добро пожаловать в банковский симулятор!")
	simulateOperations(b)
}
