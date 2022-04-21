package main

import (
	"fmt"

	"github.com/EduardMarkgraf167/QuickCheck/arbitrary"
)

const QcTestDataSize = 20
const QcStringSize = 30

func QuickCheck[T any](arbitrary arbitrary.Arbitrary[T], property func(value T) bool) {
	for i := 0; i < QcTestDataSize; i++ {
		generatedValue := arbitrary.Generate()
		if property(generatedValue) {
			fmt.Println("+++ OK")
		} else {
			fmt.Println("+++ Failed")
		}
	}
}

func QuickCheckVerbose[T any](arbitrary arbitrary.Arbitrary[T], property func(value T) bool) {
	for i := 0; i < QcTestDataSize; i++ {
		generatedValue := arbitrary.Generate()
		if property(generatedValue) {
			fmt.Println("+++ OK")
		} else {
			fmt.Println("+++ Failed")
		}
		fmt.Printf("\"%v\"\n", generatedValue)
	}
}
