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

type AGData struct {
	Scores func(uint8, uint8) int // скоринг-функция
	Gap    int                    // штраф за пропуск
	GapAdd int                    // штраф за серию
}

func (ag *AGData) Align(a, b string) (string, string, int) {
	return ag.best(a, b, ag.fill(a, b))
}

// Последняя строка матрицы из алгоритма Нидльмана-Вунша
func (ag *AGData) nwScore(a, b string) []int {
	// Для нахождения последней строки достаточно хранить в памяти лишь две
	score := make([][]int, 2)
	score[0] = make([]int, len(b)+1)
	score[1] = make([]int, len(b)+1)

	for i := range score[0] {
		score[0][i] = i * ag.Gap
	}

	for i := range a {
		score[1][0] = score[0][0] + ag.Gap
		for j := 1; j < len(b)+1; j++ {
			sub := score[0][j-1] + ag.Scores(a[i], b[j-1])
			del := score[0][j] + ag.Gap
			ins := score[1][j-1] + ag.Gap
			score[1][j] = max(sub, del, ins)
		}
		// Удаляем предыдущую строку из памяти
		score[0] = score[1]
		score[1] = make([]int, len(b)+1)
	}

	return score[0]
}

func (ag *AGData) fill(a, b string) [][]item {
	G := make([][]item, len(a)+1)
	V := make([][]int, len(a)+1)
	H := make([][]*int, len(a)+1)

	for i := range G {
		G[i] = make([]item, len(b)+1)
		V[i] = make([]int, len(b)+1)
		H[i] = make([]*int, len(b)+1)

		gap := ag.Gap + ag.GapAdd*(i-1)
		G[i][0] = item{
			direction: up,
			value:     gap,
		}
		V[i][0] = gap
	}

	for i := range G[0] {
		gap := ag.Gap + ag.GapAdd*(i-1)
		G[0][i] = item{
			direction: left,
			value:     gap,
		}
		V[0][i] = gap
	}

	G[0][0] = item{
		value:     0,
		direction: zero,
	}
	V[0][0] = 0

	for i := 1; i < len(G); i++ {
		for j := 1; j < len(G[0]); j++ {
			g := max(V[i-1][j]+ag.Gap, G[i-1][j].value+ag.GapAdd)
			G[i][j].value = g

			h := V[i][j-1] + ag.Gap
			if H[i][j-1] != nil {
				h = max(V[i][j-1]+ag.Gap, *H[i][j-1]+ag.GapAdd)
			}
			H[i][j] = &h

			diagVal := V[i-1][j-1] + ag.Scores(a[i-1], b[j-1])
			maxVal := max(diagVal, g, h)
			switch maxVal {
			case diagVal:
				G[i][j].direction = diag
			case g:
				G[i][j].direction = up
			case h:
				G[i][j].direction = left
			}
			V[i][j] = maxVal
		}
	}

	return G
}

func (ag *AGData) best(a, b string, matrix [][]item) (string, string, int) {
	// восстановление выравнивания
	alignA := ""
	alignB := ""
	i := len(a)
	j := len(b)
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
	return alignA, alignB, matrix[len(a)][len(b)].value
}

func argMax(a, b []int) int {
	if len(a) != len(b) {
		panic("len(a) must be equal len(b)")
	}

	if len(a) == 0 {
		panic("len(a) must be non-zero")
	}

	ind, val := -1, 0
	for i := range a {
		if a[i]+b[i] > val || ind < 0 {
			ind, val = i, a[i]+b[i]
		}
	}

	return ind
}

func revStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func revInt(a []int) []int {
	res := make([]int, len(a))
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = a[j], a[i]
	}
	return res
}
