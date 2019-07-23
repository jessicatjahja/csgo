package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"sync"
)

func main() {
	operator := New()
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go operator.handleConnection(conn)
	}
}

// Operator does things
type Operator struct {
	mutex sync.Mutex
	Cache map[string]string
}

// New is an Operator constructor
func New() *Operator {
	operator := new(Operator)
	operator.Cache = make(map[string]string)

	return operator
}

func (op *Operator) handleConnection(conn net.Conn) {
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
		result := op.Cache[value]
		conn.Write([]byte(result + "\n"))
		break
	case "set":
		setKey := strings.Split(message, " ")[1]
		setValue := strings.Split(message, " ")[2]
		op.Cache[setKey] = setValue

		conn.Write([]byte("200\n"))
		break
	default:
		conn.Write([]byte("nothing ever happens\n"))
		break
	}

	op.handleConnection(conn)
}
