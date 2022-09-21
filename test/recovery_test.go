package test

import (
	Recovery "bitbucket.org/coinswitch/go-bricks/recovery"
	"fmt"
	"testing"
	"time"
)

func RaisesPanic() {
	names := []string{
		"lobster",
		"sea urchin",
		"sea cucumber",
	}
	fmt.Println("My favorite sea creature is:", names[len(names)])
}

func PrintPanicError(err error) {
	fmt.Println("callback function got error", err)
}

func TestGoRoutineRecoveryWithCallback(t *testing.T) {
	Recovery.GoRoutineCustomRecovery(RaisesPanic, PrintPanicError)
	time.Sleep(10 * time.Second)
}

func TestGoRoutineRecovery(t *testing.T) {
	Recovery.GoRoutineRecovery(RaisesPanic)
	time.Sleep(10 * time.Second)
}
