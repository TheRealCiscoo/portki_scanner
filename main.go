package main

import (
	"portki_scanner/internals/paramsutils"
	"portki_scanner/internals/scannerutils"
	"sync"
	"time"
)


func main() {

	validParams := []string{"-t", "-r"}
	target, err := paramsutils.GetParams(validParams)
	timeout := 1 // Timeout in seconds for each connection

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	scannerutils.ScanPorts(target, time.Duration(timeout), &wg)

}
