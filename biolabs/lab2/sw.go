package main

const (
	zero = iota
	left
	up
	diag
)

type item struct {
	value     int // набранный балл в матрице
	direction int // направление движения по матрице для восстановления выравнивания
}

type SWData struct {
	Scores func(uint8, uint8) int // скоринг-функция
	Gap    int                    // штраф за пропуск
}

func (sw *SWData) Align(a, b string) (string, string, int) {
	matrix, imax, jmax := sw.fill(a, b)
	return sw.best(a, b, matrix, imax, jmax)
}

func (sw *SWData) fill(a, b string) ([][]item, int, int) {
	// создание и заполнение первого ряда и первой строки матрицы
	res := make([][]item, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		res[i] = make([]item, len(b)+1)
	}

	// заполнение оставшейся части матрицы
	mx := 0
	imax := 0
	jmax := 0
	for i := 1; i < len(a)+1; i++ {
		for j := 1; j < len(b)+1; j++ {
			cur := max(res[i-1][j-1].value+sw.Scores(a[i], b[j]), res[i-1][j].value+sw.Gap, res[i][j-1].value+sw.Gap, 0)
			res[i][j] = item{value: cur}
			switch cur {
			case res[i-1][j-1].value + sw.Scores(a[i-1], b[j-1]):
				res[i][j].direction = diag
			case res[i-1][j].value + sw.Gap:
				res[i][j].direction = up
			case res[i][j-1].value + sw.Gap:
				res[i][j].direction = left
			default:
				res[i][j].direction = zero
			}
			if res[i][j].value >= mx {
				mx = res[i][j].value
				imax = i
				jmax = j
			}
		}
	}

	return res, imax, jmax
}

func (sw *SWData) best(a, b string, matrix [][]item, imax, jmax int) (string, string, int) {
	// восстановление выравнивания
	alignA := ""
	alignB := ""
	i := imax
	j := jmax
	for matrix[i][j].direction != zero {
		switch matrix[i][j].direction {
		case left: // здесь был штраф за пропуск в первой строке
			alignA = "-" + alignA
			alignB = toStr(b[j-1]) + alignB
			j--
		case up: // здесь был штраф за пропуск во второй строке
			alignA = toStr(a[i-1]) + alignA
			alignB = "-" + alignB
			i--
		default: // без пропусков
			alignA = toStr(a[i-1]) + alignA
			alignB = toStr(b[j-1]) + alignB
			i--
			j--
		}
	}
	return alignA, alignB, matrix[imax][jmax].value
}
