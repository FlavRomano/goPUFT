package main

import (
	"fmt"
	"net"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

const (
	UDP         string = "udp"
	SERVER_HOST string = "HOST"
	SERVER_PORT string = "PORT"
)

func echo(conn *net.Conn, str string) {
	length := len(str)
	n, err := (*conn).Write([]byte(str))
	if err != nil || length != n {
		logger.Errorf("Error sending %d bytes to server", length)
		return
	}
	logger.Info("echo ", str)
	logger.Infof("Client sent %d bytes", n)
}

func readFromServer(conn *net.Conn) (string, int) {
	buffer := make([]byte, 1024)
	logger.Info("Reading from server...")
	n, err := (*conn).Read(buffer)
	if err != nil {
		logger.Error("Error receving message from server")
	}
	return string(buffer)[0:n], n
}

func connInit(serverAddress string) net.Conn {
	conn, err := net.Dial(UDP, serverAddress)
	if err != nil {
		logger.Errorf("Client can't open connection with server at address %s", serverAddress)
		panic(1)
	}
	logger.Infof("Client connected with server at %s", serverAddress)
	return conn
}

func main() {
	baseLogger, _ := zap.NewProduction()
	logger = baseLogger.Sugar()
	defer logger.Sync()

	host := "127.0.0.1"
	port := "8080"
	serverAddress := fmt.Sprintf("%s:%s", host, port)

	conn := connInit(serverAddress)
	defer conn.Close()

	for {
		echo(&conn, file)
		message, n := readFromServer(&conn)
		logger.Infof("%d:%s", n, message)
	}

	logger.Info("Client closed, bye bye")
}
