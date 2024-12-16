package main

import (
	"capstan/irc"
	"capstan/structs"
	"fmt"
	"net"
	"strings"
)

func HandleIRC(osuconn structs.OsuIrcClient) {
	defer osuconn.Conn.Close()
	fmt.Println("Handling IRC connection")
	for {
		recvBuf := make([]byte, 1024)
		n, err := osuconn.Conn.Read(recvBuf)
		if err != nil {
			fmt.Println(err)
		}
		msg := strings.Replace(string(recvBuf[:n]), "\r", "", -1)
		msg = strings.Replace(msg, "\n", "", -1)
		msgSplit := strings.Split(msg, " ")

		command := msgSplit[0]
		args := msgSplit[1:]
		fmt.Printf("Read %d bytes, message: %s\n", n, msg)
		switch command {
		case irc.IRC_NICK:
			nick := args[0]
			fmt.Println("Nick received", nick)
			osuconn.Nickname = nick
			fmt.Printf("kutas %02x", osuconn.Nickname)
		case irc.IRC_JOIN:
			channelName := args[0]
			fmt.Println("Join received from", osuconn.Nickname, "to channel", channelName)
			osuconn.WriteJoinAck(channelName)
		default:
			fmt.Println("Unknown command", msgSplit[0])
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
		osuconn := structs.OsuIrcClient{Conn: conn, Hostname: strings.Split(conn.RemoteAddr().String(), ":")[0]}
		go HandleIRC(osuconn)
	}
}
