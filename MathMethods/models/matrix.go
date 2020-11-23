package models

import (
	"fmt"
	"math"
	"math/rand"
)

type Matrix []Vector

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
		sums = append(sums, Vector(v).SumElems())
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

	for i := 0; i < n; i++ {
		value[i][i] = sum(value[i]) * 2
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
