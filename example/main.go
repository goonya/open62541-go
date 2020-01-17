package main

import (
	"fmt"

	"github.com/goonya/open62541-go/open62541"
)

func main() {
	server := open62541.ServerNew()
	//open62541.ServerConfigSetDefault(open62541.ServerGetConfig(server))
	fmt.Println(server)
}
