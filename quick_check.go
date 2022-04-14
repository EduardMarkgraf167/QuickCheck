package main

import (
	"fmt"

	"github.com/EduardMarkgraf167/QuickCheck/generators"
)

const QcTestDataSize = 20
const QcStringSize = 30

func QuickCheck[T any](generator generators.Generator[T], property func(value T) bool) {
	for i := 0; i < QcTestDataSize; i++ {
		generatedValue := generator.Generate()
		if property(generatedValue) {
			fmt.Println("+++ OK")
		} else {
			fmt.Println("+++ Failed")
		}
	}
}

func QuickCheckVerbose[T any](generator generators.Generator[T], property func(value T) bool) {
	for i := 0; i < QcTestDataSize; i++ {
		generatedValue := generator.Generate()
		if property(generatedValue) {
			fmt.Println("+++ OK")
		} else {
			fmt.Println("+++ Failed")
		}
		fmt.Printf("\"%v\"\n", generatedValue)
	}
}
