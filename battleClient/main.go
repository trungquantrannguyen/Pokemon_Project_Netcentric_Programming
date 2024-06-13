package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	startTCPClient()
}

func startTCPClient() {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to the TCP server.")

	// Get the user's username
	username := getUsernameFromInput()

	// Send the username to the server
	_, err = conn.Write([]byte(username))
	if err != nil {
		fmt.Println("Error writing to connection:", err)
		return
	}

	// Start a goroutine to read responses from the server
	go readResponsesFromServer(conn)

	// Read user input and send it to the server
	readAndSendPokemons(conn, username)

	// Read user battle inputs and send it to server
	readAndSendBattle(conn, username)

}

func getUsernameFromInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	return strings.TrimSpace(username)
}

func readAndSendPokemons(conn net.Conn, username string) {
	reader := bufio.NewReader(os.Stdin)
	for i := 1; i < 4; i++ {
		fmt.Printf("Enter Pokemon %d: ", i)
		text, _ := reader.ReadString('\n')
		if strings.TrimSpace(text) == "exit" {
			break
		}

		// Append the username and Pokemon number to the input text
		text = fmt.Sprintf("%s %d %s", username, i, strings.TrimSpace(text))

		// Send the message to the server
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}

func readAndSendBattle(conn net.Conn, username string) {
	reader := bufio.NewReader(os.Stdin)
	//buf := make([]byte, 1024)
	for {
		//fmt.Printf("Choose your next move (type in 'normal' or 'special'): ")
		text, _ := reader.ReadString('\n')
		text = fmt.Sprintf("%s %s a a",username, strings.TrimSpace(text))
		// Send the message to the server
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to connection:", err)
			return
		}
	}
}
func readResponsesFromServer(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		response := strings.TrimSpace(string(buf[:n]))
		fmt.Println(response)
		if strings.Contains(response, "The winner is") {
			fmt.Println("Thank you for playing")
			os.Exit(0)
		}
	}
}