package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"sync"
)

func handleClient(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		conn.Write([]byte(string(line) + " -- hello bro\n"))
	}
}

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("tcp connection failed:", err)
		return
	}
	defer l.Close()
	fmt.Println("Server listening on :6379")
	wg := &sync.WaitGroup{}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("connection accept failed:", err)
			continue
		}
		wg.Add(1)
		go handleClient(conn, wg)
	}
	wg.Wait()
}
