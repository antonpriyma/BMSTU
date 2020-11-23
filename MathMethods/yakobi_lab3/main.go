package main

import (
	"errors"
	"fmt"
	"github.com/AntonPriyma/MathMethods/models"
	"math"
)

var ATOL = 1e-12 // 0.00000000000000001
var A = models.Matrix{
	{10., -1., 2., 0.},
	{-1., 11., -1., 3.},
	{2., -1., 10., -1.},
	{0.0, 3., -1., 8.},
}
var b = models.Vector{6., 25., -11., 15.}

func scalarProduct(vec1, vec2 []float64) (float64, error) {
	if len(vec2) != len(vec1) {
		return 0, errors.New("different length")
	}
	var res float64 = 0
	for i := 0; i < len(vec1); i++ {
		res += vec1[i] * vec2[i]
	}
	return res, nil
}

func allClose(vec1, vec2 []float64) bool {
	f := 0
	for i := 0; i < len(vec1); i++ {
		if math.Abs(vec1[i]-vec2[i]) < ATOL {
			f++
		}
	}
	return f == len(vec1)
}

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

func Gauss(a models.Matrix, b models.Vector) models.Vector {
	index := make([]int, len(a))
	for i := range index {
		index[i] = i
	}

	for i, _ := range a {
		pivot := a[i][index[i]]

		// если главный элемент равен нулю, нужно найти другой
		// методом перестановки колонок в матрице
		if pivot == 0 {
			var kk int

			// двигаемся вправо от диаганаотного элемента, для поиска максимального по модулю элемента
			for k := i; k < len(a); k++ {
				if math.Abs(a[i][index[k]]) > pivot {
					kk = k
				}
			}

			// если удалось найти главный элемент
			if kk > 0 {
				// меняем местами колонки, так чтобы главный элемент встал в диагональ матрицы
				index[i], index[kk] = index[kk], index[i]
			}

			// получаем главный элемента, текущей строки из диагонали
			pivot = a[i][index[i]]
		}

		// если главный элемент строки равен 0, метод гаусса не работает
		if pivot == 0 {
			if b[i] == 0 {
				fmt.Println("система имеет множество решений")
			} else {
				fmt.Println("система не имеет решений")
			}
			return nil
		}

		for j := 0; j < len(a[i]); j++ {
			a[i][index[j]] /= pivot
		}
		b[i] /= pivot

		for k := i + 1; k < len(a); k++ {
			pivot = a[k][index[i]]
			for j := 0; j < len(a[i]); j++ {
				a[k][index[j]] = a[k][index[j]] - a[i][index[j]]*pivot
			}
			b[k] = b[k] - b[i]*pivot
		}

		//fmt.Println("Matrix A")
		//a.printMatrix(index)
		//fmt.Println("Vector B")
		//b.printVector()
	}

	x := make(models.Vector, len(b))

	for i := len(a) - 1; i >= 0; i-- {
		x[i] = b[i]

		for j := i + 1; j < len(a); j++ {
			x[i] = x[i] - (x[j] * a[i][index[j]])
		}
	}

	return x
}

func main() {
	//n := 1000
	//var A = models.RandomMatrix(n)
	//var b = models.RandomVector(n)

	x := jacobiMethod(A, b)
	//fmt.Println("Solution:", x)
	absoluteError(x, b, A)

	x = seideliMethod(A, b)
	//fmt.Println("Solution:", x)
	absoluteError(x, b, A)

	x = Gauss(A, b)
	absoluteError(x, b, A)

}
