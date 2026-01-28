package seismicactivitysensor

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func seismicSensor(
	wg *sync.WaitGroup,
	ctx context.Context,
	id int,
	seismicTransferPoint chan<- int,
	coordinates string,
) {
	defer wg.Done()

	for {

		select {
		case <-ctx.Done():
			fmt.Println("------->>>>>>>Я датчик сейсмической акстивности номер:", id, "завершаю свою работу на сегодня!")
			return
		default:
			fmt.Println("Я датчик сейсмической акстивности номер:", id, "начал измерять давление в:", coordinates)
			time.Sleep(1 * time.Second)
			seismic := 3 + rand.Intn(4)
			fmt.Println("Я датчик сейсмической акстивности номер:", id, "результат давления:", seismic, "в:", coordinates)
			seismicTransferPoint <- seismic
			fmt.Println("Я датчик сейсмической акстивности номер:", id, "передал сведения о давлении в:", coordinates, " в метеослужбу:", seismic)
		}

	}

}

func SeismicSensorPool(ctx context.Context, n int) <-chan int {
	wg := &sync.WaitGroup{}
	seismicTransferPoint := make(chan int)
	for i := 1; i <= n; i++ {
		wg.Add(1)
		go seismicSensor(wg, ctx, i, seismicTransferPoint, coordinatesToSeismicSensor(i))
	}

	go func() {
		wg.Wait()
		close(seismicTransferPoint)
	}()

	return seismicTransferPoint
}

func coordinatesToSeismicSensor(i int) string {
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
