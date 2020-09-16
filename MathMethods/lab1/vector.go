package main

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
