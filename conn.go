package gooob

import (
	"net"
	"strconv"
)

type OOBConnection struct {
	addr     *net.UDPAddr
	password string
}

func newConnection(raddr *net.UDPAddr, password string) (*OOBConnection, error) {
	return &OOBConnection{raddr, password}, nil
}

func (conn *OOBConnection) Send(cmd string) (response string, err error) {
	buffer, err := conn.sendAndClean(cmd)
	return string(buffer), err
}

func (conn *OOBConnection) Info(version int, flags string) (response InfoResponse, err error) {
	cmd := "info " + strconv.Itoa(version) + " " + flags
	buffer, err := conn.sendAndClean(cmd)
	if err != nil {
		return
	}

	return conn.parseInfo(buffer)
}

func (conn *OOBConnection) GetInfo() (response InfoResponse, error error) {
	challenge := conn.newChallenge()
	cmd := "getinfo " + challenge
	buffer, err := conn.sendAndClean(cmd)
	if err != nil {
		return
	}

	return conn.parseGetInfo(buffer)
}

func (conn *OOBConnection) GetStatus() (response StatusResponse, error error) {
	challenge := conn.newChallenge()
	cmd := "getstatus " + challenge
	buffer, err := conn.sendAndClean(cmd)
	if err != nil {
		return
	}

	return conn.parseGetStatus(buffer)
}

func (conn *OOBConnection) Rcon(cmd string) (response RconResponse, err error) {
	cmd = "rcon " + conn.password + cmd
	buffer, err := conn.sendAndClean(cmd)
	if err != nil {
		return
	}

	return conn.parseRcon(buffer)
}
