// Package retry A utility to invoke a method that might fail, but should eventually succeed.
package retry

import (
	"fmt"
	"reflect"
	"runtime"
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
	i := 0
	var err error
	for i < times {
		err = method()
		if err == nil {
			return nil
		}
		fmt.Printf("Invocation %s failed with error: %s, attempt: %d/%d\n", getFunctionName(method), err.Error(), i+1, times)
		i++
	}
	return err
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
