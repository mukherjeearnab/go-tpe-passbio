package tpe

import (
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
func (tpe *TPE) Encrypt(X []float64) string {
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
	X_dash[len(X)] = 0

	// Step 3: Permute x' to pi(x') = x''
	X_dd := make([]float64, len(X_dash))
	for i := 0; i < len(X_dash); i++ {
		X_dd[tpe.key.Pi[i]] = X_dash[i]
	}

	// TODO: Steps 4,5,6

	return ""
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
