package main

import (
	"fmt"
	"math"
)

const (
	a float64 = 0
)

var (
	b = math.Log(4)
)

func f(x float64) float64 {
	return x * math.Exp(x)
}

func sympsonMethod(f func(float64) float64, a float64, b float64, n float64) float64 {
	h := (b - a) / n

	var sum float64
	for i := 1; i < int(n+1); i++ {
		sum += f(a + float64(i)*h - h/2)
	}

	var sum1 float64
	for i := 1; i < int(n); i++ {
		sum1 += f(a + float64(i)*h)
	}

	return h / 6 * (f(a) + f(b) + 4*sum + 2*sum1)
}

func trapezoidMethod(f func(float64) float64, a float64, b float64, n float64) float64 {
	h := (b - a) / n
	var sum float64
	for i := 1; i < int(n); i++ {
		sum += f(a + float64(i)*h)
	}

	return h * ((f(a)+f(b))/2 + sum)
}

func richarson(h float64, h1 float64, k float64) float64 {
	return (h - h1) / (math.Pow(2, k) - 1)
}

func calculateIntegral(e float64, method func(f func(float64) float64, a float64, b float64, n float64) float64, k float64, f func(float64) float64, a float64, b float64) float64 {
	fmt.Printf("eps = %f\n", e)
	n := 1
	r := 10.0
	iter := 0
	h := 0.0

	for math.Abs(r) >= e {
		n = n * 2
		h1 := h
		h = method(f, a, b, float64(n))

		r = richarson(h, h1, k)
		iter++
	}

	fmt.Printf("iters: %d\n", iter)
	fmt.Printf("res: %.20f\n", h)
	fmt.Printf("richarson: %.20f\n", h+r)
	return h + r
}

func main() {
	eps := 0.001
	fmt.Println("\nSYMPSON:\n")
	sympRes := calculateIntegral(eps, sympsonMethod, 4, f, a, b)
	fmt.Println("\nTRAPEZOID:\n")
	trapRes := calculateIntegral(eps, trapezoidMethod, 2, f, a, b)

	fmt.Printf("diff = %.20f\n", math.Abs(sympRes-trapRes))
}
