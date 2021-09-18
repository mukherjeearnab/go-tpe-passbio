package main

import (
	"fmt"
	tpe "github.com/mukherjeearnab/gotpe"
	"runtime"
)

func runMemTest() {
	START := 5
	STOP := 20
	INTERVAL := 5

	for N := START; N <= STOP; N = N + INTERVAL {
		fmt.Printf("---------- N = %d\n", N)
		MemTestBench(N)
	}

}

func MemTestBench(N_size int) {
	p_a, p_ta, p_sys := PrintMemUsage("Start . . . .", 0, 0, 0)

	// Init TPE instance
	var TPE tpe.TPE

	// Setup TPE instance
	TPE.Setup(N_size, 28.1)
	p_a, p_ta, p_sys = PrintMemUsage("TPE.Setup Phase", p_a, p_ta, p_sys)

	// Generate a new Secret Key
	TPE.KeyGen(1999)
	p_a, p_ta, p_sys = PrintMemUsage("TPE.KeyGen Phase", p_a, p_ta, p_sys)

	// Create Vector X
	x := make([]float64, N_size)
	for i := range x {
		x[i] = 1
	}
	p_a, p_ta, p_sys = PrintMemUsage("X DEF Phase", p_a, p_ta, p_sys)

	// Encrypt Vector X using Secret Key
	cipher := TPE.Encrypt(x)
	p_a, p_ta, p_sys = PrintMemUsage("TPE.Encrypt Phase", p_a, p_ta, p_sys)

	// Create Vector Y
	y := make([]float64, N_size)
	for i := range x {
		y[i] = 2
	}
	p_a, p_ta, p_sys = PrintMemUsage("Y DEF Phase", p_a, p_ta, p_sys)

	// Generate a new Token using Y and Secret Key
	token := TPE.TokenGen(y)
	p_a, p_ta, p_sys = PrintMemUsage("TPE.TokenGen Phase", p_a, p_ta, p_sys)

	// Decrypt Cipher and obtain result
	TPE.Decrypt(cipher, token)
	p_a, p_ta, p_sys = PrintMemUsage("TPE.Decrypt Phase", p_a, p_ta, p_sys)
}

// Function to print Memory Usage
func PrintMemUsage(testName string, p_a uint64, p_ta uint64, p_sys uint64) (uint64, uint64, uint64) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	// Gen Current Values
	Alloc, TotalAlloc, Sys := bToMb(m.Alloc)-p_a, bToMb(m.TotalAlloc)-p_ta, bToMb(m.Sys)-p_sys
	fmt.Println("TEST: " + testName)
	fmt.Printf("Alloc = %v KB", Alloc)
	fmt.Printf("\tTotalAlloc = %v KB", TotalAlloc)
	fmt.Printf("\tSys = %v KB", m.Sys)
	fmt.Printf("\tNumGC = %v\n", m.NumGC)

	return Alloc, TotalAlloc, Sys
}

// Convert Bytes to MegaBytes
func bToMb(b uint64) uint64 {
	return b / 1024 /// 1024
}
