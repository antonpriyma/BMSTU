package main

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
