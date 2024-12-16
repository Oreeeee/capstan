package structs

import (
	"capstan/constants"
	"capstan/irc"
	"fmt"
	"net"
	"strings"
)

type OsuIrcClient struct {
	Conn     net.Conn
	Nickname string
	Hostname string
}

func (c *OsuIrcClient) RecvLoop() {
	for {
		recvBuf := make([]byte, 1024)
		n, err := c.Conn.Read(recvBuf)
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
			c.Nickname = nick
			fmt.Printf("kutas %02x", c.Nickname)
		case irc.IRC_JOIN:
			channelName := args[0]
			fmt.Println("Join received from", c.Nickname, "to channel", channelName)
			//c.WriteJoinAck(channelName)
		default:
			fmt.Println("Unknown command", msgSplit[0])
		}
	}
}

func (c *OsuIrcClient) SendLoop() {
	// TODO
}

func (c *OsuIrcClient) WriteJoinAck(channelName string) {
	c.Conn.Write([]byte(fmt.Sprintf(":%s!%s@%s %s :%s", c.Nickname, constants.OSU_USERNAME, c.Hostname, irc.IRC_JOIN, channelName)))
}
