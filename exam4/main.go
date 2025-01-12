package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Car struct {
	Start     int
	End       int
	Capacity  int
	NumberCar int
}

type Order struct {
	Start       int
	NumberCar   int
	NumberOrder int
}

type OrderSet struct {
	Orders []Order
	Cars   []Car
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	numSets, _ := strconv.Atoi(strings.TrimSpace(line))

	var orderSets []OrderSet

	for numOrder := 0; numOrder < numSets; numOrder++ {

		line, _ := reader.ReadString('\n')
		strconv.Atoi(strings.TrimSpace(line))

		line, _ = reader.ReadString('\n')
		orderTimeStr := strings.Fields(line)

		var orders []Order

		for k, time := range orderTimeStr {
			orderTime, _ := strconv.Atoi(time)
			orders = append(orders, Order{orderTime, -1, k + 1})
		}

		line, _ = reader.ReadString('\n')
		numCars, _ := strconv.Atoi(strings.TrimSpace(line))

		var cars []Car

		for j := 0; j < numCars; j++ {

			line, _ = reader.ReadString('\n')
			carParamsStr := strings.Fields(line)

			start, _ := strconv.Atoi(carParamsStr[0])
			end, _ := strconv.Atoi(carParamsStr[1])
			capacity, _ := strconv.Atoi(carParamsStr[2])

			cars = append(cars, Car{start, end, capacity, j + 1})
		}

		orderSets = append(orderSets, OrderSet{orders, cars})
	}

	for numOrder, _ := range orderSets {
		sort.SliceStable(orderSets[numOrder].Cars, func(i, j int) bool {
			return orderSets[numOrder].Cars[i].Start < orderSets[numOrder].Cars[j].Start
		})
		sort.SliceStable(orderSets[numOrder].Orders, func(i, j int) bool {
			return orderSets[numOrder].Orders[i].Start < orderSets[numOrder].Orders[j].Start
		})

		for _, car := range orderSets[numOrder].Cars {
			count := 1
			for i, _ := range orderSets[numOrder].Orders {
				if orderSets[numOrder].Orders[i].NumberCar > 0 {
					continue
				}
				if orderSets[numOrder].Orders[i].Start >= car.Start && orderSets[numOrder].Orders[i].Start <= car.End {
					orderSets[numOrder].Orders[i].NumberCar = car.NumberCar
					count++
				}
				if count > car.Capacity {
					break
				}
			}
		}
		sort.SliceStable(orderSets[numOrder].Orders, func(i, j int) bool {
			return orderSets[numOrder].Orders[i].NumberOrder < orderSets[numOrder].Orders[j].NumberOrder
		})

		for _, order := range orderSets[numOrder].Orders {
			fmt.Printf("%d ", order.NumberCar)
		}
		fmt.Print("\n")
	}
}
