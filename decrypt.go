package tpe

import (
	"encoding/base64"
	"fmt"
	"gonum.org/v1/gonum/mat"
)

// Encrypt vector --> x
func (tpe *TPE) Decrypt(Cx_str string, Ty_str string) int {
	// Load Matrices Cx and Ty from Base64-String
	var Cx mat.Dense
	CxBytes, _ := base64.StdEncoding.DecodeString(Cx_str)
	Cx.UnmarshalBinary(CxBytes)

	var Ty mat.Dense
	TyBytes, _ := base64.StdEncoding.DecodeString(Ty_str)
	Ty.UnmarshalBinary(TyBytes)

	// Initial Check for Dimentions
	c_r, c_c := Cx.Dims()
	t_r, t_c := Ty.Dims()
	if c_r != c_c || t_r != t_c {
		fmt.Println("ERROR! Matrices need to be Square Matrices.")
		return -1
	}
	if c_r != t_r || c_r != tpe.setup.N+3 || t_r != tpe.setup.N+3 {
		fmt.Println("ERROR! Matrices need same Dimentions |OR| Matrices need to be of Set Dimentions.")
		return -1
	}

	// Step 1: Compute I = Trace(Cx Ty)
	var CxTy mat.Dense
	CxTy.Mul(&Cx, &Ty)
	I := CxTy.Trace()

	// Step 2: Result = 1 if I <= 0, else Result = 0
	var Result int
	if I <= 0 {
		Result = 1
	} else {
		Result = 0
	}

	return Result
}

/* ***************************
   Auxiliary Functions
   *************************** */
// NONE
