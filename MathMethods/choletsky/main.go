package main

import (
	"fmt"
	"github.com/AntonPriyma/MathMethods/models"
	"math"
	"time"
)


var ATOL = 1e-12 // 0.00000000000000001

func CholeskyDecomposition(a models.Matrix) models.Matrix {
	L := make(models.Matrix, len(a))

	for i, _ := range L {
		L[i] = make(models.Vector, len(a))
	}

	for i, _ := range L {
		var temp float64

		for j := 0; j <= i; j++ {
			temp = 0
			if j == i {
				for k := 0; k < j; k++ {
					temp += math.Pow(L[j][k], 2)
				}
				L[j][j] = math.Sqrt(a[j][j] - temp)
			} else {
				for k := 0; k < j; k++ {
					temp += L[i][k] * L[j][k]
				}
				L[i][j] = (a[i][j] - temp)/L[j][j]
			}
		}
	}

//	fmt.Printf("L: %v", L)

	return L
}

func fillOPMatrix(n int) (models.Matrix) {
	m := models.RandomMatrix(n)


	for i := 0; i < len(m); i++ {
		for j := 0; j < i; j++ {
			m[i][j] = m[j][i]
		}
	}

	for i := 0; i < len(m); i++ {
		m[i][i] = m[i].SumElems() * 2
	}

	return m
}

func main() {
	for n:=100; n<500; n+=100 {

		matrix := fillOPMatrix(n)
		vector := models.RandomVector(n)

		timeS := time.Now()
		l := CholeskyDecomposition(matrix)
		y := Gauss(l, vector)
		_ = Gauss(l.Transponse(), y)
		fmt.Println("chol n=", n, " ",time.Now().Sub(timeS).Seconds())

		timeS = time.Now()
		_ = seideliMethod(matrix, vector)
		fmt.Println("seidel n=", n, " ",time.Now().Sub(timeS).Seconds())

		timeS = time.Now()
		_ = jacobiMethod(matrix, vector)
		fmt.Println("jaco n=", n, " ",time.Now().Sub(timeS).Seconds())

	//	fmt.Printf("%v", x)
	}
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
			//fmt.Printf("iters: %d\n", k)
			return xNew
		}

		for i := 0; i < len(x); i++ {
			x[i] = xNew[i]
		}
	}
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
			//fmt.Println("Iters:", k)

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
