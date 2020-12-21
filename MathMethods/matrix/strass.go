package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main(){
	sizes := []int{1, 100, 200, 300, 400, 500}
	res := ""

	classic := classicTest(sizes)
	for _, v := range classic {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	str1 := strassTest(sizes, false)
	for _, v := range str1 {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	str2 := strassTest(sizes, true)
	for _, v := range str2 {
		res += fmt.Sprintf("%.15f ", v)
	}
	res += "\n"

	fmt.Println(strings.ReplaceAll(res, ".", ","))
}


func classicTest(sizes []int) []float64 {
	res := make([]float64, len(sizes))

	for i, size := range sizes {
		m1, _ := fillRandomMatrix(size)
		m2, _ := fillRandomMatrix(size)

		start := time.Now()
		_, _ = m1.MulMatrix(m2)
		duration := float64(time.Since(start).Milliseconds())

		res[i] = duration
	}

	return res
}

func strassTest(sizes []int, concurrency bool) []float64 {
	res := make([]float64, len(sizes))

	for i, size := range sizes {
		m1, _ := fillRandomMatrix(size)
		m2, _ := fillRandomMatrix(size)

		start := time.Now()
		_, _ = m1.MulStrass(m2, concurrency)
		duration := float64(time.Since(start).Milliseconds())

		res[i] = duration
	}

	return res
}

func fillRandomMatrix(n int) (*Matrix, error) {
	value := make([][]float64, n)

	for i := 0; i < n; i++ {
		value[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			value[i][j] = rand.Float64() * 1000
		}
	}

	return NewMatrix(value)
}