package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"

	ua "github.com/goonya/open62541-go/open62541"
)

var (
	lang, _ = os.LookupEnv("LANG")
	locale  = strings.Split(lang, ".")[0]
)

func manuallyDefinePump(server *ua.Server) {
	var pumpID ua.NodeId /* get the nodeid assigned by the server */
	var oAttr ua.ObjectAttributes
	oAttr.DisplayName = ua.LOCALIZEDTEXT([]byte(locale), []byte("Pump (Manual)"))
	//oAttr.PassRef()
	ua.ServerAddObjectNode(
		server, ua.NodeId{},
		ua.NODEIDNUMERIC(0, ua.NS0IDOBJECTSFOLDER),
		ua.NODEIDNUMERIC(0, ua.NS0IDORGANIZES),
		ua.QUALIFIEDNAME(1, []byte("Pump (Manual)")),
		ua.NODEIDNUMERIC(0, ua.NS0IDBASEOBJECTTYPE),
		oAttr, nil, &pumpID,
	)

	var mnAttr ua.VariableAttributes
	// var manufacturerName ua.String = ua.STRING([]byte("Pump King Ltd."))
	// ua.VariantSetScalar(&mnAttr.Value, unsafe.Pointer(&manufacturerName), &ua.TYPES[ua.TYPESSTRING])
	mnAttr.DisplayName = ua.LOCALIZEDTEXT([]byte(locale), []byte("ManufacturerName"))
	ua.ServerAddVariableNode(server, ua.NodeId{}, pumpID,
		ua.NODEIDNUMERIC(0, ua.NS0IDHASCOMPONENT),
		ua.QUALIFIEDNAME(1, []byte("ManufacturerName")),
		ua.NODEIDNUMERIC(0, ua.NS0IDBASEDATAVARIABLETYPE), mnAttr, nil, nil,
	)

	var modelAttr ua.VariableAttributes
	// var modelName ua.String = ua.STRING([]byte("Mega Pump 3000"))
	// ua.VariantSetScalar(&modelAttr.Value, unsafe.Pointer(&modelName), &ua.TYPES[ua.TYPESSTRING])
	modelAttr.DisplayName = ua.LOCALIZEDTEXT([]byte(locale), []byte("ModelName"))
	ua.ServerAddVariableNode(server, ua.NodeId{}, pumpID,
		ua.NODEIDNUMERIC(0, ua.NS0IDHASCOMPONENT),
		ua.QUALIFIEDNAME(1, []byte("ModelName")),
		ua.NODEIDNUMERIC(0, ua.NS0IDBASEDATAVARIABLETYPE), modelAttr, nil, nil)

	var statusAttr ua.VariableAttributes
	// var status ua.Boolean = true
	// ua.VariantSetScalar(&statusAttr.Value, unsafe.Pointer(&status), &ua.TYPES[ua.TYPESBOOLEAN])
	statusAttr.DisplayName = ua.LOCALIZEDTEXT([]byte(locale), []byte("Status"))
	ua.ServerAddVariableNode(server, ua.NodeId{}, pumpID,
		ua.NODEIDNUMERIC(0, ua.NS0IDHASCOMPONENT),
		ua.QUALIFIEDNAME(1, []byte("Status")),
		ua.NODEIDNUMERIC(0, ua.NS0IDBASEDATAVARIABLETYPE), statusAttr, nil, nil)

	var rpmAttr ua.VariableAttributes
	// var rpm ua.Double = 50.0
	//ua.VariantSetScalar(&rpmAttr.Value, unsafe.Pointer(&rpm), &ua.TYPES[ua.TYPESDOUBLE])
	rpmAttr.DisplayName = ua.LOCALIZEDTEXT([]byte(locale), []byte("MotorRPM"))
	ua.ServerAddVariableNode(server, ua.NodeId{}, pumpID,
		ua.NODEIDNUMERIC(0, ua.NS0IDHASCOMPONENT),
		ua.QUALIFIEDNAME(1, []byte("MotorRPMs")),
		ua.NODEIDNUMERIC(0, ua.NS0IDBASEDATAVARIABLETYPE), rpmAttr, nil, nil)
}

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

	manuallyDefinePump(server)

	status := ua.ServerRun(server, &running)
	ua.ServerDelete(server)
	os.Exit(int(status))
}
