/*
Package retry A utility to invoke a method that might fail, but should eventually succeed.
*/
package retry

import (
	"reflect"
	"runtime"

	"github.com/ecnepsnai/logtic"
)

// TryAsync will invoke the given method asynchronously. If the method returns an error it will retry up-to the
// specified number of attempts. If every invocation returns an error the last error is provided in the finished
// callback. If successful the error is nil.
func TryAsync(method func() error, times int, finished func(error)) {
	go func() {
		err := Try(method, times)
		finished(err)
	}()
}

// Try will invoke the given method and wait for it to return. If the method returns an error it will retry up-to the
// specified number of attempts. If every invocation returns an error the last error is returned.
// If successful nil is returned.
func Try(method func() error, times int) error {
	functionName := getFunctionName(method)
	log := logtic.Log.Connect("retry(" + functionName + ")")
	i := 0
	var err error
	for i < times {
		err = method()
		if err == nil {
			return nil
		}
		log.Warn("Invoking function %s failed with error: %s, attempt: %d/%d\n", functionName, err.Error(), i+1, times)
		i++
	}
	return err
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
