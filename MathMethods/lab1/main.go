package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	a := Matrix{[]float64{3, -9, 3}, []float64{2, -4, 4}, []float64{1, 8, -18}}
	b := Vector{-18, -10, 35}

	// индекс, определяет порядок колонок в матрице
	index := make([]int, len(a))
	for i := range index {
		index[i] = i
	}

	// отображаем исходные данные
	fmt.Println("Matrix A")
	a.dump(index)
	fmt.Println("Vector B")
	b.dump()

	// прямой ход
	for i := 0; i < len(a); i++ {

		// главный элемент, значение по умолчанию
		r := a[i][index[i]]

		// если главный элемент равен нулю, нужно найти другой
		// методом перестановки колонок в матрице
		if r == 0 {
			var kk int

			// двигаемся вправо от диаганаотного элемента, для поиска максимального по модулю элемента
			for k := i; k < len(a); k++ {
				if math.Abs(a[i][index[k]]) > r {
					kk = k
				}
			}

			// если удалось найти главный элемент
			if kk > 0 {
				// меняем местами колонки, так чтобы главный элемент встал в диагональ матрицы
				index[i], index[kk] = index[kk], index[i]
			}

			// получаем главный элемента, текущей строки из диагонали
			r = a[i][index[i]]
		}

		// если главный элемент строки равен 0, метод гаусса не работает
		if r == 0 {
			if b[i] == 0 {
				fmt.Println("система имеет множество решений")
			} else {
				fmt.Println("система не имеет решений")
			}
			os.Exit(1)
		}

		// деление элементов текущей строки, на главный элемент
		for j := 0; j < len(a[i]); j++ {
			a[i][index[j]] /= r
		}
		b[i] /= r

		// вычитание текущей строки из всех ниже расположенных строк с занулением I - ого элемента в каждой из них
		for k := i + 1; k < len(a); k++ {
			r = a[k][index[i]]
			for j := 0; j < len(a[i]); j++ {
				a[k][index[j]] = a[k][index[j]] - a[i][index[j]]*r
			}
			b[k] = b[k] - b[i]*r
		}

		// отображаем дамп матрицы A и вектора B
		fmt.Println("++++++++++++\n")
		fmt.Println("Matrix A")
		a.dump(index)
		fmt.Println("Vector B")
		b.dump()
	}

	var x Vector = make(Vector, len(b))

	// обратный ход
	for i := len(a) - 1; i >= 0; i-- {
		// Задается начальное значение элемента x[I].
		x[i] = b[i]

		// Корректируется искомое значение x[I].
		// В цикле по J от I+1 до N (в случае, когда I=N, этот шаг не выполняется) производятся вычисления x[I]:=  x[I] - x[J]* A[I, J].
		for j := i + 1; j < len(a); j++ {
			x[i] = x[i] - (x[j] * a[i][index[j]])
		}
	}

	fmt.Println("++++++++++++\n")
	fmt.Println("Vector X")
	for i := 0; i < len(x); i++ {
		fmt.Printf("[%v] ", x[index[i]])
	}
	fmt.Println()
}

func (m Matrix) dump(index []int) {
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

func (v Vector) dump() {

	for i := 0; i < len(v); i++ {
		fmt.Printf("[%v] ", v[i])
	}

	fmt.Println()
	fmt.Println()
}
