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

	// TODO not enough expression for pre release expected & build metadata
	regExpPattern = `([0-9]+)\.([0-9]+)\.([0-9]+)(\.[0-9]+)?(\-[a-zA-Z0-9]+)?(.*)`
)

func Increment(version string, target VersionNumber) (string, error) {
	rep := regexp.MustCompile(regExpPattern)
	result := rep.FindAllStringSubmatch(version, -1)
	if len(result) != 1 {
		return "", fmt.Errorf("not a valid semantic expected='%s'", version)
	}
	numbers := result[0]
	if numbers[6] != "" {
		return "", fmt.Errorf("not a valid semantic expected='%s'", version)
	}

	isTargetBuild := target == Build
	var startIndex int
	if isTargetBuild {
		startIndex = 1
	}
	lastVersion := Build
	if numbers[Build] == "" {
		if isTargetBuild {
			return "", fmt.Errorf("no build expected: %s", version)
		}
		lastVersion = Patch
	}
	if target == Unknown {
		target = lastVersion
	}

	num, _ := strconv.Atoi(numbers[target][startIndex:])
	numbers[target] = strconv.Itoa(num + 1)
	for i := target + 1; i < lastVersion+1; i++ {
		numbers[i] = "0"
	}
	return strings.Join(numbers[1:lastVersion+1], "."), nil
}
