package helpers

import (
	"errors"
	"reflect"
)

// MARK: - Reflection Helpers
// Deprecated: incorrect behavior under certain conditions.
func PointerElement(rv reflect.Value) (reflect.Value, error) {
	el := rv

	for el.Kind() == reflect.Pointer {
		if el.IsNil() {
			return el, errors.New("nil pointer")
		}

		el = el.Elem()
	}

	return el, nil
}

// Deprecated: switch to PointerTo
func PointTo[T any](data T) *T { return &data }

// PointerTo - returns a reference to value.
func PointerTo[T any](value T) *T { return &value }

// Deprecated: switch to UnwrapOrDefault
func UnwrapPointer[T any](data *T, fallback T) T {
	if data != nil {
		return *data
	}

	return fallback
}

// UnwrapOrDefault - returns the value of the optional argument or a fallback value if the optional is nil.
func UnwrapOrDefault[T any](optional *T, fallback T) T {
	if optional != nil {
		return *optional
	}

	return fallback
}
