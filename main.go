package main

import (
	"github.com/ProgrammerToGo/scripts/python"
	"log"
	"fmt"
)

func main(){
	py := scripts.NewPython()
	hasPip := py.ScanPip()
	if hasPip != nil {
		log.Print("Installing PIP...")
		//py.InstallPip()
	}
	err := py.ParseLibraries()
	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("libs: %+v", py.Libraries)
}