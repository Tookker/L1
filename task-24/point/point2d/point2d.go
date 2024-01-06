package point2d

import (
	"L1/task-24/point"
	"errors"
	"fmt"
	"math"
)

type Point2D struct {
	x float64
	y float64
}

var (
	errCast = "Wrong cast!"
)

func NewPoint2d(x float64, y float64) point.Point {
	return &Point2D{
		x: x,
		y: y,
	}
}

func (p *Point2D) Lenght(point point.Point) (float64, error) {
	castPoint, ok := point.(*Point2D)
	if !ok {
		fmt.Println(errCast)
		return 0, errors.New(errCast)
	}

	const powVal = 2

	return math.Sqrt(math.Pow(castPoint.x-p.x, powVal) + math.Pow(castPoint.y-p.y, powVal)), nil
}
