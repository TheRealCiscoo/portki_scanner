package main

import (
	"portki_scanner/internals/paramsutils"
	"portki_scanner/internals/scannerutils"
	"sync"
)


func main() {

	validParams := []string{"-t", "-r"}
	target, err := paramsutils.GetParams(validParams)

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	scannerutils.ScanPorts(target, &wg)

}
