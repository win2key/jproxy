//go:build windows
// +build windows

package main

import (
	"github.com/win2key/jproxy"
)

func main() {
	jproxy.RunWindowsProxy("0.0.0.0:20202")
}
