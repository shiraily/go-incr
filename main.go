package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}
	fmt.Println(args[0])

	buf, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
	return
}
