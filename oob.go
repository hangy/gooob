package gooob

import (
	"net"
)

type OOB interface {
	Send(cmd string) (response string, err error)
	Info(version int, flags string) (response InfoResponse, err error)
	GetInfo() (response InfoResponse, err error)
	GetStatus() (response StatusResponse, err error)
	Rcon(cmd string) (response RconResponse, err error)
}

type InfoResponse interface {
	Data() map[string]string
}

type StatusResponse interface {
	Data() map[string]string
	Players() []Player
}

type RconResponse interface {
	String() string
}

type Player struct {
	Score      int64
	Ping       int64
	Name       string
	TeamNumber int64
}

func New(server *net.UDPAddr, password string) (OOB, error) {
	return newConnection(server, password)
}
