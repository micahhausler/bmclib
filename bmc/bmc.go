package bmc

import "fmt"

// Metadata represents details about a bmc method
type Metadata struct {
	// SuccessfulProvider is the name of the provider that successfully executed
	SuccessfulProvider string
	// ProvidersAttempted is a slice of all providers that were attempt to execute
	ProvidersAttempted []string
	// SuccessfulOpenConns is a slice of provider names that were opened successfully
	SuccessfulOpenConns []string
	// SuccessfulCloseConns is a slice of provider names that were closed successfully
	SuccessfulCloseConns []string
}

func getProviderName(provider interface{}) string {
	switch p := provider.(type) {
	case Provider:
		return p.Name()
	}
	return fmt.Sprintf("%T", provider)
}
