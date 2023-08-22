# FSD-Bot-Socket (Golang Version)

![License](https://img.shields.io/badge/Code-Golang-blue.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

**FSD-Bot-Socket** is a Go program designed to establish communication with an FSD (Flight Simulation Data) server using sockets. It allows you to monitor and interact with various events, such as ATC (Air Traffic Control) controllers and pilots going online and offline, within a flight simulation environment.

This repository contains the Golang version of the FSD-Bot-Socket program. Build and run the program to start monitoring flight simulation activities.

## Features

- **Real-time Monitoring**: The program establishes a connection to the FSD server and monitors activities such as ATC controllers and pilots logging in and out.

- **Interactive Responses**: Based on server messages, you can implement custom functionality when an ATC controller or pilot goes online or offline. This makes it a versatile tool for managing flight simulation sessions.

## Dependencies

Before building the program, make sure you have the required Go modules installed:

- `net`

You can install the required modules using:

```bash
go get -u net
```

## Usage
1. Configure the user settings in the program.
2. Build and run the program: go run fsd_bot_socket.go.
3. The program will establish a connection to the FSD server and start monitoring relevant events.

## Languages

- [Python Version](https://github.com/XLiaovo/FSD-Bot-Socket-Python-Version
)
- [C++ Version](https://github.com/XLiaovo/FSD-Bot-Socket-Cpp-Version
)
- [Go Version](https://github.com/XLiaovo/FSD-Bot-Socket-Golang-Version
)

## License
This project is licensed under the terms of the MIT License. See [LICENSE](LICENSE) for more details.

***

<p align="center">
  Looking forward to your ⭐️ Star!
</p>
