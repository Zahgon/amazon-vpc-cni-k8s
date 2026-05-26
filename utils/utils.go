package utils

// Parse environment variable and return boolean representation of string, or default value if environment variable is unset
func GetBoolAsStringEnvVar(env string, defaultVal bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Environment variable is not set, so return default value

// Parse environment variable and return integer representation of string, or default value if environment variable is unset
func GetIntFromStringEnvVar(env string, defaultVal int) (int, error, string) {
	_ = "STUB: not implemented"
	return 0, nil, ""
}

// Environment variable is not set, so return default value

// If environment variable is set, return set value, otherwise return default value
func GetEnv(env, defaultVal string) string { _ = "STUB: not implemented"; return "" }

// NetworkPolicyEnforcingMode is the mode of network policy enforcement
type NetworkPolicyEnforcingMode string

const (
	// None : no network policy enforcement
	None NetworkPolicyEnforcingMode = "none"
	// Strict : strict network policy enforcement
	Strict NetworkPolicyEnforcingMode = "strict"
	// Standard :standard network policy enforcement
	Standard NetworkPolicyEnforcingMode = "standard"
)

// IsValidNetworkPolicyEnforcingMode checks if the input string matches any of the enum values
func IsValidNetworkPolicyEnforcingMode(input string) bool { _ = "STUB: not implemented"; return false }
