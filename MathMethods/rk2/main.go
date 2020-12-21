package main

import (
	"fmt"
	"github.com/AntonPriyma/MathMethods/models"
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack/lapack64"
	"gonum.org/v1/plot/vg"
	"math"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

var xMatrix models.Matrix

func main() {
	xV := models.Vector{-1,0,1}
	x := make(models.Matrix, len(xV))
	//x := models.Matrix{{0,-1,0}, {0,0,0}, {0,1,0}}
	y := models.Vector{1, 0,  1 }
	n := len(y)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"



	for i := 0; i < n; i++ {
		x[i] = make(models.Vector, len(xV))
		for j := 0; j < n; j++ {
			x[i][j] = math.Pow(xV[i], float64(j))
		}
	}

	xL, bL := createForLapack(x, y)


	v := Gauss(x,y)
	v1:=GaussSimple(x,y)
	v2 := lapack(uint32(len(y)), xL, bL)
	v2 = convertFromLapack(len(y), v2)
	//print(v)

	printSolve(v)
	printSolve(v1)
	printSolve(v2)
	err = plotutil.AddLinePoints(p,
		"First", makePoints(v, xV))
	if err != nil {
		panic(err)
	}

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}


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
func printSolve(v models.Vector) {
	for i, a:= range v{
		fmt.Printf("+%f*x^%d",a,i)
	}
	fmt.Println()
}

func makePoints(v models.Vector, coords models.Vector) plotter.XYs {
	xy := make(plotter.XYs, len(coords))
	for i, coord := range coords {
		var res float64
		for j, a := range v {
			res += a*math.Pow(coord, float64(j))
		}
		xy[i].X = coord
		xy[i].Y = res
	}

	return xy
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
			os.Exit(1)
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

		fmt.Println("Matrix A")
		//a.printMatrix(index)
		fmt.Println("Vector B")
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


func GaussSimple(a models.Matrix, b models.Vector) models.Vector {
	index := make([]int, len(a))
	for i := range index {
		index[i] = i
	}

	for i, _ := range a {
		pivot := a[i][index[i]]

		// если главный элемент строки равен 0, метод гаусса не работает
		if pivot == 0 {
			if b[i] == 0 {
				fmt.Println("система имеет множество решений")
			} else {
				fmt.Println("система не имеет решений")
			}
			os.Exit(1)
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

		fmt.Println("Matrix A")
		//a.printMatrix(index)
		fmt.Println("Vector B")
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

	fmt.Println("Lapack:")
	//printRowise(b, int(n), 1, int(n), false)
	fmt.Println()

	return b
}