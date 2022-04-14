package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/EduardMarkgraf167/QuickCheck/generators"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	quickCheckCountMyString()
}

func quickCheckCountMyString() {
	fmt.Println("prop1_myString")
	QuickCheckVerbose[string](generators.NewStringGenerator(QcStringSize), prop1MyString)
	fmt.Println("prop2_myString")
	QuickCheckVerbose[string](generators.NewStringGenerator(QcStringSize), prop2MyString)
	fmt.Println("prop3_myString")
	QuickCheckVerbose[string](generators.NewStringGenerator(QcStringSize), prop3MyString)
}

func prop1MyString(myString string) bool {
	return len(myString) >= 0
}

func prop2MyString(myString string) bool {
	reversed := reverseString(myString)
	return len(reversed) == len(myString)
}

func prop3MyString(myString string) bool {
	doubledMyString := myString + myString
	return len(doubledMyString) == 2*len(myString)
}
