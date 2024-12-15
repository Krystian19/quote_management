package utils

// gets a pointer to anything
func GetPtr[T any](s T) *T {
	return &s
}

// returns the value of a pointer if it's not nil
// otherwise it returns the zero value of the type
func Deref[T any](ptr *T) T {
	if ptr == nil {
		var zeroValue T
		return zeroValue
	}

	return *ptr
}
