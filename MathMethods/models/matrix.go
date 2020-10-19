package models

import (
	"fmt"
	"math"
)

type Matrix [][]float64

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

	return Vector(sums).max()
}