package util

import "strings"

// ParserDockerPorts Parsing docker port information.
func ParserDockerPorts(portArr []string) map[string]string {
	portMap := make(map[string]string)
	for _, portInfo := range portArr {
		if !strings.Contains(portInfo, "->") {
			continue
		}

		portInfo = strings.ReplaceAll(portInfo, ",", "")
		portLinkArr := strings.Split(portInfo, "->")
		key := strings.ReplaceAll(portLinkArr[0], "0.0.0.0:", "")
		value := strings.ReplaceAll(portLinkArr[1], "/tcp", "")
		portMap[key] = value
	}
	return portMap
}
