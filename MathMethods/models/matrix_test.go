package models

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMatrixVectorMult(t *testing.T) {
	matrix := Matrix{
		[]float64{1, 2},
		[]float64{3, 4},
	}

	vector := Vector{1, 2}

	res := matrix.VectorMult(vector)

	require.Equal(t, Vector{7, 10}, res)
}

func TestMatrixMatrixMult(t *testing.T) {
	matrixA := Matrix{
		[]float64{1, 2},
		[]float64{3, 4},
	}

	matrixB := Matrix{
		[]float64{1, 2},
		[]float64{3, 4},
	}

	require.Equal(t, Matrix{[]float64{7, 10}, []float64{15, 22}}, matrixA.MatrixMult(matrixB))
}
