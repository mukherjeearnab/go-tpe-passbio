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
	PrintMemUsage("Start . . . .")

	// Init TPE instance
	var TPE tpe.TPE

	// Setup TPE instance
	TPE.Setup(N_size, 28.1)
	PrintMemUsage("TPE.Setup Phase")

	// Generate a new Secret Key
	TPE.KeyGen(1999)
	PrintMemUsage("TPE.KeyGen Phase")

	// Create Vector X
	x := make([]float64, N_size)
	for i := range x {
		x[i] = 1
	}
	PrintMemUsage("X DEF Phase")

	// Encrypt Vector X using Secret Key
	cipher := TPE.Encrypt(x)
	PrintMemUsage("TPE.Encrypt Phase")

	// Create Vector Y
	y := make([]float64, N_size)
	for i := range x {
		y[i] = 2
	}
	PrintMemUsage("Y DEF Phase")

	// Generate a new Token using Y and Secret Key
	token := TPE.TokenGen(y)
	PrintMemUsage("TPE.TokenGen Phase")

	// Decrypt Cipher and obtain result
	TPE.Decrypt(cipher, token)
	PrintMemUsage("TPE.Decrypt Phase")
}

// Function to print Memory Usage
func PrintMemUsage(testName string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Println("TEST: " + testName)
	fmt.Printf("Alloc = %v KB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v KB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v KB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

// Convert Bytes to MegaBytes
func bToMb(b uint64) uint64 {
	return b / 1024 /// 1024
}
