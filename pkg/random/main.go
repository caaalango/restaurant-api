package random

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	DefaultStringLength = 10
	DefaultDateStart    = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	DefaultDateEnd      = time.Now()
	DefaultFloatMin     = 0.0
	DefaultFloatMax     = 100.0
	DefaultIntMin       = 0
	DefaultIntMax       = 100
)

func Enum[T any]() T {
	values := getEnumValues[T]()
	if len(values) == 0 {
		panic("no values registered for this enum type")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return values[r.Intn(len(values))]
}

func String(n ...int) string {
	length := DefaultStringLength
	if len(n) > 0 {
		length = n[0]
	}
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func Date(startEnd ...time.Time) time.Time {
	start, end := DefaultDateStart, DefaultDateEnd
	if len(startEnd) > 0 {
		start = startEnd[0]
	}
	if len(startEnd) > 1 {
		end = startEnd[1]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	duration := end.Sub(start)
	randomDuration := time.Duration(r.Int63n(int64(duration)))
	return start.Add(randomDuration)
}

func Float(minMax ...float64) float64 {
	min, max := DefaultFloatMin, DefaultFloatMax
	if len(minMax) > 0 {
		min = minMax[0]
	}
	if len(minMax) > 1 {
		max = minMax[1]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Float64()*(max-min)
}

func Int(minMax ...int) int {
	min, max := DefaultIntMin, DefaultIntMax
	if len(minMax) > 0 {
		min = minMax[0]
	}
	if len(minMax) > 1 {
		max = minMax[1]
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func Bool() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(2) == 1
}

var registry = make(map[string][]interface{})

func getEnumValues[T any]() []T {
	typeName := fmt.Sprintf("%T", new(T))
	if values, exists := registry[typeName]; exists {
		result := make([]T, len(values))
		for i, value := range values {
			result[i] = value.(T)
		}
		return result
	}
	return nil
}
