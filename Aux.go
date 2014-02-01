package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CountLines(filepath string) int {
	b, e := ioutil.ReadFile(filepath)
	if e != nil {
		log.Fatalf("Cannot get line count for %s", filepath)
	}
	return len(strings.Split(string(b), "\n"))
}

func GetHostName() string {
	n, e := os.Hostname()
	if e != nil {
		n = "unknown"
	}
	return n
}
