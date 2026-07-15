package paramsutils

import (
	"errors"
	"fmt"
	"os"
	"portki_scanner/internals/scannerutils"
	"slices"
	"strconv"
	"strings"
)


func GetParams(validParams []string) (scannerutils.TargetToScan, error) {

	params := os.Args
	target := scannerutils.TargetToScan{Proto: "tcp"}
	var includedParams []string

	if slices.Contains(params, "-h") || slices.Contains(params, "--help") {
		fmt.Printf("\nMenu options:\n-t\t\tTarget host to be scanned. (Ex: scanme.nmap.org)\n-r\t\tRange of port to scan in the target host. (Ex: 10-1024)\n\nTool developed by https://github.com/TheRealCiscoo\n\n")
		return target, nil
	}

	for key, osParam := range params {
		if key > 0 { // Avoid the first param (file path)
			if strings.HasPrefix(osParam, "-") {

				includedParams = append(includedParams, osParam)

				value := params[(key +	1)]
				if osParam == "-t" {
					target.Target = value
				}
				if osParam == "-r" {
					portRange := strings.Split(value, "-")
					
					if len(portRange) != 2 {
						return target, errors.New("Invalid value for port range param (-r): <start:MinPort>-<end:MaxPort>\n\nUse -h or --help to see the tool menu options")
					}

					minPort, errMinPort := strconv.Atoi(portRange[0])
					if errMinPort != nil {
						return target, errors.New("Invalid value for port range param (-r): port != int\n\nUse -h or --help to see the tool menu options")
					}
					maxPort, errMaxPort := strconv.Atoi(portRange[1])
					if errMaxPort != nil {
						return target, errors.New("Invalid value for port range param (-r): port != int\n\nUse -h or --help to see the tool menu options")
					}

					if minPort > maxPort || minPort < 0 || maxPort > 65535{
						return target, errors.New("Invalid value for port range param (-r): <start:MinPort( >= 0)>-<end:MaxPort( < 65535)>\n\nUse -h or --help to see the tool menu options")
					}

					target.PortsRange.Initial = minPort
					target.PortsRange.Final = maxPort
				}
			}
		}
	}

	if len(includedParams) < len(validParams) {
		return target, errors.New("Some params are missing: (-t | -r)\n\nUse -h or --help to see the tool menu options")
	}

	return target, nil

}
