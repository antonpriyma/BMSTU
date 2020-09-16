package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)


func TestVectorScalar(t *testing.T) {
	res := Vector{1,2,3}.Scalar(Vector{0,1,2})
	require.Equal(t,float64(8), res)

	res = Vector{1,1,1}.Scalar(Vector{1,1,2})
	require.Equal(t,float64(4), res)

	res = Vector{1,1,1}.Scalar(Vector{1,1,2,4})
	require.Equal(t,float64(0), res)
}