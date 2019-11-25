package increment

import (
	"errors"
	"testing"
)

type version struct {
	expected       string
	target         VersionNumber
	preserveSuffix bool
	err            error
}

func TestIncrement(t *testing.T) {
	t.Helper()
	var (
		err      = errors.New("")
		testData = map[string]version{
			"1.0.0":             {"1.0.1", Patch, false, nil},
			"1.0.0.1":           {"1.0.0.2", Build, false, nil},
			"1.0.1.2":           {"1.0.2.0", Patch, false, nil},
			"0.1.2.3":           {"0.2.0.0", Minor, false, nil},
			"1.2.3.4":           {"2.0.0.0", Major, false, nil},
			"1.2.3.4-aaa":       {"2.0.0.0", Major, false, nil},
			"1.2.3.4-bbb+a.b.c": {"2.0.0.0-bbb+a.b.c", Major, true, nil},
			"1.2.3.4-aaa\n":     {"2.0.0.0\n", Major, false, nil},
			"1.0.0\n":           {"1.0.1\n", Patch, true, nil},
			"2.0.0.1":           {"2.0.0.2", Unknown, false, nil},
			"2.0.1":             {"2.0.2", Unknown, false, nil},
			"1.0.0.0a":          {"", Patch, false, err},
			"1.0.a0":            {"", Patch, false, err},
			"1.0a.0":            {"", Patch, false, err},
			"1.0.0a":            {"", Patch, false, err},
		}
	)

	t.Run("Increment", func(t *testing.T) {
		for current, data := range testData {
			ret, err := Increment(current, data.target, data.preserveSuffix)
			if err != nil {
				if data.err == nil {
					t.Fatalf("%s should be error but returned '%s': %s", current, data.expected, err)
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
