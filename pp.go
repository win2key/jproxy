//go:build !windows
// +build !windows

package main

func main() {
	runProxy("0.0.0.0:20202")
}
