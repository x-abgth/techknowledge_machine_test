package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

type UptimeData struct {
	Server  string `json:"server"`
	Status  string `json:"status"`
	Users   string `json:"users"`
	Command string `json:"command"`
}

func main() {
	if len(os.Args) != 4 {
		log.Println("Error : Use command like - go run main.go <server-ip> <username> <password>")
		return
	}

	serverIP := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]

	output, err := runUptimeCommand(serverIP, username, password)
	if err != nil {
		log.Println(err)
		return
	}

	// err = parseUptimeOutput(output)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	fmt.Println(output)
}

func runUptimeCommand(serverIP, username, password string) (string, error) {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", serverIP+":22", config)
	if err != nil {
		return "", fmt.Errorf("Failed to connect to server: %v", err)
	}

	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("Failed to create SSH session: %v", err)
	}

	defer session.Close()

	output, err := session.Output("uptime")
	if err != nil {
		return "", fmt.Errorf("Failed to run command: %v", err)
	}

	fields := strings.Fields(string(output))
	// loadAvgFields := strings.Split(fields[9], ",")
	// var loadAverage []float64
	// for _, field := range loadAvgFields {
	// 	load, err := strconv.ParseFloat(field, 64)
	// 	if err == nil {
	// 		loadAverage = append(loadAverage, load)
	// 	}
	// }

	// fmt.Println(fields)
	fmt.Println(fields[48:500])

	return "", nil
}

// go run main.go <server-ip> <username> <password>
