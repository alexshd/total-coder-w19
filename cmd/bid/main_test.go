package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	a := assert.New(t)

	a.Equal(11, 11)
}

func IsEven[T interface{ ~int | ~int32 | ~int64 }](num T) bool {
	return num&1 == 0
}

func IsEven1(num int) bool {
	return num%2 == 0
}

func BenchmarkIsEven(b *testing.B) {
	for i := range b.N {
		IsEven(i)
	}
}

func BenchmarkIsEven1(b *testing.B) {
	for i := range b.N {
		IsEven1(i)
	}
}
