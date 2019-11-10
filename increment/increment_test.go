package increment

import (
	"errors"
	"testing"
)

type version struct {
	version string
	target  versionNumber
	err     error
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
		for current, expected := range testData {
			ret, err := Increment(current, expected.target)
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
