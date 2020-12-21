package main

import (
	"fmt"
	"github.com/AntonPriyma/MathMethods/models"
	"gonum.org/v1/gonum/mat"
)

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

var Eps = 1e-12
func OneParameter(a models.Matrix, f models.Vector, tau float64) (models.Vector, int, error) {
	if err := checkConstraints(a, f); err != nil {
		return nil, 0, err
	}

	id := models.IdentityMatrix(len(a), len(a[0]))
	p := id.Sub(a.MulScalar(tau))
	g := f.ScalarFloat(tau)
	x := make(models.Vector, len(f))
	xNext :=  make(models.Vector, len(f))
	iterCount := 0

	for {
		iterCount++

		px := p.VectorMult(x)
		xNext = px.Sum(g)

		if delta := x.Sub(xNext); delta.Max() < Eps {
			return xNext, iterCount, nil
		}

		x = xNext
	}
}

func checkConstraints(a models.Matrix, f models.Vector) error {
	if len(a) != len(a[0]) {
		return fmt.Errorf("not squere")
	}
	if len(a) != len(f) {
		return fmt.Errorf("error deminsion")
	}
	if isSymmetric, err := a.IsSymmetric(); err != nil || !isSymmetric {
		return fmt.Errorf("not symmetric")
	}
	if !a.IsPositive() {
		return fmt.Errorf("not positive")
	}
	return nil
}

func GetEigen(m models.Matrix) (max float64, min float64, opt float64, err error) {
	eigenvalues, err := getEigen(m)
	if err != nil {
		return 0, 0, 0, err
	}

	max, min = eigenvalues[0], eigenvalues[0]
	for _, v := range eigenvalues {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min, 2 / (min + max), nil
}

func getEigen(m models.Matrix) ([]float64, error) {
	var eigen mat.Eigen

	dm := mat.NewDense(len(m), len(m[0]), m.Flat())
	eigen.Factorize(dm, mat.EigenLeft)
	rawEigenValues := eigen.Values(nil)

	res := make([]float64, len(rawEigenValues))
	for i, v := range rawEigenValues {
		res[i] = real(v)
		if imag(v) != 0 {
			return nil, fmt.Errorf("error complex")
		}
	}
	return res, nil
}