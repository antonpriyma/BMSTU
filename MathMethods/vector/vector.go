package vector

import (
	"fmt"
	"math"
	"strconv"
)

type Vector struct {
	Size  int
	Value []float64
}

func NewVector(value []float64) *Vector {
	return &Vector{
		Size:  len(value),
		Value: value,
	}
}

func EmptyVector(size int) *Vector {
	value := make([]float64, size)

	for i := 0; i < size; i++ {
		value[i] = 0
	}

	return NewVector(value)
}

func Of(v *Vector) *Vector {
	value := make([]float64, v.Size)

	for i := 0; i < v.Size; i++ {
		value[i] = v.Value[i]
	}

	return NewVector(value)
}

func (v *Vector) ScalarProd(vector *Vector) (float64, error) {
	if !equalSize(v, vector) {
		return 0, fmt.Errorf(ErrorDimensions)
	}

	res := float64(0)

	for i := range v.Value {
		res += v.Value[i] * vector.Value[i]
	}

	return res, nil
}

func (v *Vector) Sum(vector *Vector) (*Vector, error) {
	if !equalSize(v, vector) {
		return nil, fmt.Errorf(ErrorDimensions)
	}

	value := make([]float64, v.Size)

	for i := range value {
		value[i] = v.Value[i] + vector.Value[i]
	}

	return NewVector(value), nil
}

func (v *Vector) MulScalar(x float64) *Vector {
	value := make([]float64, v.Size)

	for i := range value {
		value[i] = v.Value[i] * x
	}

	return NewVector(value)
}

func (v *Vector) UniformNorm() float64 {
	norm := math.Abs(v.Value[0])

	for i := 1; i < v.Size; i++ {
		x := math.Abs(v.Value[i])

		if x > norm {
			norm = x
		}
	}

	return norm
}

func (v *Vector) Sub(vector *Vector) (*Vector, error) {
	if !equalSize(v, vector) {
		return nil, fmt.Errorf(ErrorDimensions)
	}

	value := make([]float64, len(v.Value))

	for i := range v.Value {
		value[i] = v.Value[i] - vector.Value[i]
	}

	return NewVector(value), nil
}

func (v *Vector) Norm() float64 {
	sc, _ := v.ScalarProd(v)

	return math.Sqrt(sc)
}

func (v *Vector) SwapComponents(i, j int) error {
	if i < 0 || j < 0 || i >= v.Size || j >= v.Size {
		return fmt.Errorf(ErrorIndexOutOfRange)
	}
	v.Value[i], v.Value[j] = v.Value[j], v.Value[i]
	return nil
}

func (v *Vector) Copy() *Vector {
	value := make([]float64, v.Size)
	copy(value, v.Value)

	return NewVector(value)
}

func (v *Vector) String() string {
	s := "("

	for i, c := range v.Value {
		s += strconv.FormatFloat(c, 'f', -1, 64)

		if i != len(v.Value)-1 {
			s += ", "
		}
	}

	return s + ")"
}

func equalSize(v1, v2 *Vector) bool {
	return v1.Size == v2.Size
}

var (
	ErrorDimensions      = "the dimensions of the vectors must match"
	ErrorIndexOutOfRange = "index out of vector range"
)
