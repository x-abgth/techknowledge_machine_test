package main

import (
	"fmt"
	"strings"
)

func parseUptimeOutput(output string) error {
	fields := strings.Fields(output)
	serverName := fields[0]
	numUsers := strings.TrimSuffix(fields[3], ",")
	status := fields[1]

	fmt.Println(serverName, numUsers, status)

	return nil
}
