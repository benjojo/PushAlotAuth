package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func CountLines(filepath string) int {
	b, e := ioutil.ReadFile(filepath)
	if e != nil {
		log.Fatalf("Cannot get line count for %s", filepath)
	}
	return len(strings.Split(string(b), "\n"))
}
