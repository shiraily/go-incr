package increment

import (
	"fmt"
	"regexp"
	"strconv"
)

//const regExpPattern = `([0-9]+).([0-9]+).([0-9]+)(.[0-9]+)?`
const regExpPattern = `([0-9]+).([0-9]+).([0-9]+)`

func Increment(version string) string {
	rep := regexp.MustCompile(regExpPattern)
	result := rep.FindAllStringSubmatch(version, -1)
	if len(result) != 1 {
		return ""
	}
	versions := result[0][1:]
	if len(versions) < 3 {
		return ""
	}
	minor, _ := strconv.Atoi(versions[2])
	return fmt.Sprintf("%s.%s.%d", versions[0], versions[1], minor+1)
	return ""
}
