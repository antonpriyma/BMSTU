package models

import (
	"fmt"
	gonum "github.com/gonum/matrix/mat64"
	"math"
	"math/rand"
)

type Matrix []Vector

func (m Matrix) Dims() (r, c int) {
	return len(m), len(m)
}

func (m Matrix) At(i, j int) float64 {
	return m[i][j]
}

func (m Matrix) T() gonum.Matrix {
	return  m.Transponse()
}

func (m Matrix) VectorMult(v Vector) (res Vector) {
	if len(v) != len(m) {
		return res
	}

	for i, _ := range m[len(m)-1] {
		var temp Vector
		for j, _ := range m {
			temp = append(temp, m[j][i])
		}

		res = append(res, v.Scalar(temp))
	}
	return
}

func (m Matrix) MatrixMult(a Matrix) (res Matrix) {
	if len(m) != len(a) {
		return res
	}

	for i, _ := range m {
		res = append(res, a.VectorMult(m[i]))
	}

	return res
}

func (m Matrix) MatrixSum(m1 Matrix) Matrix {
	if len(m) != len(m1) {
		fmt.Println("Wrong matrix size")
		return nil
	}

	res := make(Matrix, len(m))
	for i, aa := range m {
		str := make([]float64, len(m))
		for j, elem := range aa {
			str[j] = elem + m1[i][j]
		}
		res[i] = str
	}

	return res
}

func (m Matrix) Norm() float64 {
	var res float64

	for _, aa := range m {

		for _, elem := range aa {
			res += elem * elem
		}

	}
	return math.Sqrt(res)
}

func (m Matrix) NormNew() float64 {
	sums := make([]float64, 0, len(m))
	for _, v := range m {
		sums = append(sums, v.SumElems())
	}

	return Vector(sums).Max()
}

func RandomMatrix(n int) Matrix {
	value := make(Matrix, n)

	for i := 0; i < n; i++ {
		value[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			value[i][j] = rand.Float64()
		}
	}
	return value
}


func (m Matrix) Transponse() Matrix {
	b := make(Matrix, len(m))

	for i := 0; i< len(m); i++{
		b[i] = make(Vector, len(m))
		for j := 0; j< len(m); j++ {
			b[i][j] = m[j][i]
		}
	}

	return b
}
func RandomDomMatrix(n int, max int) Matrix {
	value := make(Matrix, n)

	for i := 0; i < n; i++ {
		value[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			rand_num := rand.Int()%max - (max / 2)

			if i == j {
				rand_num += 2 * max
				rand_num *= n * 5

			}

			value[i][j] = float64(rand_num)
		}
	}

	return value
}

func sum(s []float64) float64 {
	sum := float64(0)

	for _, v := range s {
		sum += v
	}

	return sum
}

func (m Matrix) printMatrix(index []int) {
	for i := range m {
		for j := range m[i] {
			if m[i][index[j]] == 0 {
				fmt.Printf("[0] ")
			} else {
				fmt.Printf("[%v] ", m[i][index[j]])
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m Matrix) IsDominance() bool {
	for i := 0; i < len(m); i++ {
		diag := math.Abs(m[i][i])
		sum := float64(0)

		for j := 0; j < len(m); j++ {
			if i != j {
				sum += math.Abs(m[i][j])
			}
		}

		if diag <= sum {
			return false
		}
	}

	return true
}

func IdentityMatrix(rowCount, columnCount int) Matrix {
	matrix := make(Matrix, rowCount)

	for i := 0; i < rowCount; i++ {
		matrix[i] = make([]float64, columnCount)
		for j := 0; j < columnCount; j++ {
			matrix[i][j] = 0
		}
		matrix[i][i] = 1
	}

	return matrix
}

func (m Matrix) IsSymmetric() (bool, error) {
	if len(m) != len(m[0]) {
		return false, fmt.Errorf("not squere matrix")
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != m[j][i] {
				return false, nil
			}
		}
	}

	return true, nil
}

func (m Matrix) IsPositive() bool {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] <= 0 {
				return false
			}
		}
	}

	return true
}

func (m Matrix) Sub(matrix Matrix) Matrix {
	if len(m) != len(matrix) || len(m[0]) != len(matrix[0]) {
		return nil
	}

	res := make(Matrix, len(m))

	for i := 0; i < len(m); i++ {
		res[i] = make([]float64, len(m[i]))

		for j := 0; j <len(m[i]); j++ {
			res[i][j] = m[i][j] - matrix[i][j]
		}
	}
	return res
}

func (m Matrix) MulScalar(s float64) Matrix {
	value := make(Matrix, len(m))

	for i := 0; i < len(m); i++ {
		value[i] = make([]float64, len(m[i]))

		for j := 0; j < len(m[i]); j++ {
			value[i][j] = m[i][j] * s
		}
	}

	return value
}


func (m Matrix) Flat() Vector {
	res := make([]float64, len(m)*len(m[0]))

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			res[i*len(m[i])+j] = m[i][j]
		}
	}

	return res
}

func (m Matrix) NormE() float64 {
	return 0
}


//func (m Matrix) Det() Matrix {
//
//}