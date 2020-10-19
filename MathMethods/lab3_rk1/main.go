package main

import (
	"fmt"
	"github.com/AntonPriyma/MathMethods/models"
	"math"
	"os"
)

func GaussMethod(a [][]float64, b []float64) (models.Vector, uint32) {
	var n uint32 = 2
	rows := make([]int, n)

	for i := 0; i < int(n); i++ {
		rows[i] = i
	}

	index := make([]int, len(a))
	for i := range index {
		index[i] = i
	}

	for i := 0; i < len(a); i++ {

		r := a[i][index[i]]

		var kk int

		for k := i; k < len(a); k++ {

			if math.Abs(a[i][index[k]]) > r {

				kk = k

			}

		}

		if kk > 0 {

			rows[index[i]], rows[index[kk]] = rows[index[kk]],

				rows[index[i]]

			index[i], index[kk] = index[kk], index[i]

		}

		r = a[i][index[i]]

		for k := i; k < len(a); k++ {

			if math.Abs(a[k][index[i]]) > r {

				kk = k

			}

		}

		if kk > 0 {

			a[i], a[kk] = a[kk], a[i]

			b[i], b[kk] = b[kk], b[i]

		}
		r = a[i][index[i]]

		if r == 0 {

			if b[i] == 0 {

				fmt.Println("система имеет множество решений")
			} else {

				fmt.Println("система не имеет решений")
			}
			os.Exit(1)

		}

		for j := 0; j < len(a[i]); j++ {

			a[i][index[j]] /= r

		}
		b[i] /= r

		for k := i + 1; k < len(a); k++ {
			r = a[k][index[i]]

			for j := 0; j < len(a[i]); j++ {

				a[k][index[j]] = a[k][index[j]] - a[i][index[j]]*r

			}

			b[k] = b[k] - b[i]*r

		}

	}

	var x = make([]float64, len(b))

	for i := len(a) - 1; i >= 0; i-- {
		x[i] = b[i]

		for j := i + 1; j < len(a); j++ {

			x[i] = x[i] - (x[j] * a[i][index[j]])

		}

	}

	return x, n

}

func main() {
	a := models.Matrix{{2, 7}, {9, 4},}
	b := models.Vector{3, 7}

	a1 := models.Matrix{{3, 7}, {4, 6},}
	b1 := models.Vector{3, 1}
	aa := a.MatrixSum(a1)
	bb := b.Sum(b1)
	rightPart := b1.Norm()/b.Norm() + a1.Norm()/a.Norm()
	rightNew := b1.NormNew()/b.NormNew() + a1.NormNew()/a.NormNew()


	aInverse := models.Matrix{{-55, 0}, {0, -55},}

	fmt.Printf("Right: %v\n", rightPart*a.Norm()*aInverse.Norm())
	fmt.Printf("RightNew: %v\n", rightNew*a.NormNew()*aInverse.NormNew())

	x1, _ := GaussMethod(a, b)
	x2, _ := GaussMethod(aa, bb)
	x3 := x2.Sub(x1)

	res := x3.Norm() / x1.Norm()
	fmt.Printf("Percent: %.2f\n", res*100)
	resNew := x3.NormNew() / x1.NormNew()
	fmt.Printf("Percent: %.2f", resNew*100)

}
