package main

import (
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Set the SSH configuration parameters
	sshConfig := &ssh.ClientConfig{
		User: "Swaroop",
		Auth: []ssh.AuthMethod{
			ssh.Password("cvTOaF06RjdArA"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the remote SSH server
	conn, err := net.Dial("tcp", "tty.sdf.org:22")
	if err != nil {
		fmt.Printf("Failed to connect to remote server: %s", err)
		os.Exit(1)
	}

	// Establish an SSH connection
	sshConn, chans, reqs, err := ssh.NewClientConn(conn, "tty.sdf.org", sshConfig)
	if err != nil {
		fmt.Printf("Failed to establish SSH connection: %s", err)
		os.Exit(1)
	}

	// Create an SSH client using the connection
	sshClient := ssh.NewClient(sshConn, chans, reqs)

	// Run an SSH command on the remote server
	session, err := sshClient.NewSession()
	if err != nil {
		fmt.Printf("Failed to create SSH session: %s", err)
		os.Exit(1)
	}
	defer session.Close()

	output, err := session.CombinedOutput("ls -la")
	if err != nil {
		fmt.Printf("Failed to run command: %s", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
