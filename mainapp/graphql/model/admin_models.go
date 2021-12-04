package model

type Admin struct {
	BusinessName    string `json:"business_name"`
	DefaultTimezone string `json:"default_timezone"`
	SetupComplete   bool   `json:"setup_complete`
	Version         string `json:"version`
}
