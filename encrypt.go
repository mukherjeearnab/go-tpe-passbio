package tpe

import (
	"encoding/base64"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"math/rand"
	"time"
)

// Encrypt vector --> x
func (tpe *TPE) Encrypt(X []float64) string {
	// Initial Check (Check Vector Length)
	if len(X) != tpe.setup.N {
		fmt.Printf("ERROR! Vector (%d) length not equal to Set (%d) Length.\n", len(X), tpe.setup.N)
		return "ERROR!"
	}

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	// Step 1: Generate Random Numbers beta, rx
	beta := generateNumber(true)
	rx := generateNumber(false)

	// Step 2: Generate new vector x' with len = n+3
	X_dash := make([]float64, len(X)+3)
	for i := 0; i < len(X); i++ {
		X_dash[i] = X[i] * beta
	}
	X_dash[len(X)] = -1 * beta * tpe.setup.Theta
	X_dash[len(X)+1] = rx
	X_dash[len(X)+2] = 0

	// Step 3: Permute x' to pi(x') = x''
	X_dd := make([]float64, len(X_dash))
	for i := 0; i < len(X_dash); i++ {
		X_dd[i] = X_dash[tpe.key.Pi[i]]
	}

	// Step 4: Transform x'' to a Diagonal Matrix with diag(Matrix) = x''
	X_Diag := mat.NewDiagDense(len(X_dd), X_dd)

	// Step 5: Generate a Random Lower Triangular Matrix with diagonal elements = 1
	_Sx := generateLTMatrix(tpe.setup.N + 3)
	var Sx mat.Dense
	Sx.CloneFrom(_Sx)
	for i := 0; i < tpe.setup.N+3; i++ {
		Sx.Set(i, i, 1)
	}

	// Step 6: Compute Cipher = M_1 Sx X_Diag M_2
	var Cx mat.Dense
	Cx.Mul(tpe.key.M_1, &Sx)
	Cx.Mul(&Cx, X_Diag)
	Cx.Mul(&Cx, tpe.key.M_2)

	// Convert Cx to Byte[]
	CxBytes, err := Cx.MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return "ERROR!"
	}

	// Return Base64-String of CxBytes
	return base64.StdEncoding.EncodeToString(CxBytes)
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

// Generate A Random Lower Triangular Matrix
func generateLTMatrix(matrixSize int) mat.Matrix {
	for {
		// Generate Random slice
		data := make([]float64, matrixSize*matrixSize)
		for i := range data {
			data[i] = rand.Float64()
		}

		// Generate Random Matrix
		A := mat.NewTriDense(matrixSize, false, data)

		return A
	}
}
