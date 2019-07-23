package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Fatal(err)
	}

	cache := make(map[string]string)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn, cache)
	}
}

func handleConnection(conn net.Conn, cache map[string]string) {
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Println(err)
		conn.Close()
		return
	}

	message := string(bufferBytes)

	message = strings.TrimSuffix(message, "\n")

	key := strings.Split(message, " ")[0]

	switch key {
	case "get":
		value := strings.Split(message, " ")[1]
		result := cache[value]
		conn.Write([]byte(result + "\n"))
		break
	case "set":
		setKey := strings.Split(message, " ")[1]
		setValue := strings.Split(message, " ")[2]
		cache[setKey] = setValue

		conn.Write([]byte("200\n"))
		break
	default:
		conn.Write([]byte("nothing ever happens\n"))
		break
	}

	handleConnection(conn, cache)
}
