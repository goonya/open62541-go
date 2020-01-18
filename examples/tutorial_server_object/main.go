package main

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"unsafe"

	ua "github.com/goonya/open62541-go/open62541"
)

var (
	lang, _ = os.LookupEnv("LANG")
	locale  = strings.Split(lang, ".")[0]
)

func manuallyDefinePump(server *ua.Server) {
	var pumpId ua.NodeId /* get the nodeid assigned by the server */
	var oAttr ua.ObjectAttributes
	oAttr.DisplayName = ua.LOCALIZEDTEXT([]byte(locale), []byte("Pump (Manual)"))
	//oAttr.PassRef()
	ua.ServerAddObjectNode(
		server, ua.NodeId{},
		ua.NODEIDNUMERIC(0, ua.NS0IDOBJECTSFOLDER),
		ua.NODEIDNUMERIC(0, ua.NS0IDORGANIZES),
		ua.QUALIFIEDNAME(1, []byte("Pump (Manual)")),
		ua.NODEIDNUMERIC(0, ua.NS0IDBASEOBJECTTYPE),
		oAttr, nil, &pumpId,
	)

	var mnAttr ua.VariableAttributes
	var manufacturerName ua.String = ua.STRING([]byte("Pump King Ltd."))
	ua.VariantSetScalar(&mnAttr.Value, unsafe.Pointer(&manufacturerName), &ua.TYPES[ua.TYPESSTRING])
	mnAttr.DisplayName = ua.LOCALIZEDTEXT([]byte(locale), []byte("ManufacturerName"))
	ua.ServerAddVariableNode(server, ua.NodeId{}, pumpId,
		ua.NODEIDNUMERIC(0, ua.NS0IDHASCOMPONENT),
		ua.QUALIFIEDNAME(1, []byte("ManufacturerName")),
		ua.NODEIDNUMERIC(0, ua.NS0IDBASEDATAVARIABLETYPE), mnAttr, nil, nil,
	)

	// ua.VariableAttributes modelAttr = ua.VariableAttributes_default;
	// ua.String modelName = ua.STRING("Mega Pump 3000");
	// ua.Variant_setScalar(&modelAttr.value, &modelName, &ua.TYPES[ua.TYPES_STRING]);
	// modelAttr.displayName = ua.LOCALIZEDTEXT([]byte(locale), "ModelName");
	// ua.Server_addVariableNode(server, ua.NodeId{}, pumpId,
	//                           ua.NODEID_NUMERIC(0, ua.NS0ID_HASCOMPONENT),
	//                           ua.QUALIFIEDNAME(1, "ModelName"),
	//                           ua.NODEID_NUMERIC(0, ua.NS0ID_BASEDATAVARIABLETYPE), modelAttr, nil, nil);

	// ua.VariableAttributes statusAttr = ua.VariableAttributes_default;
	// ua.Boolean status = true;
	// ua.Variant_setScalar(&statusAttr.value, &status, &ua.TYPES[ua.TYPES_BOOLEAN]);
	// statusAttr.displayName = ua.LOCALIZEDTEXT([]byte(locale), "Status");
	// ua.Server_addVariableNode(server, ua.NodeId{}, pumpId,
	//                           ua.NODEID_NUMERIC(0, ua.NS0ID_HASCOMPONENT),
	//                           ua.QUALIFIEDNAME(1, "Status"),
	//                           ua.NODEID_NUMERIC(0, ua.NS0ID_BASEDATAVARIABLETYPE), statusAttr, nil, nil);

	// ua.VariableAttributes rpmAttr = ua.VariableAttributes_default;
	// ua.Double rpm = 50.0;
	// ua.Variant_setScalar(&rpmAttr.value, &rpm, &ua.TYPES[ua.TYPES_DOUBLE]);
	// rpmAttr.displayName = ua.LOCALIZEDTEXT([]byte(locale), "MotorRPM");
	// ua.Server_addVariableNode(server, ua.NodeId{}, pumpId,
	//                           ua.NODEID_NUMERIC(0, ua.NS0ID_HASCOMPONENT),
	//                           ua.QUALIFIEDNAME(1, "MotorRPMs"),
	//                           ua.NODEID_NUMERIC(0, ua.NS0ID_BASEDATAVARIABLETYPE), rpmAttr, nil, nil);
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

	manuallyDefinePump(server)

	status := ua.ServerRun(server, &running)
	ua.ServerDelete(server)
	os.Exit(int(status))
}
