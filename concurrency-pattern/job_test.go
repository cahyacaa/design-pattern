package concurrency_pattern

import (
	"fmt"
	"testing"
)

func BenchmarkWithoutPipelineModule(b *testing.B) {
	for i := 0; i < 5; i++ {
		fmt.Println(addComments(sqrt[int, int](square[int, int](multiplyTwo[int, int](i)))))
	}
}

func BenchmarkWithPipelineModule(b *testing.B) {
	outC := New(func(inC chan<- any) {
		defer close(inC)
		for i := 0; i < 5; i++ {
			inC <- i
		}
	}).
		Pipe(func(in any) (any, error) {
			return multiplyTwo[int, int](in.(int)), nil
		}).
		Pipe(func(in any) (any, error) {
			return square[int, int](in.(int)), nil
		}).
		Pipe(func(in any) (any, error) {
			return sqrt[int, int](in.(int)), nil
		}).
		Pipe(func(i any) (any, error) {
			return addComments[int, string](i.(int)), nil
		}).
		Merge()

	for items := range outC {
		fmt.Println(items)
		// empty out channel
	}
}
