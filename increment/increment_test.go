package increment

import (
	"errors"
	"testing"
)

type version struct {
	version string
	err     error
}

func TestIncrement(t *testing.T) {
	var (
		err      = errors.New("")
		testData = map[string]version{
			"1.0.0":    {"1.0.1", nil},
			"1.0.0.0":  {"1.0.0.1", nil},
			"1.0.0.0a": {"", err},
			"1.0.a0":   {"", err},
			"1.0a.0":   {"", err},
			"1.0.0a":   {"", err},
		}
	)

	t.Run("Increment", func(t *testing.T) {
		for current, expected := range testData {
			ret, err := Increment(current)
			if err != nil {
				if expected.err == nil {
					t.Fatalf("%s should be error but '%s': %s", current, expected.version, err)
				}
			} else {
				if expected.err != nil {
					t.Fatalf("increment %s must be err but %s", current, ret)
				}
				if ret != expected.version {
					t.Fatalf("%s must be expected to %s but '%s'", current, expected.version, ret)
				}
			}
		}
	})
}
