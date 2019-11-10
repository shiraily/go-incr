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
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	fmt.Println(args[0])

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
	incremented, err := increment.Increment(version, increment.Patch)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filePath, []byte(incremented), fi.Mode()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("after :", incremented)
	return
}
