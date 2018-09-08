package main

import (
	"github.com/goSwap/scripts/python"
	"log"
)

func main(){
	libs, err := scripts.GetLibraries()
	if err != nil{
		log.Fatal(err)
	}

	pyLibInstall := "pip install"
	for _, lib := range libs{
		pyLibInstall += " " + lib
	}

	print(pyLibInstall)
}