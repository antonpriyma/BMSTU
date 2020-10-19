package main

import (
	"fmt"
	"github.com/AntonPriyma/MathMethods/models"
	"github.com/alex-ant/gomath/gaussian-elimination"
	"github.com/alex-ant/gomath/rational"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {

	a := make(models.Matrix,3)
	b := models.Vector{}
	t := make(models.Matrix,3)

	m := make([][]float64, 3)
	m[0] = []float64{}
	m[1] = []float64{}
	m[2] = []float64{}

	rand.Seed(time.Now().Unix())
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			r := rand.Float64()
			a[i] = append(a[i], r)
			m[i] = append(m[i], r)
			t[i] = append(t[i], r)
		}
		d := rand.Float64()
		b = append(b, d)
		m[i] = append(m[i], d)
	}

	index := make([]int, len(a))
	for i := range index {
		index[i] = i
	}

	fmt.Println("Matrix A")
	//a.printMatrix(index)
	fmt.Println("Vector B")
	//b.printVector()

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
			pivot = a[i][index[i]];
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

	fmt.Println("Vector X")
	for i := 0; i < len(x); i++ {
		fmt.Printf("[%v] ", x[index[i]])
	}
	fmt.Println()


	checkResult(t, x)

	m2 := make([][]rational.Rational, len(m))
	for i, iv := range m {
		mr := make([]rational.Rational, len(m[i]))
		for j, jv := range iv {
			mr[j], _ = rational.NewFromFloat(jv)
		}
		m2[i] = mr
	}

	res, _ := gaussian.SolveGaussian(m2, false)

	fmt.Println()
	for _, elem := range res {
		fmt.Print(elem[0].Float64())
		fmt.Println()
	}
}

func (m models.Matrix) printMatrix(index []int) {
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

func (v models.Vector) printVector() {
	for i := 0; i < len(v); i++ {
		fmt.Printf("[%v] ", v[i])
	}

	fmt.Println()
	fmt.Println()
}

func checkResult(m models.Matrix, v models.Vector) bool {

	var max float64
	if len(v) > 0 {
		for i, _ := range m {
			row := m[i]
			correct_answ := row[len(row)-1]
			row = row[:len(row)-1]
			if max < math.Abs(math.Abs(v.Scalar(row)) - correct_answ){
				max = math.Abs(math.Abs(v.Scalar(row)) - correct_answ)
			}
		}
	}

	fmt.Println("max: ",max)

	return true
}
