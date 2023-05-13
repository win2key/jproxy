//go:build !windows
// +build !windows

package main

import (
	"log"
	"os"

	"github.com/win2key/jproxy"
)

func main() {
	address, err := os.LookupEnv("JPROXY") // example: JPROXY=:20202
	if !err {
		log.Fatal("JPROXY environment variable not set")
		os.Exit(1)
	}
	jproxy.RunProxy(address)
}
