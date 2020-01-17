package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/goonya/open62541-go/open62541"
)

func main() {
	var running bool = true
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChannel
		running = false
	}()

	server := open62541.ServerNew()
	open62541.ServerConfigSetDefault(open62541.ServerGetConfig(server))
	status := open62541.ServerRun(server, &running)
	open62541.ServerDelete(server)
	os.Exit(int(status))
}
