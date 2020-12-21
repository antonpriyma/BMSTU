package main

import (
	"fmt"
	"github.com/AntonPriyma/MathMethods/vector"
	"math"
	"strconv"
	"sync"
)

type Matrix struct {
	RowCount    int
	ColumnCount int
	Value       [][]float64
}

func NewMatrix(value [][]float64) (*Matrix, error) {
	if !equalColumnsCount(value) {
		return nil, fmt.Errorf(ErrorColumnsNumber)
	}

	return &Matrix{
		RowCount:    len(value),
		ColumnCount: len(value[0]),
		Value:       value,
	}, nil
}

func IdentityMatrix(rowCount, columnCount int) *Matrix {
	value := make([][]float64, rowCount)

	for i := 0; i < rowCount; i++ {
		value[i] = make([]float64, columnCount)
		for j := 0; j < columnCount; j++ {
			value[i][j] = 0
		}
		value[i][i] = 1
	}

	matrix, _ := NewMatrix(value)
	return matrix
}

func (m *Matrix) MulMatrix(secondMatrix *Matrix) (*Matrix, error) {
	if m.ColumnCount != secondMatrix.RowCount {
		return nil, fmt.Errorf(ErrorMulMatricesSize)
	}

	res := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		res[i] = make([]float64, secondMatrix.ColumnCount)

		for k := 0; k < secondMatrix.ColumnCount; k++ {
			elem := float64(0)

			for j := 0; j < m.ColumnCount; j++ {
				elem += m.Value[i][j] * secondMatrix.Value[j][k]
			}

			res[i][k] = elem
		}
	}

	return NewMatrix(res)
}

func (m *Matrix) MulVector(v *vector.Vector) (*vector.Vector, error) {
	if m.ColumnCount != v.Size {
		return nil, fmt.Errorf(ErrorMulVectorDimension)
	}

	res := make([]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		elem := float64(0)

		for j := 0; j < m.ColumnCount; j++ {
			elem += m.Value[i][j] * v.Value[j]
		}

		res[i] = elem
	}

	return vector.NewVector(res), nil
}

func (m *Matrix) MulScalar(s float64) *Matrix {
	value := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		value[i] = make([]float64, m.ColumnCount)

		for j := 0; j < m.ColumnCount; j++ {
			value[i][j] = m.Value[i][j] * s
		}
	}

	newM, _ := NewMatrix(value)
	return newM
}

func (m *Matrix) Sum(matrix *Matrix) *Matrix {
	if m.RowCount != matrix.RowCount || m.ColumnCount != matrix.ColumnCount {
		return nil
	}

	res := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		res[i] = make([]float64, m.ColumnCount)

		for j := 0; j < m.ColumnCount; j++ {
			res[i][j] = m.Value[i][j] + matrix.Value[i][j]
		}
	}

	newMatrix, _ := NewMatrix(res)
	return newMatrix
}

func (m *Matrix) Sub(matrix *Matrix) *Matrix {
	if m.RowCount != matrix.RowCount || m.ColumnCount != matrix.ColumnCount {
		return nil
	}

	res := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		res[i] = make([]float64, m.ColumnCount)

		for j := 0; j < m.ColumnCount; j++ {
			res[i][j] = m.Value[i][j] - matrix.Value[i][j]
		}
	}

	newMatrix, _ := NewMatrix(res)
	return newMatrix
}

func (m *Matrix) Inverse2() (*Matrix, error) {
	if m.ColumnCount != 2 || m.RowCount != 2 {
		return nil, fmt.Errorf(ErrorMatrix2by2)
	}

	a, b, c, d := m.Value[0][0], m.Value[0][1], m.Value[1][0], m.Value[1][1]
	factor := a*d - b*c

	value := [][]float64{
		{d / factor, -b / factor},
		{-c / factor, a / factor},
	}

	return NewMatrix(value)
}

func (m *Matrix) ConditionNumber2() (float64, error) {
	if m.ColumnCount != 2 || m.RowCount != 2 {
		return 0, fmt.Errorf(ErrorMatrix2by2)
	}
	i, _ := m.Inverse2()
	return i.UniformNorm() * m.UniformNorm(), nil
}

func (m *Matrix) UniformNorm() float64 {
	norm := math.Abs(m.Value[0][0])

	for i := 0; i < m.RowCount; i++ {
		sum := float64(0)

		for j := 0; j < m.ColumnCount; j++ {
			sum += math.Abs(m.Value[i][j])
		}

		if sum > norm {
			norm = sum
		}
	}

	return norm
}

func (m *Matrix) Norm() float64 {
	res := float64(0)

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			res += m.Value[i][j] * m.Value[i][j]
		}
	}

	return math.Sqrt(res)
}

func (m *Matrix) IsTriangle() bool {
	for i := 1; i < m.RowCount; i++ {
		for j := 0; j < i; j++ {
			if m.Value[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) IsDominance() bool {
	for i := 0; i < m.RowCount; i++ {
		diag := math.Abs(m.Value[i][i])
		sum := float64(0)

		for j := 0; j < m.ColumnCount; j++ {
			if i != j {
				sum += math.Abs(m.Value[i][j])
			}
		}

		if diag <= sum {
			return false
		}
	}

	return true
}

func (m *Matrix) IsSymmetric() (bool, error) {
	if m.RowCount != m.ColumnCount {
		return false, fmt.Errorf(ErrorSquareMatrix)
	}

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if m.Value[i][j] != m.Value[j][i] {
				return false, nil
			}
		}
	}

	return true, nil
}

func (m *Matrix) IsPositive() bool {
	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			if m.Value[i][j] <= 0 {
				return false
			}
		}
	}

	return true
}

func (m *Matrix) SwapColumns(i, j int) error {
	if i < 0 || j < 0 || i >= m.ColumnCount || j >= m.ColumnCount {
		return fmt.Errorf(ErrorSwapIndexOutOfRange)
	}

	for k := 0; k < m.RowCount; k++ {
		m.Value[k][i], m.Value[k][j] = m.Value[k][j], m.Value[k][i]
	}

	return nil
}

func (m *Matrix) SwapRows(i, j int) error {
	if i < 0 || j < 0 || i >= m.RowCount || j >= m.RowCount {
		return fmt.Errorf(ErrorSwapIndexOutOfRange)
	}

	for k := 0; k < m.ColumnCount; k++ {
		m.Value[i][k], m.Value[j][k] = m.Value[j][k], m.Value[i][k]
	}

	return nil
}

func (m *Matrix) Copy() *Matrix {
	rows := make([][]float64, m.RowCount)

	for i := 0; i < m.RowCount; i++ {
		rows[i] = make([]float64, m.ColumnCount)
		copy(rows[i], m.Value[i])
	}

	newMatrix, _ := NewMatrix(rows)

	return newMatrix
}

func (m *Matrix) Flat() []float64 {
	res := make([]float64, m.RowCount*m.ColumnCount)

	for i := 0; i < m.RowCount; i++ {
		for j := 0; j < m.ColumnCount; j++ {
			res[i*m.ColumnCount+j] = m.Value[i][j]
		}
	}

	return res
}

func (m *Matrix) String() string {
	s := ""

	for i := 0; i < m.RowCount; i++ {
		row := "|"

		for j := 0; j < m.ColumnCount; j++ {
			row += strconv.FormatFloat(m.Value[i][j], 'f', -1, 64)

			if j != m.ColumnCount-1 {
				row += "\t"
			}
		}

		row += "|\n"
		s += row
	}

	return s
}


type SliceRange struct {
	Start int
	End   int
}

func (m *Matrix) Slice(rowRange, colRange SliceRange) *Matrix {
	rows := m.Value[rowRange.Start:rowRange.End]
	s := make([][]float64, len(rows))

	for i, r := range rows {
		s[i] = r[colRange.Start:colRange.End]
	}

	newMatrix, _ := NewMatrix(s)
	return newMatrix
}

func NearestPower2(x int) int {
	x--
	x |= x >> 1
	x |= x >> 2
	x |= x >> 4
	x |= x >> 8
	x |= x >> 16
	x |= x >> 32
	return x + 1
}

func (m *Matrix) CompleteToSquare(size int) *Matrix {
	value := make([][]float64, size)

	for i := 0; i < size; i++ {
		value[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			if i < m.RowCount && j < m.ColumnCount {
				value[i][j] = m.Value[i][j]
			} else {
				value[i][j] = 0
			}
		}
	}

	newMatrix, _ := NewMatrix(value)
	return newMatrix
}

func CopyToSlice(slice *Matrix, src *Matrix) {
	for i := 0; i < slice.RowCount; i++ {
		for j := 0; j < slice.ColumnCount; j++ {
			slice.Value[i][j] = src.Value[i][j]
		}
	}
}

func equalColumnsCount(rows [][]float64) bool {
	for i := 0; i < len(rows)-1; i++ {
		if len(rows[i]) != len(rows[i+1]) {
			return false
		}
	}

	return true
}

const (
	ErrorColumnsNumber       = "the number of columns in rows must be equal"
	ErrorMulMatricesSize     = "the number of columns of the first matrix must match the number of rows of the second"
	ErrorMulVectorDimension  = "incorrect vector dimension"
	ErrorSwapIndexOutOfRange = "swap index out of matrix range"
	ErrorMatrix2by2          = "the method is applicable only to 2-by-2 matrices"
	ErrorSquareMatrix        = "matrix must be square"
)

const nMin = 128

func (m *Matrix) MulStrass(secondMatrix *Matrix, concurrency bool) (*Matrix, error) {
	if m.ColumnCount != secondMatrix.RowCount {
		return nil, fmt.Errorf(ErrorMulMatricesSize)
	}

	n := m.RowCount
	if secondMatrix.RowCount > n {
		n = secondMatrix.RowCount
	}
	if secondMatrix.ColumnCount > n {
		n = secondMatrix.ColumnCount
	}
	n = NearestPower2(n)

	mSquare1 := m.CompleteToSquare(n)
	mSquare2 := secondMatrix.CompleteToSquare(n)
	mul := mulStrass(mSquare1, mSquare2, concurrency)

	return mul.Slice(
		SliceRange{Start: 0, End: m.RowCount},
		SliceRange{Start: 0, End: secondMatrix.ColumnCount},
	), nil
}

func mulStrass(a *Matrix, b *Matrix, concurrency bool) *Matrix {
	n := a.RowCount

	if n <= nMin {
		mul, _ := a.MulMatrix(b)
		return mul
	}

	middle := n / 2
	u := SliceRange{Start: 0, End: middle}
	v := SliceRange{Start: middle, End: n}
	c := IdentityMatrix(a.RowCount, a.ColumnCount)

	var p1, p2, p3, p4, p5, p6, p7 *Matrix
	wg := sync.WaitGroup{}

	if concurrency {
		wg.Add(7)
		go func() {
			p1 = mulStrass(a.Slice(u, u).Sum(a.Slice(v, v)), b.Slice(u, u).Sum(b.Slice(v, v)), concurrency)
			wg.Done()
		}()
		go func() {
			p2 = mulStrass(a.Slice(v, u).Sum(a.Slice(v, v)), b.Slice(u, u), concurrency)
			wg.Done()
		}()
		go func() {
			p3 = mulStrass(a.Slice(u, u), b.Slice(u, v).Sub(b.Slice(v, v)), concurrency)
			wg.Done()
		}()
		go func() {
			p4 = mulStrass(a.Slice(v, v), b.Slice(v, u).Sub(b.Slice(u, u)), concurrency)
			wg.Done()
		}()
		go func() {
			p5 = mulStrass(a.Slice(u, u).Sum(a.Slice(u, v)), b.Slice(v, v), concurrency)
			wg.Done()
		}()
		go func() {
			p6 = mulStrass(a.Slice(v, u).Sub(a.Slice(u, u)), b.Slice(u, u).Sum(b.Slice(u, v)), concurrency)
			wg.Done()
		}()
		go func() {
			p7 = mulStrass(a.Slice(u, v).Sub(a.Slice(v, v)), b.Slice(v, u).Sum(b.Slice(v, v)), concurrency)
			wg.Done()
		}()
	} else {
		p1 = mulStrass(a.Slice(u, u).Sum(a.Slice(v, v)), b.Slice(u, u).Sum(b.Slice(v, v)), concurrency)
		p2 = mulStrass(a.Slice(v, u).Sum(a.Slice(v, v)), b.Slice(u, u), concurrency)
		p3 = mulStrass(a.Slice(u, u), b.Slice(u, v).Sub(b.Slice(v, v)), concurrency)
		p4 = mulStrass(a.Slice(v, v), b.Slice(v, u).Sub(b.Slice(u, u)), concurrency)
		p5 = mulStrass(a.Slice(u, u).Sum(a.Slice(u, v)), b.Slice(v, v), concurrency)
		p6 = mulStrass(a.Slice(v, u).Sub(a.Slice(u, u)), b.Slice(u, u).Sum(b.Slice(u, v)), concurrency)
		p7 = mulStrass(a.Slice(u, v).Sub(a.Slice(v, v)), b.Slice(v, u).Sum(b.Slice(v, v)), concurrency)
	}

	wg.Wait()

	CopyToSlice(c.Slice(u, u), p1.Sum(p4).Sub(p5).Sum(p7))
	CopyToSlice(c.Slice(u, v), p3.Sum(p5))
	CopyToSlice(c.Slice(v, u), p2.Sum(p4))
	CopyToSlice(c.Slice(v, v), p1.Sum(p3).Sub(p2).Sum(p6))

	return c
}
