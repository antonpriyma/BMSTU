package models

import (
	"fmt"
	"math"
	"math/rand"
)

//type Vector struct {
//	X, Y, Z float64
//}
//
//func (a Vector) Add(b Vector) Vector {
//	return Vector{
//		X: a.X + b.X,
//		Y: a.Y + b.Y,
//		Z: a.Z + b.Z,
//	}
//}
//
//func (a Vector) Sub(b Vector) Vector {
//	return Vector{
//		X: a.X - b.X,
//		Y: a.Y - b.Y,
//		Z: a.Z - b.Z,
//	}
//}
//
//func (a Vector) MultiplyByScalar(s float64) Vector {
//	return Vector{
//		X: a.X * s,
//		Y: a.Y * s,
//		Z: a.Z * s,
//	}
//}
//
//func (a Vector) Dot(b Vector) float64 {
//	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
//}
//
//func (a Vector) Cross(b Vector) Vector {
//	return Vector{
//		X: a.Y*b.Z - a.Z*b.Y,
//		Y: a.Z*b.X - a.X*b.Z,
//		Z: a.X*b.Y - a.Y*b.X,
//	}
//}

type Vector []float64

func (v Vector) Scalar(a Vector) (res float64) {
	if len(v) != len(a) {
		return 0
	}

	for i, _ := range v {
		res += v[i] * a[i]
	}

	return res
}


func (v Vector) Sum(v1 Vector) Vector {
	if len(v) != len(v1) {
		fmt.Println("Wrong vector size")
		return nil
	}
	res := make([]float64, len(v))
	for i, elem := range v {
		res[i] = elem + v1[i]
	}
	return res

}


func (v Vector) Norm() float64 {
	var res float64

	for _, elem := range v {
		res += elem * elem
	}
	return math.Sqrt(res)
}

func (v Vector) NormNew() float64 {
	return v.Max()
}

func (v Vector) Max() float64 {
	var max float64
	max = math.Abs(v[0])
	for _, x := range v{
		if math.Abs(x) > max {
			max = math.Abs(x)
		}
	}

	return max
}


func(v Vector) Sub(v1 Vector) Vector {
	if len(v) != len(v1) {
		fmt.Println("Wrong vector size")
		return nil
	}
	res := make([]float64, len(v))
	for i, elem := range v {
		res[i] = elem - v1[i]
	}
	return res
}

func (v Vector) SumElems() float64 {
	var sum float64
	sum = 0
	for _, x := range  v {
		sum += x
	}

	return sum
}

func RandomVector(n int) Vector {
	v := make(Vector, n)

	for i := 0; i < n; i++ {
		v[i] = rand.Float64()
	}

	return v
}

func (v Vector) printVector() {
	for i := 0; i < len(v); i++ {
		fmt.Printf("[%v] ", v[i])
	}

	fmt.Println()
	fmt.Println()
}