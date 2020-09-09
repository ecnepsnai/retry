package retry_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ecnepsnai/retry"
)

func TestTry(t *testing.T) {
	i := 0
	err := retry.Try(func() error {
		if i < 5 {
			i++
			return fmt.Errorf("Nope")
		}
		return nil
	}, 6)
	if i != 5 || err != nil {
		t.Fail()
	}
}

func TestTryAsync(t *testing.T) {
	finished := false
	i := 0
	retry.TryAsync(func() error {
		if i < 5 {
			i++
			return fmt.Errorf("Nope")
		}
		return nil
	}, 6, func(err error) {
		if err != nil {
			t.Fail()
		}
		finished = true
	})
	time.Sleep(1 * time.Second)
	if !finished {
		t.Fail()
	}
}
