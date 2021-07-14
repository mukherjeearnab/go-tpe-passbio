package tpe

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

// Parent Struct
type SecretKey struct {
	key Key
}

// Key Struct
type Key struct {
	M_1  mat.Matrix `json:"m1"`
	M_2  mat.Matrix `json:"m2"`
	M_1i mat.Matrix `json:"m1i"`
	M_2i mat.Matrix `json:"m2i"`
	Pi   []int      `json:"pi"`
}

// Export Key Format Struct
type exportKey struct {
	M_1  []byte `json:"m1"`
	M_2  []byte `json:"m2"`
	M_1i []byte `json:"m1i"`
	M_2i []byte `json:"m2i"`
	Pi   []int  `json:"pi"`
}

// Generate New Secret Key
func (sk *SecretKey) KeyGen(seed int64, matrixSize int) {
	// Set Random Seed
	rand.Seed(seed)

	// Generate Random Non-singular Matrices
	sk.key.M_1, sk.key.M_1i = generateMatrix(matrixSize)
	sk.key.M_2, sk.key.M_2i = generateMatrix(matrixSize)

	// Generate Permutation 'pi'
	sk.key.Pi = generatePermutation(matrixSize)
}

// Export Secret Key as JSON string
func (sk *SecretKey) ExportKey() string {
	// Marshal mat.Matrix to byte[]
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

	// Marshal byte[] struct to JSON string
	expKey := exportKey{M_1Bytes, M_1iBytes, M_2Bytes, M_2iBytes, sk.key.Pi}
	keyJSON, _ := json.Marshal(expKey)
	return string(keyJSON)
}

// Import Secret Key from JSON string
func (sk *SecretKey) ImportKey(JSON string) {
	// JSON Unmarshal
	var importKey exportKey
	json.Unmarshal([]byte(JSON), &importKey)

	// Unmarshal byte[] to mat.Dense
	var M_1 mat.Dense
	M_1.UnmarshalBinary(importKey.M_1)
	var M_1i mat.Dense
	M_1i.UnmarshalBinary(importKey.M_1i)
	var M_2 mat.Dense
	M_2.UnmarshalBinary(importKey.M_2)
	var M_2i mat.Dense
	M_2i.UnmarshalBinary(importKey.M_2i)

	// Set sk Vars
	sk.key.M_1 = &M_1
	sk.key.M_1i = &M_1i
	sk.key.M_2 = &M_2
	sk.key.M_2i = &M_2i
	sk.key.Pi = importKey.Pi
}

/* ***************************
   Auxiliary Functions
   *************************** */

// Generate A Random Non-singular Matrix
func generateMatrix(matrixSize int) (a, aInv mat.Matrix) {
	for {
		// Generate Random slice
		data := make([]float64, matrixSize*matrixSize)
		for i := range data {
			data[i] = float64(rand.Int())
		}

		// Generate Random Matrix
		A := mat.NewDense(matrixSize, matrixSize, data)

		// Check of Inverse Exists
		var A_Inv mat.Dense
		err := A_Inv.Inverse(A)

		// If Inverse Exists, return the Matrix
		if err == nil {
			return A, &A_Inv
		}
	}

}

// Generate a Random Permutation
func generatePermutation(matrixSize int) []int {
	// Generate Ordered Slice
	Permutation := make([]int, matrixSize)
	for i := 0; i < matrixSize; i++ {
		Permutation[i] = i
	}

	// Randomize the Order to generate Permutation
	rand.Shuffle(matrixSize, func(i, j int) { Permutation[i], Permutation[j] = Permutation[j], Permutation[i] })
	return Permutation
}
