package main

import (
	"fmt"
	"net"
	"os"
	"redis-lite/commands"
	"redis-lite/instance"
	"redis-lite/resp"
	"strings"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "6379"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	instance := instance.MakeInstance()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(instance, conn)

		go instance.RunExpirationChecker()
	}
}

func handleRequest(instance *instance.Instance, conn net.Conn) {
	for {
		buffer := make([]byte, 1024)

		_, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			break
		}

		cmp, _ := resp.DeserializeArrayOfBulkStrings(string(buffer))

		command := strings.ToUpper((*cmp)[0])

		var args []string
		if len(*cmp) > 1 {
			args = (*cmp)[1:]
		} else {
			args = make([]string, 0)
		}

		commands.HandleCommand(instance, command, args, conn)
	}
}
