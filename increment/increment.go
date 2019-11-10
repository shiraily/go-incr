package increment

import (
	"fmt"
	"regexp"
	"strconv"
)

const regExpPattern = "([0-9]+).([0-9]+).([0-9]+)"

func Increment(version string) string {
	rep := regexp.MustCompile(regExpPattern)
	minor, _ :=strconv.Atoi(rep.ReplaceAllString(version, "$3"))
	return rep.ReplaceAllString(version, fmt.Sprintf("$1.$2.%d",minor+1))
}