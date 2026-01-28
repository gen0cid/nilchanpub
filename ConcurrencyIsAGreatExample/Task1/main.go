package main

import (
	airhumiditysensor "Concurrency/airHumiditySensor"
	airpressuresensor "Concurrency/airPressureSensor"
	seismicactivitysensor "Concurrency/seismicActivitySensor"
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	mtx := sync.Mutex{}

	var humidity []int
	var pressure []float64
	var seismic []int

	humidityContext, humididtyCancel := context.WithCancel(context.Background())
	pressureContext, pressureCancel := context.WithCancel(context.Background())
	seismicContext, seismicCancel := context.WithCancel(context.Background())

	humidityTransferPoint := airhumiditysensor.HumiditySensorvPool(humidityContext, 100)
	pressureTransferPoint := airpressuresensor.PressureSensorPool(pressureContext, 12345)
	seismicTransferPoint := seismicactivitysensor.SeismicSensorPool(seismicContext, 432)

	go func() {
		time.Sleep(3 * time.Second)
		humididtyCancel()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		pressureCancel()
	}()

	go func() {
		time.Sleep(7 * time.Second)
		seismicCancel()
	}()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range humidityTransferPoint {
			mtx.Lock()
			humidity = append(humidity, v)
			mtx.Unlock()
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()

		for v := range pressureTransferPoint {
			mtx.Lock()
			pressure = append(pressure, v)
			mtx.Unlock()
		}
	}()
	wg.Add(1)

	go func() {
		defer wg.Done()

		for v := range seismicTransferPoint {
			mtx.Lock()
			seismic = append(seismic, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()
	mtx.Lock()
	fmt.Println("влажность:", len(humidity))
	mtx.Unlock()

	mtx.Lock()
	fmt.Println("давление:", len(pressure))
	mtx.Unlock()

	mtx.Lock()
	fmt.Println("сейсмическая активность:", len(seismic))
	mtx.Unlock()

}
