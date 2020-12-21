package main

import (
	"fmt"
	"github.com/AntonPriyma/MathMethods/models"
	"github.com/gonum/blas"
	"github.com/gonum/blas/blas64"
	"github.com/gonum/lapack/lapack64"
	"math"
	gonum "github.com/gonum/matrix/mat64"
)

func LU(m models.Matrix) (l models.Matrix, u models.Matrix) {
	n := len(m)

	l = make(models.Matrix,n)
	u = make(models.Matrix,n)

	for i,_ := range m {
		l[i] = make(models.Vector, n)
		u[i] = make(models.Vector, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < i {
				l[j][i] = 0
			} else {
				l[j][i] = m[j][i]
				for k := 0; k < i; k++ {
					l[j][i] = l[j][i] - l[j][k]*u[k][i]
				}
			}
		}
		for j := 0; j < n; j++ {
			if j < i {
				u[i][j] = 0
			} else if j == i {
				u[i][j] = 1
			} else {
				u[i][j] = m[i][j] / l[i][i]
				for k := 0; k < i; k++ {
					u[i][j] = u[i][j] - ((l[i][k] * u[k][j]) / l[i][i])
				}
			}
		}
	}

	return l,u
}


func main() {
	n := 5
	matrix := fillOPMatrix(n)
	vector := models.RandomVector(n)

	l, u := LU(matrix)

	mLU := l.MatrixMult(u)


	fmt.Printf("m: %v\n",matrix)
	fmt.Printf("L: %v\n",l)
	fmt.Printf("U: %v\n",u)
	fmt.Printf("mLU: %v\n",mLU)

	fmt.Printf("m - LU: %f\n",matrix.Sub(mLU).Norm())

	fmt.Printf("m - mT: %f\n",matrix.Transponse().Sub(matrix).Norm())

	if matrix.Transponse().Sub(matrix).NormNew() < 1 {
		fmt.Printf("u-lT: %f\n",u.Sub(l.Transponse()).NormE())
	}

	// Solve ax = b
	y := Gauss(l,vector)
	x := Gauss(matrix, vector)

	fmt.Printf("X LU: %v\n",x)

	mL, vL := createForLapack(matrix, vector)
	xLRaw := lapack(uint32(len(matrix)), mL, vL)
	xL := convertFromLapack(len(matrix), xLRaw)
	fmt.Printf("X Lapack: %v\n",xL)

	xG := Gauss(matrix, vector)
	fmt.Printf("X G: %v\n",xG)



	fmt.Println(y)

	/// DET
	det :=  gonum.Det(matrix)
	fmt.Printf("det: %20f\n",det)

	detL :=  gonum.Det(l)
	detU :=  gonum.Det(u)
	fmt.Printf("det: %20f\n",detL*detU)

	//inverse
	d:=&gonum.Dense{}
	_ = d.Inverse(matrix)
	fmt.Printf("A-1: %v\n",d)

	reverseLU := make([][]float64, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i < j {
				sum := 0.0
				for k := i+1; k < n; k++ {
					sum += u[i][k]*u[k][j]
				}
				reverseLU[i][j] = -sum/u[i][i]
			}
			if i == j {
				sum := 0.0
				for k := i+1; k < n; k++ {
					sum += u[i][k]*u[k][j]
				}
				reverseLU[i][j] = (1 - sum)/u[i][i]
			} else {
				sum := 0.0
				for k := j+1; k < n; k++ {
					sum += reverseLU[i][k]*l[k][j]
				}
				reverseLU[i][j] = -sum
			}
		}
	}

	d=&gonum.Dense{}
	_ = d.Inverse(l.MatrixMult(u))
	fmt.Printf("LU-1: %v\n",d)
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

func fillOPMatrix(n int) models.Matrix {
	m := models.RandomMatrix(n)


	for i := 0; i < len(m); i++ {
		for j := 0; j < i; j++ {
			m[i][j] = m[j][i]
		}
	}

	return m
}


func convertFromLapack(n int, v2 []float64) models.Vector {
	res := make(models.Vector, 0, n)
	for i, v := range v2 {
		if i%n == 0 {
			res = append(res, v)
		}
	}

	return res
}

func createForLapack(m models.Matrix, b models.Vector) (x models.Vector, bRes models.Vector) {
	for _, v := range m {
		x = append(x, v...)
	}

	for _, v := range  b{
		temp := make(models.Vector, len(b))
		temp[0] = v
		bRes = append(bRes, temp...)
	}

	return
}

func lapack(n uint32, a, b []float64) []float64 {
	aa := blas64.General{
		Rows:   int(n),
		Cols:   int(n),
		Data:   a,
		Stride: int(n),
	}

	bb := blas64.General{
		Rows:   int(n),
		Cols:   int(n),
		Data:   b,
		Stride: int(n),
	}
	ipiv := make([]int, n)

	lapack64.Getrf(aa, ipiv)
	lapack64.Getrs(blas.NoTrans, aa, bb, ipiv)

	//fmt.Println("Lapack:")
	//printRowise(b, int(n), 1, int(n), false)
	fmt.Println()

	return b
}