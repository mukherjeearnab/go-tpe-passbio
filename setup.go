package tpe

import (
	"encoding/json"
)

// Parent Struct
type TPE struct {
	setup SetupConfig
	key   Key
}

// Config Struct
type SetupConfig struct {
	N     int     `json:"n"`
	Theta float64 `json:"theta"`
}

// Set n and theta
func (tpe *TPE) Setup(n int, theta float64) {
	tpe.setup.N = n
	tpe.setup.Theta = theta
}

// Get n and theta
func (tpe *TPE) GetConfig() (n int, theta float64) {
	return tpe.setup.N, tpe.setup.Theta
}

// Export Setup Vars as JSON string
func (tpe *TPE) ExportSetup() string {
	// Marshal Struct to JSON string
	configJSON, _ := json.Marshal(tpe.setup)
	return string(configJSON)
}

// Import Setup Vars from JSON string
func (tpe *TPE) ImportSetup(JSON string) {
	// Unmarshal and load Config
	json.Unmarshal([]byte(JSON), &tpe.setup)
}
