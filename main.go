package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
	// TODO detailed usage for git option
	git := flag.Bool("git", false, "git commit -m 'Update version'")
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
	if err := ioutil.WriteFile(filePath, []byte(incremented), fi.Mode()); err != nil {
		log.Fatal(err)
	}
	if !*git {
		fmt.Println("before:", strings.TrimSpace(version))
		fmt.Println("after :", strings.TrimSpace(incremented))
		return
	}

	// TODO use go-git?
	if err := isOnlyVersionFileStaged(filePath); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("git", "commit", "-m", "Update version", filePath)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("git commit succeeded")
	// TODO can preserve color
	cmd = exec.Command("git", "diff", "HEAD^")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(output))
	return
}

func isOnlyVersionFileStaged(targetFile string) error {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	outputByte, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("git diff: %s", err)
	}
	output := string(outputByte)
	lines := strings.Split(output, "\n")
	if len(lines) > 2 || (lines[0] != "" && lines[0] != targetFile) {
		return fmt.Errorf("only target version file can be staged:\n%s", output)
	}
	return nil
}
