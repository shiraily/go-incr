package increment

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type VersionNumber int

const (
	Unknown VersionNumber = iota
	Major
	Minor
	Patch
	Build

	// TODO not enough expression for pre release & build metadata
	// TODO support Windows
	regExpPattern = `([0-9]+)\.([0-9]+)\.([0-9]+)(\.[0-9]+)?(\-[a-zA-Z0-9]+)?(.*)(\n)?`
)

func Increment(version string, target VersionNumber, preserveSuffix bool) (string, error) {
	rep := regexp.MustCompile(regExpPattern)
	result := rep.FindAllStringSubmatch(version, -1)
	if len(result) != 1 {
		return "", fmt.Errorf("not a valid semantic expected='%s'", version)
	}
	numbers := result[0]
	if numbers[5] == "" && numbers[6] != "" {
		return "", fmt.Errorf("not a valid semantic expected='%s'", version)
	}

	// TODO simplify
	lastVersion := Build
	if numbers[Build] == "" {
		if target == Build {
			return "", fmt.Errorf("no build expected: %s", version)
		}
		lastVersion = Patch
	}
	if target == Unknown {
		target = lastVersion
	}

	var startIndex int
	if target == Build {
		startIndex = 1
	}
	num, _ := strconv.Atoi(numbers[target][startIndex:])
	numbers[target] = strconv.Itoa(num + 1)
	for i := target + 1; i < lastVersion+1; i++ {
		numbers[i] = "0"
	}
	versionString := strings.Join(numbers[1:lastVersion+1], ".")
	if preserveSuffix {
		versionString += strings.Join(numbers[lastVersion+1:7], "")
	}
	return versionString + numbers[7], nil
}
