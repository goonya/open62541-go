// Mozilla Public License Version 2.0

// WARNING: This file has automatically been generated on Thu, 02 Apr 2020 14:28:43 MSK.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package open62541

/*
#cgo LDFLAGS: -lopen62541
#include "open62541/architecture_base.h"
#include "open62541/architecture_definitions.h"
#include "open62541/architecture_functions.h"
#include "open62541/client_config_default.h"
#include "open62541/client_config.h"
#include "open62541/client.h"
#include "open62541/client_highlevel_async.h"
#include "open62541/client_highlevel.h"
#include "open62541/client_subscriptions.h"
#include "open62541/config.h"
#include "open62541/constants.h"
#include "open62541/ms_stdint.h"
#include "open62541/network_tcp.h"
#include "open62541/nodeids.h"
#include "open62541/server_config_default.h"
#include "open62541/server_config.h"
#include "open62541/server.h"
#include "open62541/server_pubsub.h"
#include "open62541/statuscodes.h"
#include "open62541/types_generated.h"
#include "open62541/types_generated_handling.h"
#include "open62541/types.h"
#include "open62541/util.h"
#include "open62541/ziptree.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// ServerNew function as declared in open62541/server.h:45
func ServerNew() *Server {
	__ret := C.UA_Server_new()
	__v := *(**Server)(unsafe.Pointer(&__ret))
	return __v
}

// ServerNewWithConfig function as declared in open62541/server.h:49
func ServerNewWithConfig(config *ServerConfig) *Server {
	cconfig, _ := (*C.UA_ServerConfig)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.UA_Server_newWithConfig(cconfig)
	__v := *(**Server)(unsafe.Pointer(&__ret))
	return __v
}

// ServerDelete function as declared in open62541/server.h:52
func ServerDelete(server *Server) {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	C.UA_Server_delete(cserver)
}

// ServerGetConfig function as declared in open62541/server.h:54
func ServerGetConfig(server *Server) *ServerConfig {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	__ret := C.UA_Server_getConfig(cserver)
	__v := *(**ServerConfig)(unsafe.Pointer(&__ret))
	return __v
}

// ServerRun function as declared in open62541/server.h:65
func ServerRun(server *Server, running *bool) StatusCode {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	crunning, _ := (*C.UA_Boolean)(unsafe.Pointer(running)), cgoAllocsUnknown
	__ret := C.UA_Server_run(cserver, crunning)
	__v := (StatusCode)(__ret)
	return __v
}

// ServerWrite function as declared in open62541/server.h:334
func ServerWrite(server *Server, value *WriteValue) StatusCode {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	cvalue, _ := value.PassRef()
	__ret := C.UA_Server_write(cserver, cvalue)
	__v := (StatusCode)(__ret)
	return __v
}

// ServerWriteValue function as declared in open62541/server.h:393
func ServerWriteValue(server *Server, nodeId NodeId, value Variant) StatusCode {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	cnodeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&nodeId)), cgoAllocsUnknown
	cvalue, _ := *(*C.UA_Variant)(unsafe.Pointer(&value)), cgoAllocsUnknown
	__ret := C.UA_Server_writeValue(cserver, cnodeId, cvalue)
	__v := (StatusCode)(__ret)
	return __v
}

// ServerSetNodeTypeLifecycle function as declared in open62541/server.h:733
func ServerSetNodeTypeLifecycle(server *Server, nodeId NodeId, lifecycle NodeTypeLifecycle) StatusCode {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	cnodeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&nodeId)), cgoAllocsUnknown
	clifecycle, _ := *(*C.UA_NodeTypeLifecycle)(unsafe.Pointer(&lifecycle)), cgoAllocsUnknown
	__ret := C.UA_Server_setNodeTypeLifecycle(cserver, cnodeId, clifecycle)
	__v := (StatusCode)(__ret)
	return __v
}

// ServerAddVariableNode function as declared in open62541/server.h:1057
func ServerAddVariableNode(server *Server, requestedNewNodeId NodeId, parentNodeId NodeId, referenceTypeId NodeId, browseName QualifiedName, typeDefinition NodeId, attr VariableAttributes, nodeContext unsafe.Pointer, outNewNodeId *NodeId) StatusCode {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	crequestedNewNodeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&requestedNewNodeId)), cgoAllocsUnknown
	cparentNodeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&parentNodeId)), cgoAllocsUnknown
	creferenceTypeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&referenceTypeId)), cgoAllocsUnknown
	cbrowseName, _ := *(*C.UA_QualifiedName)(unsafe.Pointer(&browseName)), cgoAllocsUnknown
	ctypeDefinition, _ := *(*C.UA_NodeId)(unsafe.Pointer(&typeDefinition)), cgoAllocsUnknown
	cattr, _ := attr.PassValue()
	cnodeContext, _ := nodeContext, cgoAllocsUnknown
	coutNewNodeId, _ := (*C.UA_NodeId)(unsafe.Pointer(outNewNodeId)), cgoAllocsUnknown
	__ret := C.UA_Server_addVariableNode(cserver, crequestedNewNodeId, cparentNodeId, creferenceTypeId, cbrowseName, ctypeDefinition, cattr, cnodeContext, coutNewNodeId)
	__v := (StatusCode)(__ret)
	return __v
}

// ServerAddObjectNode function as declared in open62541/server.h:1089
func ServerAddObjectNode(server *Server, requestedNewNodeId NodeId, parentNodeId NodeId, referenceTypeId NodeId, browseName QualifiedName, typeDefinition NodeId, attr ObjectAttributes, nodeContext unsafe.Pointer, outNewNodeId *NodeId) StatusCode {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	crequestedNewNodeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&requestedNewNodeId)), cgoAllocsUnknown
	cparentNodeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&parentNodeId)), cgoAllocsUnknown
	creferenceTypeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&referenceTypeId)), cgoAllocsUnknown
	cbrowseName, _ := *(*C.UA_QualifiedName)(unsafe.Pointer(&browseName)), cgoAllocsUnknown
	ctypeDefinition, _ := *(*C.UA_NodeId)(unsafe.Pointer(&typeDefinition)), cgoAllocsUnknown
	cattr, _ := attr.PassValue()
	cnodeContext, _ := nodeContext, cgoAllocsUnknown
	coutNewNodeId, _ := (*C.UA_NodeId)(unsafe.Pointer(outNewNodeId)), cgoAllocsUnknown
	__ret := C.UA_Server_addObjectNode(cserver, crequestedNewNodeId, cparentNodeId, creferenceTypeId, cbrowseName, ctypeDefinition, cattr, cnodeContext, coutNewNodeId)
	__v := (StatusCode)(__ret)
	return __v
}

// ServerAddObjectTypeNode function as declared in open62541/server.h:1104
func ServerAddObjectTypeNode(server *Server, requestedNewNodeId NodeId, parentNodeId NodeId, referenceTypeId NodeId, browseName QualifiedName, attr ObjectTypeAttributes, nodeContext unsafe.Pointer, outNewNodeId *NodeId) StatusCode {
	cserver, _ := (*C.UA_Server)(unsafe.Pointer(server)), cgoAllocsUnknown
	crequestedNewNodeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&requestedNewNodeId)), cgoAllocsUnknown
	cparentNodeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&parentNodeId)), cgoAllocsUnknown
	creferenceTypeId, _ := *(*C.UA_NodeId)(unsafe.Pointer(&referenceTypeId)), cgoAllocsUnknown
	cbrowseName, _ := *(*C.UA_QualifiedName)(unsafe.Pointer(&browseName)), cgoAllocsUnknown
	cattr, _ := attr.PassValue()
	cnodeContext, _ := nodeContext, cgoAllocsUnknown
	coutNewNodeId, _ := (*C.UA_NodeId)(unsafe.Pointer(outNewNodeId)), cgoAllocsUnknown
	__ret := C.UA_Server_addObjectTypeNode(cserver, crequestedNewNodeId, cparentNodeId, creferenceTypeId, cbrowseName, cattr, cnodeContext, coutNewNodeId)
	__v := (StatusCode)(__ret)
	return __v
}

// StringFromChars function as declared in open62541/types.h:158
func StringFromChars(src string) String {
	src = safeString(src)
	csrc, _ := unpackPCharString(src)
	__ret := C.UA_String_fromChars(csrc)
	runtime.KeepAlive(src)
	__v := *(*String)(unsafe.Pointer(&__ret))
	return __v
}

// STRING function as declared in open62541/types.h:169
func STRING(chars []byte) String {
	cchars, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&chars)).Data)), cgoAllocsUnknown
	__ret := C.UA_STRING(cchars)
	__v := *(*String)(unsafe.Pointer(&__ret))
	return __v
}

// NODEIDNUMERIC function as declared in open62541/types.h:344
func NODEIDNUMERIC(nsIndex uint16, identifier uint32) NodeId {
	cnsIndex, _ := (C.UA_UInt16)(nsIndex), cgoAllocsUnknown
	cidentifier, _ := (C.UA_UInt32)(identifier), cgoAllocsUnknown
	__ret := C.UA_NODEID_NUMERIC(cnsIndex, cidentifier)
	__v := *(*NodeId)(unsafe.Pointer(&__ret))
	return __v
}

// NODEIDSTRING function as declared in open62541/types.h:351
func NODEIDSTRING(nsIndex uint16, chars []byte) NodeId {
	cnsIndex, _ := (C.UA_UInt16)(nsIndex), cgoAllocsUnknown
	cchars, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&chars)).Data)), cgoAllocsUnknown
	__ret := C.UA_NODEID_STRING(cnsIndex, cchars)
	__v := *(*NodeId)(unsafe.Pointer(&__ret))
	return __v
}

// QUALIFIEDNAME function as declared in open62541/types.h:462
func QUALIFIEDNAME(nsIndex uint16, chars []byte) QualifiedName {
	cnsIndex, _ := (C.UA_UInt16)(nsIndex), cgoAllocsUnknown
	cchars, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&chars)).Data)), cgoAllocsUnknown
	__ret := C.UA_QUALIFIEDNAME(cnsIndex, cchars)
	__v := *(*QualifiedName)(unsafe.Pointer(&__ret))
	return __v
}

// LOCALIZEDTEXT function as declared in open62541/types.h:487
func LOCALIZEDTEXT(locale []byte, text []byte) LocalizedText {
	clocale, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&locale)).Data)), cgoAllocsUnknown
	ctext, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&text)).Data)), cgoAllocsUnknown
	__ret := C.UA_LOCALIZEDTEXT(clocale, ctext)
	__v := *(*LocalizedText)(unsafe.Pointer(&__ret))
	return __v
}

// VariantSetScalar function as declared in open62541/types.h:627
func VariantSetScalar(v *Variant, p unsafe.Pointer, kind *DataType) {
	cv, _ := (*C.UA_Variant)(unsafe.Pointer(v)), cgoAllocsUnknown
	cp, _ := p, cgoAllocsUnknown
	ckind, _ := (*C.UA_DataType)(unsafe.Pointer(kind)), cgoAllocsUnknown
	C.UA_Variant_setScalar(cv, cp, ckind)
}

// VariantInit function as declared in open62541/types_generated_handling.h:747
func VariantInit(p *Variant) {
	cp, _ := (*C.UA_Variant)(unsafe.Pointer(p)), cgoAllocsUnknown
	C.UA_Variant_init(cp)
}

// WriteValueInit function as declared in open62541/types_generated_handling.h:7916
func WriteValueInit(p *WriteValue) {
	cp, _ := p.PassRef()
	C.UA_WriteValue_init(cp)
}

// ServerConfigSetMinimalCustomBuffer function as declared in open62541/server_config_default.h:43
func ServerConfigSetMinimalCustomBuffer(config *ServerConfig, portNumber uint16, certificate *ByteString, sendBufferSize uint32, recvBufferSize uint32) StatusCode {
	cconfig, _ := (*C.UA_ServerConfig)(unsafe.Pointer(config)), cgoAllocsUnknown
	cportNumber, _ := (C.UA_UInt16)(portNumber), cgoAllocsUnknown
	ccertificate, _ := (*C.UA_ByteString)(unsafe.Pointer(certificate)), cgoAllocsUnknown
	csendBufferSize, _ := (C.UA_UInt32)(sendBufferSize), cgoAllocsUnknown
	crecvBufferSize, _ := (C.UA_UInt32)(recvBufferSize), cgoAllocsUnknown
	__ret := C.UA_ServerConfig_setMinimalCustomBuffer(cconfig, cportNumber, ccertificate, csendBufferSize, crecvBufferSize)
	__v := (StatusCode)(__ret)
	return __v
}

// ServerConfigSetMinimal function as declared in open62541/server_config_default.h:55
func ServerConfigSetMinimal(config *ServerConfig, portNumber uint16, certificate *ByteString) StatusCode {
	cconfig, _ := (*C.UA_ServerConfig)(unsafe.Pointer(config)), cgoAllocsUnknown
	cportNumber, _ := (C.UA_UInt16)(portNumber), cgoAllocsUnknown
	ccertificate, _ := (*C.UA_ByteString)(unsafe.Pointer(certificate)), cgoAllocsUnknown
	__ret := C.UA_ServerConfig_setMinimal(cconfig, cportNumber, ccertificate)
	__v := (StatusCode)(__ret)
	return __v
}

// ServerConfigSetDefault function as declared in open62541/server_config_default.h:80
func ServerConfigSetDefault(config *ServerConfig) StatusCode {
	cconfig, _ := (*C.UA_ServerConfig)(unsafe.Pointer(config)), cgoAllocsUnknown
	__ret := C.UA_ServerConfig_setDefault(cconfig)
	__v := (StatusCode)(__ret)
	return __v
}
