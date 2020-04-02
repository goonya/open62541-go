package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"unsafe"

	ua "github.com/goonya/open62541-go/open62541"
)

func addVariable(server *ua.Server, value int32) {
	var attr ua.VariableAttributes
	var myInteger int32 = int32(value)
	ua.VariantSetScalar(&attr.Value, unsafe.Pointer(&myInteger), &ua.TYPES[ua.TYPESINT32])
	attr.Description = ua.LOCALIZEDTEXT([]byte("ru-RU"), []byte("the answer"))
	attr.DisplayName = ua.LOCALIZEDTEXT([]byte("ru-RU"), []byte("the answer"))
	//attr.DataType = ua.TYPES[ua.TYPESINT32].TypeId
	attr.AccessLevel = ua.ACCESSLEVELMASKREAD | ua.ACCESSLEVELMASKWRITE

	/* Add the variable node to the information model */
	var myIntegerNodeId = ua.NODEIDSTRING(1, []byte("the.answer"))
	var myIntegerName = ua.QUALIFIEDNAME(1, []byte("the answer"))
	var parentNodeId = ua.NODEIDNUMERIC(0, ua.NS0IDOBJECTSFOLDER)
	var parentReferenceNodeId = ua.NODEIDNUMERIC(0, ua.NS0IDORGANIZES)
	ua.ServerAddVariableNode(server, myIntegerNodeId, parentNodeId,
		parentReferenceNodeId, myIntegerName,
		ua.NODEIDNUMERIC(0, ua.NS0IDBASEDATAVARIABLETYPE), attr, nil, nil)
}

func writeVariable(server *ua.Server, value int32) {
	var myIntegerNodeId = ua.NODEIDSTRING(1, []byte("the.answer"))

	var myInteger int32 = int32(value)
	var myVar ua.Variant
	ua.VariantInit(&myVar)
	ua.VariantSetScalar(&myVar, unsafe.Pointer(&myInteger), &ua.TYPES[ua.TYPESINT32])
	ua.ServerWriteValue(server, myIntegerNodeId, myVar)

	var wv ua.WriteValue
	ua.WriteValueInit(&wv)
	wv.NodeId = myIntegerNodeId
	wv.AttributeId = uint32(ua.ATTRIBUTEIDVALUE)
	wv.Value.SetStatus(uint32(ua.STATUSCODEBADNOTCONNECTED))
	// wv.Value.HasStatus = true
	ua.ServerWrite(server, &wv)

	// // wv.Value.HasStatus = false
	// wv.Value.SetValue(myVar)
	// // wv.Value.HasValue = true
	// ua.ServerWrite(server, &wv)
}

func main() {
	var running bool = true
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChannel
		running = false
	}()

	server := ua.ServerNew()
	ua.ServerConfigSetDefault(ua.ServerGetConfig(server))

	var value int32 = 32

	addVariable(server, value)

	ticker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			value++
			writeVariable(server, value)
		}
	}()
	writeVariable(server, 77)

	status := ua.ServerRun(server, &running)
	ua.ServerDelete(server)
	os.Exit(int(status))
}
