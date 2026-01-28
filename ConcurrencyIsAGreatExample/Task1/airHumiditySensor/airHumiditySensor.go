package airhumiditysensor

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func humiditySensor(
	wg *sync.WaitGroup,
	ctx context.Context,
	id int,
	humidityTransferPoint chan<- int,
	coordinates string,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("------->>>>>>>Я датчик влажности воздуха номер:", id, "завершаю свою работу на сегодня!")
			return
		default:
			fmt.Println("Я датчик влажности воздуха номер:", id, "начал измерять влажность в:", coordinates)
			time.Sleep(1 * time.Second)

			humidity := 30 + rand.Intn(31)
			fmt.Println("Я датчик влажности воздуха номер:", id, "результат влажности:", humidity, "в:", coordinates)
			humidityTransferPoint <- humidity
			fmt.Println("Я датчик влажности воздуха номер:", id, "передал сведения в:", coordinates, " о влажности в метеослужбу:", humidity)
		}

	}

}

func HumiditySensorvPool(ctx context.Context, n int) <-chan int {
	wg := &sync.WaitGroup{}
	humidityTransferPoint := make(chan int)

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go humiditySensor(wg, ctx, i, humidityTransferPoint, coordinatesToHumiditySensor(i))
	}

	go func() {
		wg.Wait()
		close(humidityTransferPoint)
	}()

	return humidityTransferPoint
}

func coordinatesToHumiditySensor(i int) string {
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
