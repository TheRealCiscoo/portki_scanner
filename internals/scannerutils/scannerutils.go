package scannerutils

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type PortsRange struct {
	Initial int
	Final   int
}

type TargetToScan struct {
	Target 			string
	PortsRange  PortsRange
	Proto 			string
}

func ScanPorts(target TargetToScan, wg *sync.WaitGroup) {
	for tport := target.PortsRange.Initial; tport <= target.PortsRange.Final; tport++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			thost := target.Target + ":" + strconv.Itoa(tport)
			conn, err := net.Dial(target.Proto, thost)
			if err == nil {
				fmt.Printf("[+] Open port %d on target %s\n", tport, target.Target)
				conn.Close()
			}
		}()
	}
}