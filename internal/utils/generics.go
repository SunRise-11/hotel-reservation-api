package utils

import "encoding/json"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 |
		uint32 | uint64 | uintptr | float32 | float64
}

func ConvertWithGenerics[T any](thing T, data []byte) T {
	json.Unmarshal(data, &thing)
	return thing
}