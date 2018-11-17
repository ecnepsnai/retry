/*
Package retry A utility to invoke a method that might fail, but should eventually succeed.

Simple retry:

	retry.Try(func() error {
		return SomethingThatMightFail()
	}, 5)

Asynchronous retry:

	retry.TryAsync(func() error {
		return SomethingThatMightFail()
	}, 5, func(err error) {
		if err != nil {
			panic(err.Error())
		}
	})
*/
package retry

import (
	"reflect"
	"runtime"

	"github.com/ecnepsnai/logtic"
)

// TryAsync try to invoke the given method asynchronously. If unsuccessful, retry for the specified number of tries.
func TryAsync(method func() error, times int, finished func(error)) {
	go func() {
		err := Try(method, times)
		finished(err)
	}()
}

// Try try to invoke the given method. If unsuccessful, retry for the specified number of tries.
func Try(method func() error, times int) error {
	functionName := getFunctionName(method)
	log := logtic.Connect("retry:" + functionName)
	i := 0
	var err error
	for i < times {
		err = method()
		if err == nil {
			return nil
		}
		log.Warn("Invocation %s failed with error: %s, attempt: %d/%d\n", functionName, err.Error(), i+1, times)
		i++
	}
	return err
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
