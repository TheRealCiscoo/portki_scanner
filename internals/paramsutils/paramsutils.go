package paramsutils

import (
	"errors"
	"os"
	"portki_scanner/internals/scannerutils"
	"strconv"
	"strings"
)


func GetParams(validParams []string) (scannerutils.TargetToScan, error) {

	params := os.Args
	target := scannerutils.TargetToScan{Proto: "tcp"}
	var includedParams []string

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
						return target, errors.New("Invalid value for port range param (-r): <start:MinPort>-<end:MaxPort>")
					}

					minPort, errMinPort := strconv.Atoi(portRange[0])
					if errMinPort != nil {
						return target, errors.New("Invalid value for port range param (-r): port != int")
					}
					maxPort, errMaxPort := strconv.Atoi(portRange[1])
					if errMaxPort != nil {
						return target, errors.New("Invalid value for port range param (-r): port != int")
					}

					if minPort > maxPort {
						return target, errors.New("Invalid value for port range param (-r): <start:MinPort>-<end:MaxPort>")
					}

					target.PortsRange.Initial = minPort
					target.PortsRange.Final = maxPort
				}
			}
		}
	}

	if len(includedParams) < len(validParams) {
		return target, errors.New("Some params are missing: (-t | -r)")
	}

	return target, nil

}
