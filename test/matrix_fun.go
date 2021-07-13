package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

func main() {
	seed := 10
	matSize := 3

	// 1. Generate a n×n matrix of random values.
	data := make([]float64, matSize*matSize)
	for i := range data {
		data[i] = float64(rand.Intn(seed))
	}
	A := mat.NewDense(matSize, matSize, data)
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

	// 4. Generate Diagonal Matrix from a vector of size n
	DiagMat := mat.NewDiagDense(len(data), data)
	fmt.Printf("Diag Matrix:\n")
	matPrint(DiagMat)

	// 5. Generate a Random Lower-Triangular Matrix
	LowTriMat := mat.NewTriDense(matSize, false, data)
	fmt.Printf("Lower-Triangular Matrix:\n")
	matPrint(LowTriMat)

	// 6. Set Diagonal Entries of a matrix as x=1
	var Diag1Matrix mat.Dense
	Diag1Matrix.CloneFrom(LowTriMat)
	for i := 0; i < matSize; i++ {
		Diag1Matrix.Set(i, i, 1)
	}
	fmt.Printf("Lower-Triangular Matrix with Diagonal Elements as 1:\n")
	matPrint(&Diag1Matrix)

	// 7. Generate Trace of a matrix
	ATrace := A.Trace()
	fmt.Printf("Trace of Matrix (A):\n")
	fmt.Printf("%f", ATrace)
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
