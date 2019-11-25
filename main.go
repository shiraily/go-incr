package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/shiraily/go-incr/increment"
)

func main() {
	type argVersion struct {
		name  string
		value bool
	}
	argVersions := []argVersion{
		{name: "major"},
		{name: "minor"},
		{name: "patch"},
		{name: "build"},
	}
	for i, argVersion := range argVersions {
		flag.BoolVar(
			&argVersions[i].value,
			argVersion.name,
			false,
			fmt.Sprintf("increment %s version", argVersion.name),
		)
	}
	preserveSuffix := flag.Bool("suffix", false, "preserve suffix (pre release version / build metadata)")
	flag.Parse()
	args := flag.Args()
	argVersions = append([]argVersion{{}}, argVersions...)

	var filePath string
	if len(args) == 0 {
		filePath = "VERSION"
	} else {
		filePath = args[0]
	}

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	version := string(buf)
	fmt.Println("before:", strings.TrimSpace(version))

	fi, err := os.Stat(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var target increment.VersionNumber
	for i, argVersion := range argVersions {
		if argVersion.value {
			target = increment.VersionNumber(i)
			break
		}
	}
	incremented, err := increment.Increment(version, target, *preserveSuffix)
	if err != nil {
		log.Fatal(err)
	}
	// TODO preserve line feed
	if err := ioutil.WriteFile(filePath, []byte(incremented), fi.Mode()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("after :", strings.TrimSpace(incremented))
	return
}
