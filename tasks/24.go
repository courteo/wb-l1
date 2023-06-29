package tasks

import (
	"fmt"
	"math"
)

type point struct {								// структура с двумя инкапсулированными полями
	x int
	y int
}

func createPoint(x, y int) point {				// создает структуру точка из двух входящих параметров
	return point{
		x: x,
		y: y,
	}
}

func getDistant(p1, p2 point) float64 {
	return math.Sqrt(float64((p1.x - p2.x) * (p1.x - p2.x) + (p1.y - p2.y)*(p1.y - p2.y)))		// вычисляет расстояние между двуями точками
}

func T24() {
	p1 := createPoint(2, 1)
	p2 := createPoint(5, 2)
	fmt.Println(getDistant(p1, p2))
}