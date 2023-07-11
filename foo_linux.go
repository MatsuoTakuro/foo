//go:build linux
// +build linux

package foo

import "fmt"

func PlatformPrint() {
	fmt.Println("Running on Linux")
}
