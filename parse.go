package gooob

import (
	"strings"
)

type infoResponse struct {
	data map[string]string
}

type statusResponse struct {
	info    InfoResponse
	players []Player
}

type rconResponse struct {
	message string
}

func (conn *OOBConnection) parseInfo(response []byte) (result *infoResponse, err error) {
	return conn.parseInfoBytes(response, "info", true, "\\\\")
}

func (conn *OOBConnection) parseGetInfo(response []byte) (result *infoResponse, err error) {
	return conn.parseInfoBytes(response, "infoResponse", false, "\\")
}

func (conn *OOBConnection) parseGetStatus(response []byte) (result *statusResponse, err error) {
	lines := strings.Split(string(response), "\n")
	info, err := conn.parseInfoInternal(lines[1], "statusResponse", false, "\\")
	if err != nil {
		return
	}

	players, err := conn.parsePlayers(lines[2:len(lines)])
	return &statusResponse{info, players}, err
}

func (conn *OOBConnection) parseRcon(response []byte) (result *rconResponse, err error) {
	return &rconResponse{string(response)}, nil
}

func (conn *OOBConnection) parseInfoBytes(response []byte, name string, ignoreLastLine bool, separator string) (result *infoResponse, err error) {
	return conn.parseInfoInternal(string(response), name, ignoreLastLine, separator)
}

func (conn *OOBConnection) parseInfoInternal(response string, name string, ignoreLastLine bool, separator string) (result *infoResponse, err error) {
	str := strings.TrimLeft(response, name+"\n")

	lines := strings.Split(str, separator)

	m := make(map[string]string)

	lastLine := 0
	if ignoreLastLine {
		lastLine = len(lines) - 2
	} else {
		lastLine = len(lines) - 1
	}

	for index := 1; index < lastLine; index = index + 2 {
		key := lines[index]
		value := lines[index+1]
		m[key] = value
	}

	return &infoResponse{m}, nil
}

func (conn *OOBConnection) parsePlayers(lines []string) (players []Player, err error) {
	players = make([]Player, len(lines), len(lines))
	// TODO: Fill.
	return players, nil
}

func (r *infoResponse) Data() map[string]string {
	return r.data
}

func (r *statusResponse) Data() map[string]string {
	return r.info.Data()
}

func (r *statusResponse) Players() []Player {
	return r.players
}

func (r *rconResponse) String() string {
	return r.message
}
