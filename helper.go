package gooob

import (
	"errors"
	"github.com/dchest/uniuri"
)

func (conn *OOBConnection) validateAndRemovePrefix(response []byte) (result []byte, err error) {
	length := len(response)

	if length < 4 {
		return response, errors.New("response too short")
	}

	for i := 0; i < 4; i++ {
		if response[i] != 0xFF {
			return response, errors.New("datagram format invalid")
		}
	}

	result = response[4:length]
	return result, nil
}

func (conn *OOBConnection) newChallenge() string {
	return uniuri.NewLen(64)
}
