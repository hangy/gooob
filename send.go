package gooob

import (
	"net"
)

const (
	maxDatagramSize = 65507
)

func (conn *OOBConnection) sendAndClean(cmd string) (response []byte, err error) {
	buffer, err := conn.send(cmd)
	response, err = conn.validateAndRemovePrefix(buffer)
	return response, err
}

func (conn *OOBConnection) send(cmd string) (response []byte, err error) {
	con, err := net.DialUDP("udp", nil, conn.addr)
	if err != nil {
		return
	}

	defer con.Close()

	req := append(
		[]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x02},
		cmd...)
	con.Write(req)

	response = make([]byte, maxDatagramSize, maxDatagramSize)
	con.Read(response)
	return response, err
}
