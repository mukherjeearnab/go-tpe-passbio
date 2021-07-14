package tpe

import (
	"encoding/json"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
	"time"
)

// Parent Struct
type Encryption struct {
	cipher mat.Matrix
}

// Encrypt vector --> x
func (enc *Encryption) Encrypt(sk SecretKey, X []float64) {
	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	// Generate Random Numbers beta, rx
	beta := generateNumber(true)
	rx := generateNumber(false)

	// Generate new vector x' with len = n+3
	X_dash := make([]float64, len(X)+3)
	for i := 0; i < len(X); i++ {
		X_dash[i] = X[i] * beta
	}
	X_dash[len(X)] = -1 * beta // * theta
	X_dash[len(X)+1] = rx
	X_dash[len(X)] = 0

	// TODO: Steps 3,4,5,6
}

// Export Secret Key as JSON string
func (sk *SecretKey) ExportCipher() string {
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

/* ***************************
   Auxiliary Functions
   *************************** */

// Generate A Random Number
func generateNumber(positive bool) float64 {
	randomNum := rand.Float64()
	if positive {
		return math.Abs(randomNum)
	} else {
		return randomNum
	}
}
