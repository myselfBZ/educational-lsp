package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
)


func EncodeMessage(a any) string{
    msg, err := json.Marshal(a)
    if err != nil{
        log.Fatal("error: ", err)
    }
    return fmt.Sprintf("Content-Length %d \r\n\r\n%s", len(msg), msg)
}

func DecodeMessage(msg []byte) (int, error) {
    // We will get back to this
    header, _, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
    if !found{
        return 0, errors.New("content not found")
    }
    contentSizeBytes := header[len("Content-Length: "):]
    contentSize, err:= strconv.Atoi(string(contentSizeBytes))
    if err != nil{
        return 0, errors.New("invalid content size")
    }
    return contentSize, nil
}

