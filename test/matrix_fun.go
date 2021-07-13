package main

import (
	"encoding/base64"
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
	fmt.Printf("%f\n", ATrace)

	// 8. Generate Permutation 'pi' from Original Order
	Permutation := make([]int, matSize)
	for i := 0; i < matSize; i++ {
		Permutation[i] = i
	}
	rand.Shuffle(matSize, func(i, j int) { Permutation[i], Permutation[j] = Permutation[j], Permutation[i] })
	fmt.Printf("Permutation: %v\n", Permutation)

	// 9. Generate Permutation Matrix from Permutation 'pi'
	var PermutationMatrix mat.Dense
	PermutationMatrix.Permutation(matSize, Permutation)
	fmt.Printf("Permuation Matrix:\n")
	matPrint(&PermutationMatrix)

	// 10. Generate Permuted Matrix from Permutation Matrix
	var Permuted mat.Dense
	Permuted.Mul(A, &PermutationMatrix)
	fmt.Printf("Permuted Matrix:\n")
	matPrint(&Permuted)

	// 11. Convert Matrix as Byte[] and Base64-String
	MatBytes, err := Permuted.MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
	}
	MatString := base64.StdEncoding.EncodeToString(MatBytes)
	fmt.Printf("Base64 String: %s\n", MatString)

	// 12. Convert Base64-String String to Matrix
	RecvBytes, _ := base64.StdEncoding.DecodeString(MatString)
	var RecvMatrix mat.Dense
	RecvMatrix.UnmarshalBinary(RecvBytes)
	fmt.Printf("Recovered Matrix:\n")
	matPrint(&RecvMatrix)
}

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
