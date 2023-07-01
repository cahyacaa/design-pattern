package concurrency_pattern

import (
	"fmt"
	"math"
	"time"
)

type Calculate interface {
	int | float64
}

func multiplyTwo[T, U Calculate](v T) U {
	time.Sleep(1 * time.Second)
	return U(v * 2)
}

func square[T, U Calculate](v T) U {
	time.Sleep(2 * time.Second)
	return U(v * v)
}
func sqrt[T, U Calculate](v T) U {
	time.Sleep(1 * time.Second)
	return U(math.Pow(float64(v), 3))
}

func addComments[T any, U string](v T) U {
	return U(fmt.Sprintf("%v-test", v))
}
