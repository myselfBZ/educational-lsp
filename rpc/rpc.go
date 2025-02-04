package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func EncodeMessage(a any) string {
	msg, err := json.Marshal(a)
	if err != nil {
		log.Fatal("error: ", err)
	}
	return fmt.Sprintf("Content-Length %d \r\n\r\n%s", len(msg), msg)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) ([]byte, string, error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return nil, "", errors.New("content not found")
	}
	contentSizeBytes := header[len("Content-Length: "):]
	_, err := strconv.Atoi(string(contentSizeBytes))
	if err != nil {
		return nil, "", errors.New("invalid content size")
	}
	var baseMsg BaseMessage
	if err := json.Unmarshal([]byte(content), &baseMsg); err != nil {
		return nil, "", err
	}
	return content, baseMsg.Method, nil
}
