package main

import (
	"fmt"
	"math"
)

const (
	e = 0.001
	a = -3
	b = -2
)

var goldenRatio = (1 + math.Sqrt(5)) / 2

func f(x float64) float64 {
	return math.Pow(x, 4) - 5*math.Pow(x, 3) + 10*math.Pow(x, 2) - 5*x
}

func f_first(x float64) float64 {
	return 4*math.Pow(x, 3) - 15*math.Pow(x, 2) + 20*x - 5
}

func f_second(x float64) float64 {
	return 12*math.Pow(x, 2) - 30*x + 20
}

func binaryMethod() (float64, int) {
	var x float64
	var aLocal float64
	var bLocal float64
	aLocal = a
	bLocal = b

	i := 0
	for math.Abs(bLocal-aLocal) > e {
		x = (aLocal + bLocal) / 2

		f1 := f(x - e)
		f2 := f(x + e)

		if f1 < f2 {
			bLocal = x
		} else {
			aLocal = x
		}
		i++
	}

	x = (aLocal + bLocal) / 2
	return x, i
}

func goldenMethod() (float64, int) {
	var x float64
	var x1 float64
	var aLocal float64
	var bLocal float64
	aLocal = a
	bLocal = b

	i := 0
	for math.Abs(bLocal-aLocal) > e {
		x = bLocal - (bLocal-aLocal)/goldenRatio
		x1 = aLocal + (bLocal-aLocal)/goldenRatio
		f1 := f(x)
		f2 := f(x1)

		if f1 <= f2 {
			bLocal = x
		} else {
			aLocal = x
		}
		i++
	}

	return (aLocal + bLocal) / 2, i
}

func parabMethod() (float64, int) {

	var x float64
	var x1 float64
	x = (a + b) / 2
	i := 0
	for math.Abs(x1-x) > e {
		i++
		buf := x1
		x1 = x - f_first(x)/f_second(x)
		x = buf
	}

	if x1 > b {
		x1 = b
	}

	if x1 < a {
		x1 = a
	}
	return x1, i
}

func compareMethods(methods ...func() (float64, int)) {
	results := make([]struct {
		res float64
		x   float64
		i   int
	}, 0, 3)

	for _, method := range methods {
		x, i := method()
		results = append(results, struct {
			res float64
			x   float64
			i   int
		}{res: f(x), x: x, i: i})
	}

	for _, res := range results {
		fmt.Printf("%+v\n", res)
	}
}

func main() {
	compareMethods(goldenMethod, parabMethod, binaryMethod)
}
