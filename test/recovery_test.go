package test

import (
	Recovery "bitbucket.org/coinswitch/go-bricks/recovery"
	"fmt"
	"testing"
)

func RaisesPanic() {
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}
	fmt.Println("My favorite sea creature is:", names[len(names)])
}

func TestGoRoutineRecovery(t *testing.T) {
	Recovery.GoRoutineRecovery(RaisesPanic)
	//Recovery.GoRoutineRecovery(RaisesPanic)
	fmt.Println("go routine function didn't raise panic")
}
