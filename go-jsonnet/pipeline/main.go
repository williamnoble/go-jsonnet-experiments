package main

import (
	"fmt"
	"github.com/google/go-jsonnet"
)

func main() {

	vm := jsonnet.MakeVM()
	vm.ExtVar("env", "dev")
	// vm.EvaluateFile => string
	// vm.EvaluateFileMulti => map of string[string] (pretty cool)

	file, err := vm.EvaluateFile("template.jsonnet")
	if err != nil {
		panic(err)
	}
	fmt.Println("str:", file)

	fileMap, err := vm.EvaluateFileMulti("template.jsonnet")
	if err != nil {
		panic(err)
	}
	fmt.Printf("map: %+v\n", fileMap["config"])

	// only works when a top level object is an array
	//fileStream, err := vm.EvaluateFileStream("template.jsonnet")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("map", fileStream)

	// importedFrom - pretend you looked here
	// importedPath - where you actually fetched this data

	//vm.ImportData()
}
