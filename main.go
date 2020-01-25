package main

import (
	"fmt"
)

var (
	// BuildTime gets populated during the build proces
	BuildTime = ""

	//Version gets populated during the build process
	Version = ""
)


func main() {
	fmt.Printf("Current version is: %s and buildtime is: %s\n", Version, BuildTime)
}
