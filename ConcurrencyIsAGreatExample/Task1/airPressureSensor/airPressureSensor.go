package airpressuresensor

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func pressureSensor(
	wg *sync.WaitGroup,
	ctx context.Context,
	id int,
	pressureTransferPoint chan<- float64,
	coordinates string,
) {
	defer wg.Done()

	for {

		select {
		case <-ctx.Done():
			fmt.Println("------->>>>>>>Я датчик давления номер:", id, "завершаю свою работу на сегодня!")
			return
		default:
			fmt.Println("Я датчик давления номер:", id, "начал измерять давление в:", coordinates)
			time.Sleep(1 * time.Second)

			press := 740.0 + float64(rand.Intn(11))
			fmt.Println("Я датчик давления номер:", id, "результат давления:", press, "в:", coordinates)
			pressureTransferPoint <- press
			fmt.Println("Я датчик давления номер:", id, "передал сведения о давлении в:", coordinates, " в метеослужбу:", press)
		}

	}

}

func PressureSensorPool(ctx context.Context, n int) <-chan float64 {
	wg := &sync.WaitGroup{}
	pressureTransferPoint := make(chan float64)
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go pressureSensor(wg, ctx, i, pressureTransferPoint, coordinatesToPressureSensor(i))
	}

	go func() {
		wg.Wait()
		close(pressureTransferPoint)
	}()

	return pressureTransferPoint
}

func coordinatesToPressureSensor(i int) string {
	var ptm = map[int]string{
		1:  "Москва",
		2:  "Санкт-Петербург",
		3:  "Новосибирск",
		4:  "Екатеринбург",
		5:  "Казань",
		6:  "Нижний Новгород",
		7:  "Челябинск",
		8:  "Самара",
		9:  "Омск",
		10: "Ростов-на-Дону",
		11: "Уфа",
		12: "Красноярск",
		13: "Воронеж",
		14: "Пермь",
		15: "Волгоград",
	}

	val, ok := ptm[i]
	if !ok {
		return "Тула"
	}
	return val
}
