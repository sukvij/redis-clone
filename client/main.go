package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to Redis TCP server on localhost:6379")

	reader := bufio.NewReader(conn)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // 1. Waits here until the user types a line and presses Enter
		input := scanner.Text()

		// 2. Append "\n" if your server reads line-by-line (e.g., using ReadString('\n'))
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("send problem:", err)
			return
		}

		res, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("receive problem:", err)
			return
		}
		fmt.Println(res)
	}
}
