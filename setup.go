package tpe

import (
	"encoding/json"
)

type Config struct {
	config SetupConfig
}

type SetupConfig struct {
	N     int     `json:"n"`
	Theta float64 `json:"theta"`
}

// Set n and theta
func (c *SetupConfig) Setup(n int, theta float64) {
	c.N = n
	c.Theta = theta
}

// Get n and theta
func (c *SetupConfig) GetConfig() (n int, theta float64) {
	return c.N, c.Theta
}

// Export Setup Vars as JSON string
func (c *SetupConfig) ExportSetup() string {
	configJSON, _ := json.Marshal(c)
	return string(configJSON)
}

// Import Setup Vars from JSON string
func (c *SetupConfig) ImportSetup(JSON string) {
	json.Unmarshal([]byte(JSON), c)
}
