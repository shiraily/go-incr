package increment

import (
	"errors"
	"testing"
)

type version struct {
	expected string
	target   VersionNumber
	err      error
}

func TestIncrement(t *testing.T) {
	var (
		err      = errors.New("")
		testData = map[string]version{
			"1.0.0":    {"1.0.1", Patch, nil},
			"1.0.0.1":  {"1.0.0.2", Build, nil},
			"1.0.1.2":  {"1.0.2.0", Patch, nil},
			"0.1.2.3":  {"0.2.0.0", Minor, nil},
			"1.2.3.4":  {"2.0.0.0", Major, nil},
			"1.0.0.0a": {"", Patch, err},
			"1.0.a0":   {"", Patch, err},
			"1.0a.0":   {"", Patch, err},
			"1.0.0a":   {"", Patch, err},
		}
	)

	t.Run("Increment", func(t *testing.T) {
		for current, data := range testData {
			ret, err := Increment(current, data.target)
			if err != nil {
				if data.err == nil {
					t.Fatalf("%s should be error but '%s': %s", current, data.expected, err)
				}
			} else {
				if data.err != nil {
					t.Fatalf("increment %s must be err but %s", current, ret)
				}
				if ret != data.expected {
					t.Fatalf("%s must be data to %s but '%s'", current, data.expected, ret)
				}
			}
		}
	})
}
