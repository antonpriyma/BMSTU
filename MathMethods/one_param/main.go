package main

import (
	"fmt"
	"github.com/AntonPriyma/MathMethods/models"
	"time"
)

var ATOL = 1e-12 // 0.00000000000000001

func jacobiMethod(A models.Matrix, b models.Vector) []float64 {
	if !A.IsDominance() {
		return nil
	}

	var x = make(models.Vector, len(b))
	var xNew = make(models.Vector, len(b))

	var k int
	for ; ; k++ {
		for i := 0; i < len(A); i++ {
			xNew[i] = b[i]
			for j := 0; j < len(A[0]); j++ {
				if i != j {
					xNew[i] -= A[i][j] * x[j]
				}
			}
			xNew[i] /= A[i][i]
		}

		if delta := x.Sub(xNew); delta.Max() < ATOL {
			fmt.Printf("iters: %d\n", k)
			return xNew
		}

		for i := 0; i < len(x); i++ {
			x[i] = xNew[i]
		}
	}
}

func absoluteError(x, b models.Vector, A models.Matrix) {
	res := A.VectorMult(x)

	sub := res.Sub(b)
	abs := sub.Max()
	fmt.Printf("absolute error %0.20f\n", abs)
}

func seideliMethod(A models.Matrix, b models.Vector) []float64 {
	if !A.IsDominance() {
		return nil
	}

	var x = make(models.Vector, len(b))
	var xNew = make(models.Vector, len(b))
	for i := 0; i < len(b); i++ {
		x[i] = 0.
		xNew[i] = 0.
	}

	k := 0
	for {
		k++

		for i := 0; i < len(A); i++ {
			sum := float64(0)
			for j := 0; j < i; j++ {
				sum += A[i][j] * xNew[j]
			}
			for j := i + 1; j < len(A[0]); j++ {
				sum += A[i][j] * x[j]
			}
			xNew[i] = (b[i] - sum) / A[i][i]
		}

		if delta := x.Sub(xNew); delta.Max() < ATOL {
			fmt.Println("Iters:", k)

			return xNew
		}
		for i := 0; i < len(x); i++ {
			x[i] = xNew[i]
		}
	}
}


func main(){
	//n := 1000
	for n:=10; n<600; n+=100 {
		fmt.Printf("_______________%d_________________\n", n)
		var A = fillOPMatrix(n)
		var b = models.RandomVector(n)

		timeS := time.Now()
		x := jacobiMethod(A, b)
		fmt.Println("Time spent: ", time.Now().Sub(timeS).Microseconds())
		//fmt.Println("Solution:", x)

		absoluteError(x, b, A)

		timeS = time.Now()
		x = seideliMethod(A, b)
		fmt.Println("Time spent: ", time.Now().Sub(timeS).Microseconds())
		//fmt.Println("Solution:", x)
		absoluteError(x, b, A)

		startEigen := time.Now()
		_, _, tau, err := GetEigen(A)
		if err != nil {
			panic(err)
		}
		durationEigen := time.Since(startEigen).Milliseconds()
		fmt.Printf("Duration eigen: %d ms\n", durationEigen)
		fmt.Printf("Tau: %.15f\n", tau)

		timeS = time.Now()
		x, iterCount, _ := OneParameter(A, b, tau)
		fmt.Println("Time spent: ", time.Now().Sub(timeS).Microseconds())
		fmt.Println("Iters: ", iterCount)
		absoluteError(x, b, A)
	}

}
