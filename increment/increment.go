package increment

import (
	"fmt"
	"regexp"
	"strconv"
)

const regExpPattern = `([0-9]+)\.([0-9]+)\.([0-9]+)(\.[0-9]+)?(\-[a-zA-Z0-9]+)?(.*)`

func Increment(version string) (string, error) {
	rep := regexp.MustCompile(regExpPattern)
	result := rep.FindAllStringSubmatch(version, -1)
	if len(result) != 1 {
		return "", fmt.Errorf("not a valid semantic version='%s'", version)
	}
	numbers := result[0][1:]
	if numbers[5] != "" {
		return "", fmt.Errorf("not a valid semantic version='%s'", version)
	}
	if numbers[3] == "" {
		patch, _ := strconv.Atoi(numbers[2])
		return fmt.Sprintf("%s.%s.%d", numbers[0], numbers[1], patch+1), nil
	} else {
		build, _ := strconv.Atoi(numbers[3])
		return fmt.Sprintf("%s.%s.%s.%d", numbers[0], numbers[1], numbers[2], build+1), nil
	}

}
