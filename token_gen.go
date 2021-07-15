package tpe

import (
	"encoding/base64"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"time"
)

// TokenGen vector --> y
func (tpe *TPE) TokenGen(Y []float64) string {
	// Initial Check (Check Vector Length)
	if len(Y) != tpe.setup.N {
		fmt.Printf("ERROR! Vector (%d) length not equal to Set (%d) Length.\n", len(Y), tpe.setup.N)
		return "ERROR!"
	}

	// Set Random Seed
	rand.Seed(time.Now().UnixNano())

	// Step 1: Generate Random Numbers alpha, ry
	alpha := generateNumber(true)
	ry := generateNumber(false)

	// Step 2: Generate new vector y' with len = n+3
	Y_dash := make([]float64, len(Y)+3)
	for i := 0; i < len(Y); i++ {
		Y_dash[i] = Y[i] * alpha
	}
	Y_dash[len(Y)] = alpha
	Y_dash[len(Y)+1] = 0
	Y_dash[len(Y)+2] = ry

	// Step 3: Permute y' to pi(y') = y''
	Y_dd := make([]float64, len(Y_dash))
	for i := 0; i < len(Y_dash); i++ {
		Y_dd[i] = Y_dash[tpe.key.Pi[i]]
	}

	// Step 4: Transform y'' to a Diagonal Matrix with diag(Matrix) = y''
	Y_Diag := mat.NewDiagDense(len(Y_dd), Y_dd)

	// Step 5: Generate a Random Lower Triangular Matrix with diagonal elements = 1
	_Sy := generateLTMatrix(tpe.setup.N + 3)
	var Sy mat.Dense
	Sy.CloneFrom(_Sy)
	for i := 0; i < tpe.setup.N+3; i++ {
		Sy.Set(i, i, 1)
	}

	// Step 6: Compute Token Ty = M_2i Y_Diag Sy M_1i
	var Ty mat.Dense
	Ty.Mul(tpe.key.M_2i, Y_Diag)
	Ty.Mul(&Ty, &Sy)
	Ty.Mul(&Ty, tpe.key.M_1i)

	// Convert Ty to Byte[]
	TyBytes, err := Ty.MarshalBinary()
	if err != nil {
		fmt.Printf("BinMarshall Error: %v", err)
		return "ERROR!"
	}

	// Return Base64-String of TyBytes
	return base64.StdEncoding.EncodeToString(TyBytes)
}

/* ***************************
   Auxiliary Functions
   *************************** */
// NONE
