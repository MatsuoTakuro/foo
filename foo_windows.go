//go:build windows
// +build windows

package foo

import "fmt"

func PlatformPrint() {
	fmt.Println("Running on Windows")
}
