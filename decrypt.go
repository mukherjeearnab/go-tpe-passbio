package tpe

import (
	"encoding/base64"
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
