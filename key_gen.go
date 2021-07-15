package tpe

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

// key Struct
type key struct {
	M_1  mat.Matrix `json:"m1"`
	M_2  mat.Matrix `json:"m2"`
	M_1i mat.Matrix `json:"m1i"`
	M_2i mat.Matrix `json:"m2i"`
	Pi   []int      `json:"pi"`
}

// Export key Format Struct
type exportKey struct {
	M_1  []byte `json:"m1"`
	M_2  []byte `json:"m2"`
	M_1i []byte `json:"m1i"`
	M_2i []byte `json:"m2i"`
	Pi   []int  `json:"pi"`
}

// Generate New Secret key
func (tpe *TPE) KeyGen(seed int64) {
	// Set Random Seed
	rand.Seed(seed)

	// Step 1: Generate Random Non-singular Matrices
	tpe.key.M_1, tpe.key.M_1i = generateMatrix(tpe.setup.N + 3)
	tpe.key.M_2, tpe.key.M_2i = generateMatrix(tpe.setup.N + 3)

	// Step 2: Generate Permutation 'pi'
	tpe.key.Pi = generatePermutation(tpe.setup.N + 3)
}

// Export Secret key as JSON string
func (tpe *TPE) ExportKey() string {
	// Marshal mat.Matrix to byte[]
	M_1Bytes, err := mat.DenseCopyOf(tpe.key.M_1).MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return ""
	}
	M_1iBytes, err := mat.DenseCopyOf(tpe.key.M_1i).MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return ""
	}
	M_2Bytes, err := mat.DenseCopyOf(tpe.key.M_2).MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return ""
	}
	M_2iBytes, err := mat.DenseCopyOf(tpe.key.M_2i).MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return ""
	}

	// Marshal byte[] struct to JSON string
	expKey := exportKey{M_1Bytes, M_1iBytes, M_2Bytes, M_2iBytes, tpe.key.Pi}
	keyJSON, _ := json.Marshal(expKey)
	return string(keyJSON)
}

// Import Secret key from JSON string
func (tpe *TPE) ImportKey(JSON string) {
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

	// Initial Check for Dimentions
	c_r, c_c := M_1.Dims()
	t_r, t_c := M_2.Dims()
	if c_r != c_c || t_r != t_c {
		fmt.Println("ERROR! Matrices need to be Square Matrices.")
		return
	}
	if c_r != t_r || c_r != tpe.setup.N+3 || t_r != tpe.setup.N+3 {
		fmt.Println("ERROR! Matrices need same Dimentions |OR| Matrices need to be of Set Dimentions.")
		return
	}
	if len(importKey.Pi) != tpe.setup.N+3 {
		fmt.Println("ERROR! Permutation Must be of Set Dimentions.")
		return
	}

	// Set sk Vars
	tpe.key.M_1 = &M_1
	tpe.key.M_1i = &M_1i
	tpe.key.M_2 = &M_2
	tpe.key.M_2i = &M_2i
	tpe.key.Pi = importKey.Pi
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
			data[i] = rand.Float64()
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
