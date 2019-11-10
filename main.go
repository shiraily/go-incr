package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/shiraily/go-incr/increment"
)

func main() {
	argVersions := []struct {
		name  string
		value bool
	}{
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
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		return
	}
	filePath := args[0]
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	version := string(buf)
	fmt.Println("before:", version)

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
	incremented, err := increment.Increment(version, target)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filePath, []byte(incremented), fi.Mode()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("after :", incremented)
	return
}
