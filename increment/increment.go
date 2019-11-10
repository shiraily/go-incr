package increment

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type versionNumber int

const (
	Major versionNumber = iota
	Minor
	Patch
	Build

	regExpPattern = `([0-9]+)\.([0-9]+)\.([0-9]+)(\.[0-9]+)?(\-[a-zA-Z0-9]+)?(.*)`
)

func Increment(version string, target versionNumber) (string, error) {
	rep := regexp.MustCompile(regExpPattern)
	result := rep.FindAllStringSubmatch(version, -1)
	if len(result) != 1 {
		return "", fmt.Errorf("not a valid semantic version='%s'", version)
	}
	numbers := result[0][1:]
	if numbers[5] != "" {
		return "", fmt.Errorf("not a valid semantic version='%s'", version)
	}

	isTargetBuild := target == Build
	var startIndex int
	if isTargetBuild {
		startIndex = 1
	}
	lastVersion := Build
	if numbers[Build] == "" {
		if isTargetBuild {
			return "", fmt.Errorf("no build version: %s", version)
		}
		lastVersion = Patch
	}

	num, _ := strconv.Atoi(numbers[target][startIndex:])
	numbers[target] = strconv.Itoa(num + 1)
	for i := target + 1; i < lastVersion+1; i++ {
		numbers[i] = "0"
	}
	return strings.Join(numbers[0:lastVersion+1], "."), nil
}
