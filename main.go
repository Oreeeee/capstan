package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleIRC(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Handling IRC connection")
	var sentNames bool
	for {
		recvBuf := make([]byte, 1024)
		n, err := conn.Read(recvBuf)
		if err != nil {
			fmt.Println(err)
		}
		msg := strings.Replace(string(recvBuf[:n]), "\n", "", -1)
		fmt.Printf("Read %d bytes, message: %s\n", n, msg)
		if strings.Contains(msg, "NAMES #peppy-osu-r") {
			sentNames = true
		}
		if sentNames {
			for _, str := range []string{"#osu", "#global"} {
				fmt.Println("Jfjfsdjfjsfjsj fjdsfj,diohjkhjla " + str)
				conn.Write([]byte(":peppy-osu!OSU@localhost JOIN :" + str + "\n"))
			}

		}
	}
}

func main() {
	fmt.Println("Hello World")

	l, err := net.Listen("tcp", "127.0.0.1:6667")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Println("Listening on 127.0.0.1:6667")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go HandleIRC(conn)
	}
}
