package v3

type MSSResponse struct {
	ActualMss     int  `json:"actualMss"` // exclude TCP option(timestample)
	Mss           int  `json:"mss"`
	IsIPv4        bool `json:"isIPv4"`
	EstinamtedMtu int  `json:"estimatedMtu"`
}
