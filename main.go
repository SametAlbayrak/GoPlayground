package main

import (
	"fmt"
	"os"
)

func main() {

	//Getenv examples
	uName := os.Getenv("USER")
	uDomain := os.Getenv("USERDOMAIN")
	processorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	goRootPAth := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println(uName)
	fmt.Println(uDomain)
	fmt.Println(processorArchitecture)
	fmt.Println(goRootPAth)
	fmt.Println(goPath)
	fmt.Println(homePath)

}
