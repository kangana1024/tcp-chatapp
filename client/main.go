package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gookit/color"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	fmt.Println("Main Client")
	connection, err := net.Dial("tcp", "localhost:8080")
	logFatal(err)

	defer connection.Close()

	color.Cyan.Println("Enter your username : ")

	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')

	logFatal(err)

	username = strings.Trim(username, " \r\n")

	wellcomeMsg := fmt.Sprintf("Welcome : %s, to the chat, say hi to your friends.\n", username)
	color.Green.Println(wellcomeMsg)
	// read
	go read(connection)
	// write
	write(connection, username)
}
func read(connection net.Conn) {
	for {
		reader := bufio.NewReader(connection)
		message, err := reader.ReadString('\n')

		if err == io.EOF {
			connection.Close()
			fmt.Println("Connection close.")
			os.Exit(0)
		}

		color.Magenta.Println(message)
		color.Red.Println("---------------------------")
	}
}
func write(connection net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		message, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		// username: - message
		message = fmt.Sprintf("%s:- %s \n", username, strings.Trim(message, "\r\n"))

		connection.Write([]byte(message))
	}
}
