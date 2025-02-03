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

func TestDecode(t *testing.T){
    msg := []byte("Content-Length: 17\r\n\r\n{\"method\":\"hiii\"}") 
    contentLength := 17
    methodExpected := "hiii"
    cntLngth, method,err := DecodeMessage(msg)
    if err != nil{
        t.Fatalf("couldn't decode: %v", err)
    }
    if len(cntLngth) != contentLength{
        t.Fatalf("expected %d got: %d", contentLength, len(cntLngth))
    }
    if methodExpected != method {
        t.Fatalf("expected %s got : %s", methodExpected, method)
    }
}
