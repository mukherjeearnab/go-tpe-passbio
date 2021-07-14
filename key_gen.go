package tpe

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

type SecretKey struct {
	key Key
}

type Key struct {
	M_1  mat.Matrix `json:"m1"`
	M_2  mat.Matrix `json:"m2"`
	M_1i mat.Matrix `json:"m1i"`
	M_2i mat.Matrix `json:"m2i"`
	Pi   []int      `json:"pi"`
}

type exportKey struct {
	M_1  []byte `json:"m1"`
	M_2  []byte `json:"m2"`
	M_1i []byte `json:"m1i"`
	M_2i []byte `json:"m2i"`
	Pi   []int  `json:"pi"`
}

// Generate New Secret Key
func (sk *SecretKey) KeyGen(seed int64, matrixSize int) {
	rand.Seed(seed)
	sk.key.M_1, sk.key.M_1i = generateMatrix(matrixSize)
	sk.key.M_2, sk.key.M_2i = generateMatrix(matrixSize)
	sk.key.Pi = generatePermutation(matrixSize)
}

// Export Secret Key as JSON string
func (sk *SecretKey) ExportKey() string {
	M_1Bytes, err := mat.DenseCopyOf(sk.key.M_1).MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return ""
	}
	M_1iBytes, err := mat.DenseCopyOf(sk.key.M_1i).MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return ""
	}
	M_2Bytes, err := mat.DenseCopyOf(sk.key.M_2).MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return ""
	}
	M_2iBytes, err := mat.DenseCopyOf(sk.key.M_2i).MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return ""
	}

	expKey := exportKey{M_1Bytes, M_1iBytes, M_2Bytes, M_2iBytes, sk.key.Pi}
	keyJSON, _ := json.Marshal(expKey)
	return string(keyJSON)
}

/* ***************************
   Auxiliary Functions
   *************************** */

// Generate A Random Non-singular Matrix
func generateMatrix(matrixSize int) (a, aInv mat.Matrix) {
	for {
		data := make([]float64, matrixSize*matrixSize)
		for i := range data {
			data[i] = float64(rand.Int())
		}
		A := mat.NewDense(matrixSize, matrixSize, data)
		var A_Inv mat.Dense
		err := A_Inv.Inverse(A)
		if err == nil {
			return A, &A_Inv
		}
	}

}

// Generate a Random Permutation
func generatePermutation(matrixSize int) []int {
	Permutation := make([]int, matrixSize)
	for i := 0; i < matrixSize; i++ {
		Permutation[i] = i
	}
	rand.Shuffle(matrixSize, func(i, j int) { Permutation[i], Permutation[j] = Permutation[j], Permutation[i] })
	return Permutation
}
