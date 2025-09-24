package utils

import "time"

func PtrString(s string) *string {
	return &s
}

func PtrInt(i int) *int {
	return &i
}

func PtrUint(i uint) *uint {
	return &i
}

func PtrFloat64(f float64) *float64 {
	return &f
}

func PtrBool(b bool) *bool {
	return &b
}

func PtrTime(b time.Time) *time.Time {
	return &b
}

func PtrInterface(i interface{}) *interface{} {
	return &i
}
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
