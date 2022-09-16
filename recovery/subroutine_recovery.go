package Recovery

import (
	"fmt"
	"log"
)

func GoRoutineRecoveryExecute(fn func()) {
	go func() {
		defer recoverFromPanic(nil)
		fn()
	}()
}

func GoRoutineCustomRecoveryExecute(fn func(), recoveryFunc func(err error)) {
	go func() {
		defer recoverFromPanic(recoveryFunc)
		fn()
	}()
}

func recoverFromPanic(callback func(err error)) {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("%v", r)
		}
		fmt.Println(err)
		newStack := stack(0)
		log.Println("%s\n", string(newStack))
		if callback != nil {
			callback(err)
		}
	}
}
