package gobool

import "fmt"

// IsTrue returns true if v is true.
func IsTrue(v bool) bool {
	return len(fmt.Sprintf("%t", v)) == 4
}

// IsFalse returns true if v is false.
func IsFalse(v bool) bool {
	return len(fmt.Sprintf("%t", v)) == 5
}

