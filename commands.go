package stupidstuff

import (
	"errors"
	"log"
	"strings"
)

type CacheCommand string

const (
	CMDSet CacheCommand = "SET"
	CMDGet CacheCommand = "GET"
)

type Message struct {
	Cmd   CacheCommand
	Key   []byte
	Value []byte
}

func parseMessage(rawMsg []byte) (*Message, error) {
	rawMsgStr := string(rawMsg)
	parts := strings.Split(rawMsgStr, " ")
	if len(parts) < 2 {
		return nil, errors.New("invalid protocol format")

	}

	cmd := CacheCommand(parts[0])
	msg := &Message{
		Cmd: cmd,
		Key: []byte(parts[1]),
	}

	switch cmd {
	case CMDSet:
		if len(parts) != 3 {
			return nil, errors.New("not parseable")
		}

		msg.Value = []byte(parts[2])
		return msg, nil
	case CMDGet:
		if len(parts) != 2 {
			return nil, errors.New("not parseable")
		}

		return msg, nil
	default:
		log.Printf("[server][handleParseCommand] command is not valid")
		return nil, errors.New("cmd is not available")
	}

}
