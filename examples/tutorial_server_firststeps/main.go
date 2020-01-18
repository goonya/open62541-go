package main

import (
	"os"
	"os/signal"
	"syscall"

	ua "github.com/goonya/open62541-go/open62541"
)

func main() {
	var running ua.Boolean = true
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChannel
		running = false
	}()

	server := ua.ServerNew()
	ua.ServerConfigSetDefault(ua.ServerGetConfig(server))
	status := ua.ServerRun(server, &running)
	ua.ServerDelete(server)
	os.Exit(int(status))
}
