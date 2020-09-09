package retry_test

import (
	"github.com/ecnepsnai/retry"
)

func ExampleTry() {
	err := retry.Try(func() error {
		// Invoke a method that may initially fail, but should eventually suceed
		return nil
	}, 5)
	if err != nil {
		// After 5 attempts it never suceeded
	}
}

func ExampleTryAsync() {
	retry.TryAsync(func() error {
		// Invoke a method that may initially fail, but should eventually suceed
		return nil
	}, 5, func(err error) {
		if err != nil {
			// After 5 attempts it never succeeded
		}
	})
}
