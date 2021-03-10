package terraform

import "strings"

// IsModuleAddress determines if passed string is a module address and
// returns true if thats the case
func IsModuleAddress(a string) bool {
	return isSpecificAddressType(a, func(parts []string) bool { return "module" == parts[0] })
}

// IsDataAddress determines if passed string is a data address and
// returns true if thats the case
func IsDataAddress(a string) bool {
	return isSpecificAddressType(a, func(parts []string) bool { return "data" == parts[0] })
}

// IsManagedAddress determines if passed string is a managed resource address and
// returns true if thats the case
func IsManagedAddress(a string) bool {
	return isSpecificAddressType(a, func(parts []string) bool { return 3 > len(parts) })
}

func isSpecificAddressType(a string, v func([]string) bool) bool {
	if parts := strings.Split(a, "."); 1 < len(parts) {
		return v(parts)
	}
	return false
}
