package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/myseflBZ/lsp/lsp"
	"github.com/myseflBZ/lsp/rpc"
)

func main() {
	logger := getLogger("/home/boburmirzoalivobjonov/Programming/Projects/LSP/log.txt")
	logger.Println("Hi")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(Split)
	for scanner.Scan() {
		content, method, err := rpc.DecodeMessage(scanner.Bytes())
		if err != nil {
			logger.Printf("got an error: %v", err)
			continue
		}
		handleMsg(method, content, logger)
	}
}

func handleMsg(method string, content []byte, l *log.Logger) {
	l.Println("got a new message: ", method)
    switch method{
    case "initialize":
        var req lsp.InitializeRequest
        if err := json.Unmarshal(content, &req); err != nil{
            l.Printf("couldn't parse the request: %s", err)
            return
        }
        l.Printf("init method from: %s Version: %s", req.Params.ClientInfo.Name, req.Params.ClientInfo.Version)   
    }
}

// testing a comment
func Split(msg []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, errors.New("content not found")
	}
	contentSizeBytes := header[len("Content-Length: "):]
	contentSize, err := strconv.Atoi(string(contentSizeBytes))
	if err != nil {
		return 0, nil, errors.New("invalid content size")
	}
	if len(content) < contentSize {
		return 0, nil, nil
	}
	totalLength := len(header) + 4 + contentSize
	return totalLength, msg[:totalLength], nil
}

func getLogger(filename string) *log.Logger {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal("error: ", err)
	}
	return log.New(file, "[educationallsp]", log.Ldate|log.Lshortfile)
}
