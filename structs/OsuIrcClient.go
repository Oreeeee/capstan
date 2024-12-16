package structs

import (
	"capstan/constants"
	"capstan/irc"
	"fmt"
	"net"
)

type OsuIrcClient struct {
	Conn     net.Conn
	Nickname string
	Hostname string
}

func (c *OsuIrcClient) WriteJoinAck(channelName string) {
	c.Conn.Write([]byte(fmt.Sprintf(":%s!%s@%s %s :%s", c.Nickname, constants.OSU_USERNAME, c.Hostname, irc.IRC_JOIN, channelName)))
}
