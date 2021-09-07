package main

import (
	"fmt"
	tpe "github.com/mukherjeearnab/gotpe"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Create("benchmark.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Size,Encrypt,Token,Decrypt")

	if err2 != nil {
		log.Fatal(err2)
	}

	for N := 100; N <= 2000; N = N + 100 {
		T_Enc := int64(0)
		T_Tok := int64(0)
		T_Dec := int64(0)
		for T := 0; T < 10; T++ {
			e, t, d := runTest(N)
			T_Enc = T_Enc + e
			T_Tok = T_Tok + t
			T_Dec = T_Dec + d
		}
		T_Enc = T_Enc / 10
		T_Tok = T_Tok / 10
		T_Dec = T_Dec / 10
		out := fmt.Sprintf("%d,%d,%d,%d\n", N, T_Enc, T_Tok, T_Dec)

		_, err2 := f.WriteString(out)

		if err2 != nil {
			log.Fatal(err2)
		}
	}

}

func runTest(N_size int) (int64, int64, int64) {
	// Create a seed
	seed := time.Now().UnixNano()

	// Init TPE instance
	var TPE tpe.TPE

	// Setup TPE instance
	TPE.Setup(N_size, 28.1)

	// Generate a new Secret Key
	TPE.KeyGen(seed)

	// Print Secret Key
	// fmt.Println("Secret Key: " + TPE.ExportKey())

	// Create Vector X
	x := make([]float64, N_size)
	for i := range x {
		x[i] = 1
	}

	//----------------------------------------------------------------
	// BENCHMARK READ 1 => START (RECORD ENCRYPTION TIME)
	//----------------------------------------------------------------

	BenchmarkR1Start := time.Now()
	// Encrypt Vector X using Secret Key
	cipher := TPE.Encrypt(x)
	BenchmarkR1Elapsed := time.Since(BenchmarkR1Start)
	log.Printf("Encrypt took %d", BenchmarkR1Elapsed.Milliseconds())

	//----------------------------------------------------------------
	// BENCHMARK READ 1 => END
	//----------------------------------------------------------------

	// Create Vector Y
	y := make([]float64, N_size)
	for i := range x {
		y[i] = 2
	}

	//----------------------------------------------------------------
	// BENCHMARK READ 2 => START (RECORD TOKEN GEN TIME)
	//----------------------------------------------------------------

	BenchmarkR2Start := time.Now()
	// Generate a new Token using Y and Secret Key
	token := TPE.TokenGen(y)
	BenchmarkR2Elapsed := time.Since(BenchmarkR2Start)
	log.Printf("Token took %d", BenchmarkR2Elapsed.Milliseconds())

	//----------------------------------------------------------------
	// BENCHMARK READ 2 => END
	//----------------------------------------------------------------

	//----------------------------------------------------------------
	// BENCHMARK READ 3 => START (RECORD DECRYPT TIME)
	//----------------------------------------------------------------

	BenchmarkR3Start := time.Now()
	// Decrypt Cipher and obtain result
	TPE.Decrypt(cipher, token)
	BenchmarkR3Elapsed := time.Since(BenchmarkR3Start)
	log.Printf("Decrypt took %d", BenchmarkR3Elapsed.Milliseconds())

	//----------------------------------------------------------------
	// BENCHMARK READ 3 => END
	//----------------------------------------------------------------

	// Return Readings
	return BenchmarkR1Elapsed.Milliseconds(), BenchmarkR2Elapsed.Milliseconds(), BenchmarkR3Elapsed.Milliseconds()
}
