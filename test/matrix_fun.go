package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

func main() {
	seed := 10

	// 1. Generate a n×n matrix of random values.
	data := make([]float64, 9)
	for i := range data {
		data[i] = float64(rand.Intn(seed))
	}
	A := mat.NewDense(3, 3, data)
	fmt.Printf("A:\n")
	matPrint(A)

	// 2. Calculate Inverse of matrix
	var A_Inv mat.Dense
	err := A_Inv.Inverse(A)
	if err != nil {
		fmt.Printf("A is not invertible: %v", err)
	}
	fmt.Printf("A Inverse:\n")
	matPrint(&A_Inv)

	// 3. Verify Inverse of matrix (A * A^-1 = I)
	var I mat.Dense
	I.Mul(A, &A_Inv)
	fmt.Printf("I matrix:\n")
	matPrint(&I)
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
