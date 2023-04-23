//go:build !windows
// +build !windows

package jproxy

func RunWindowsProxy(address string) {
	RunProxy(address)
}
