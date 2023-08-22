"""
Software Name: FSD-Bot-Socket
Version: 1.0.0
Author: XLiao
Copyright (c) 2023 by XLiao ❤
License: This software is licensed under the MIT License.
"""
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// FsdSocketClient represents the FSD socket client
type FsdSocketClient struct {
	hostName string
	port     int
	userCall string
	userName string
	userPwd  string
}

func main() {
	hostName := "" // Host name or IP
	port := 6809
	userCall := "" // User name
	userName := "" // Position name
	userPwd := "" // User password

	client := NewFsdSocketClient(hostName, port, userCall, userName, userPwd)
	client.FsdBot()
}

// NewFsdSocketClient creates a new instance of FsdSocketClient
func NewFsdSocketClient(hostName string, port int, userCall string, userName string, userPwd string) *FsdSocketClient {
	return &FsdSocketClient{
		hostName: hostName,
		port:     port,
		userCall: userCall,
		userName: userName,
		userPwd:  userPwd,
	}
}

// FsdBot handles communication with the FSD server
func (c *FsdSocketClient) FsdBot() {
	fmt.Println("Software Name: FSD-Bot-Socket")
	fmt.Println("Version: 1.0.0")
	fmt.Println("Author: XLiao")
	fmt.Println("License: This software is licensed under the MIT License.")
	fmt.Println("If customization is needed, please contact XLiao(QQ: 2456666787)")

	// Create a socket and connect to the FSD server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.hostName, c.port))
	if err != nil {
		fmt.Printf("Failed to connect to the server: %s\n", err)
		return
	}
	defer conn.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s Successfully established a connection with the server!\n", timestamp)

	// Write login message
	loginMessage := fmt.Sprintf("#AA%s:SERVER:%s:%s:%s:12:9:1:0:30.872780:104.391390:0\n%%%s:21600:12:20:2:30.872780:104.391390:0\n$SXC%s:SERVER:ATC:%s\n$SXC%s:SERVER:CAPS\n$SXC%s:SERVER:127.0.0.1\n",
		c.userName, c.userName, c.userCall, c.userPwd, c.userName, c.userName, c.userName, c.userName)
	_, err = conn.Write([]byte(loginMessage))
	if err != nil {
		fmt.Printf("Failed to send login message: %s\n", err)
		return
	}

	// Read login response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Printf("Failed to read login response: %s\n", err)
		return
	}
	response := string(buffer[:n])

	responseLines := strings.Split(response, "\n")
	if strings.HasPrefix(responseLines[0], "#TMserver") {
		// If the server response starts with "#TMserver", login is successful

		for {
			// Enter an infinite loop to continuously receive messages from the server
			n, err = conn.Read(buffer)
			if err != nil {
				fmt.Printf("Error while receiving message: %s\n", err)
				return
			}
			message := string(buffer[:n])
			if message != "" {
				// Process the received message
				c.processMessage(message)
			}
		}
	} else {
		// If the server response doesn't start with "#TMserver", login failed
		timestamp = time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("%s Login failed! Server response data: %s\n", timestamp, responseLines[len(responseLines)-1])
	}
}

// processMessage processes the received message from the server
func (c *FsdSocketClient) processMessage(message string) {
	// Process the message based on its type
	if strings.HasPrefix(message, "#AA") {
		// ATC is online
		// Implement the functionality needed for ATC online here
	} else if strings.HasPrefix(message, "#DA") {
		// ATC is offline
		// Implement the functionality needed for ATC offline here
	} else if strings.HasPrefix(message, "#AP") {
		// Pilot is online
		// Implement the functionality needed for user online here
	} else if strings.HasPrefix(message, "#DP") {
		// Pilot is offline
		// Implement the functionality needed for user offline here
	}
}
