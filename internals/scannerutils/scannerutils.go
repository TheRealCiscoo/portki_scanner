package scannerutils

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
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

func ScanPorts(target TargetToScan, timeout time.Duration, wg *sync.WaitGroup) {
	for tport := target.PortsRange.Initial; tport <= target.PortsRange.Final; tport++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			thost := target.Target + ":" + strconv.Itoa(tport)
			dialer := net.Dialer{Timeout: timeout * time.Second}
			conn, err := dialer.Dial(target.Proto, thost)
			if err == nil {
				fmt.Printf("[+] Open port %d on target %s\n", tport, target.Target)
				conn.Close()
			}
		}()
	}
}