package rpc

import (
	"testing"
)

type Message struct{
    Msg string `json:"msg"`
}

func TestEncode(t *testing.T) {
    msg := Message{"hello"}
    expected := "Content-Length 15 \r\n\r\n{\"msg\":\"hello\"}"
    if expected != EncodeMessage(msg){
        t.Fatalf("didn't match: got %s", EncodeMessage(msg))
    }
}
