package main

import (
	"strings"
	"unicode/utf8"
)

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

type HBData struct {
	Scores func(uint8, uint8) int // скоринг-функция
	Gap    int                    // штраф за пропуск
}

func (hb *HBData) Align(a, b string) (string, string, int) {
	switch {
	case a == "":
		count := utf8.RuneCountInString(b)
		return strings.Repeat("-", count), b, hb.Gap * count

	case b == "":
		count := utf8.RuneCountInString(a)
		return a, strings.Repeat("-", count), hb.Gap * count

	case len(a) == 1 || len(b) == 1:
		resa, resb, score := hb.nw(a, b)
		return resa, resb, score

	default:
		mid := len(a) / 2

		// Делим первую строку пополам
		leftScore := hb.nwScore(a[:mid], b)
		rightScore := hb.nwScore(revStr(a[mid:]), revStr(b))

		// Ищем оптимальное разделение второй строки
		index, _ := argMax(leftScore, revInt(rightScore))
		leftResA, leftResB, sc1 := hb.Align(a[:mid], b[:index])
		rightResA, rightResB, sc2 := hb.Align(a[mid:], b[index:])

		return leftResA + rightResA, leftResB + rightResB, sc1 + sc2
	}
}

// Последняя строка матрицы из алгоритма Нидльмана-Вунша
func (hb *HBData) nwScore(a, b string) []int {
	// Для нахождения последней строки достаточно хранить в памяти лишь две
	score := make([][]int, 2)
	score[0] = make([]int, len(b)+1)
	score[1] = make([]int, len(b)+1)

	for i := range score[0] {
		score[0][i] = i * hb.Gap
	}

	for i := range a {
		score[1][0] = score[0][0] + hb.Gap
		for j := 1; j < len(b)+1; j++ {
			sub := score[0][j-1] + hb.Scores(a[i], b[j-1])
			del := score[0][j] + hb.Gap
			ins := score[1][j-1] + hb.Gap
			score[1][j] = max(sub, del, ins)
		}
		// Удаляем предыдущую строку из памяти
		score[0] = score[1]
		score[1] = make([]int, len(b)+1)
	}

	return score[0]
}

func (hb *HBData) nw(a, b string) (string, string, int) {
	return hb.best(a, b, hb.fill(a, b))
}

func (hb *HBData) fill(a, b string) [][]item {
	// создание и заполнение первого ряда и первой строки матрицы
	res := make([][]item, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		res[i] = make([]item, len(b)+1)
		res[i][0] = item{
			value:     hb.Gap * i,
			direction: up,
		}
	}

	for i := 0; i < len(b)+1; i++ {
		res[0][i] = item{
			value:     hb.Gap * i,
			direction: left,
		}
	}

	res[0][0] = item{
		value:     0,
		direction: zero,
	}

	// заполнение оставшейся части матрицы
	for i := 1; i < len(a)+1; i++ {
		for j := 1; j < len(b)+1; j++ {
			cur := max(res[i-1][j-1].value+hb.Scores(a[i-1], b[j-1]), res[i-1][j].value+hb.Gap, res[i][j-1].value+hb.Gap)
			res[i][j] = item{value: cur}
			switch cur {
			case res[i-1][j-1].value + hb.Scores(a[i-1], b[j-1]):
				res[i][j].direction = diag
			case res[i-1][j].value + hb.Gap:
				res[i][j].direction = up
			default:
				res[i][j].direction = left
			}

		}
	}

	return res
}

func (hb *HBData) best(a, b string, matrix [][]item) (string, string, int) {
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

func argMax(a, b []int) (int, int) {
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

	return ind, val
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
