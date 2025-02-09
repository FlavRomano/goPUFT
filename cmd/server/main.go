package main

import (
	"net"
	"time"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

const (
	UDP         string = "udp"
	SERVER_HOST string = "HOST"
	SERVER_PORT string = "PORT"
)

func readFromClient(conn *net.UDPConn) (string, *net.UDPAddr) {
	buffer := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(buffer)

	logger.Infof("Reading from %s", remoteAddr.IP)
	if err != nil {
		logger.Errorf("Unexpected error on read: %s", err.Error())
		return "EOF", nil
	}

	return string(buffer)[0:n], remoteAddr
}

func sendToClient(conn *net.UDPConn, addr *net.UDPAddr, mess string) {
	logger.Info("Sending message to client...")
	(*conn).WriteToUDP([]byte(mess), addr)
}

func connInit(host string, port int) *net.UDPConn {
	addr := net.UDPAddr{
		IP:   net.ParseIP(host),
		Port: port,
	}
	server, err := net.ListenUDP(UDP, &addr)
	if err != nil {
		logger.Errorf("Can't open server connection at %s:%s", addr.IP, addr.Port)
		panic(1)
	}

	return server
}

func main() {
	baseLogger, _ := zap.NewProduction()
	logger = baseLogger.Sugar()
	defer logger.Sync()

	host := "127.0.0.1"
	port := 8080

	logger.Infof("Server started, listening on port %d", port)

	server := connInit(host, port)
	defer server.Close()
	for {
		clientMess, remoteAddr := readFromClient(server)
		logger.Infof("received %s", clientMess)
		time.Sleep(1e9 * 5)
		sendToClient(server, remoteAddr, clientMess)
	}
}
