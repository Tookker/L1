package main

import (
	"L1/task-24/point/point2d"
	"fmt"
)

//24.Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

func main() {
	point1 := point2d.NewPoint2d(4, 1)
	point2 := point2d.NewPoint2d(-2, 5)

	res, err := point1.Lenght(point2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Result =", res)
}
