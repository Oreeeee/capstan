package main

import (
	"capstan/globals"
	"capstan/structs"
	"fmt"
	"net"
	"strings"
)

func HandleIRC(osuconn structs.OsuIrcClient) {
	fmt.Println("Handling IRC connection")
	go osuconn.SendLoop()
	go osuconn.RecvLoop()
}

func main() {
	globals.InitRedisSession()

	l, err := net.Listen("tcp", "0.0.0.0:6667")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Println("Listening on 0.0.0.0:6667")

	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		osuconn := structs.OsuIrcClient{Conn: conn, Hostname: strings.Split(conn.RemoteAddr().String(), ":")[0]}
		go HandleIRC(osuconn)
	}
}
