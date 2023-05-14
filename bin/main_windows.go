//go:build windows
// +build windows

package main

import (
	"github.com/win2key/jproxy"
)

func main() {
	address, err := os.LookupEnv("JPROXY") // example: JPROXY=0.0.0.0:20202
	if !err {
		log.Fatal("JPROXY environment variable not set, using default 0.0.0.0:20202")
		address = ":20202"
	}
	jproxy.RunWindowsProxy(address)
}
