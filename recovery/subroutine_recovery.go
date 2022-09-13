package recovery

import (
	"fmt"
	"log"
)

func Execute(fn func()) {
	go func() {
		defer recoverPanic()
		fn()
	}()
}

func CustomRecoveryExecute(fn func(), recoveryFunc func()) {
	go func() {
		defer recoveryFunc()
		fn()
	}()
}

func recoverPanic() {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		fmt.Println(err)
		newStack := stack(0)
		log.Printf("%s\n", string(newStack))
	}
}
