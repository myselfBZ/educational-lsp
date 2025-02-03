package main

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os"
	"strconv"
)

func main(){
    logger := getLogger("/home/boburmirzoalivobjonov/Programming/Projects/LSP/log.txt")
    logger.Println("Hi")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(Split)
    for scanner.Scan() {
        handleMsg(scanner.Text(), logger)   
    }
}


func handleMsg(d any, l *log.Logger) {
    l.Println("got a new message: ", d.(string))    
}

func Split(msg []byte, _ bool) (advance int, token []byte, err error){
    header, content, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
    if !found{
        return 0, nil, errors.New("content not found")
    }
    contentSizeBytes := header[len("Content-Length: "):]
    contentSize, err := strconv.Atoi(string(contentSizeBytes))
    if err != nil{
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
    if err != nil{
        log.Fatal("error: ", err)
    }
    return log.New(file, "[educationallsp]", log.Ldate|log.Lshortfile)
}





