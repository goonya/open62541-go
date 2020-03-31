// Mozilla Public License Version 2.0

// WARNING: This file has automatically been generated on Tue, 31 Mar 2020 21:39:35 MSK.
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
	"sync"
	"unsafe"
)

// Ref returns a reference to C object as it is.
func (x *ServerConfig) Ref() *C.UA_ServerConfig {
	if x == nil {
		return nil
	}
	return (*C.UA_ServerConfig)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *ServerConfig) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewServerConfigRef converts the C object reference into a raw struct reference without wrapping.
func NewServerConfigRef(ref unsafe.Pointer) *ServerConfig {
	return (*ServerConfig)(ref)
}

// NewServerConfig allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewServerConfig() *ServerConfig {
	return (*ServerConfig)(allocServerConfigMemory(1))
}

// allocServerConfigMemory allocates memory for type C.UA_ServerConfig in C.
// The caller is responsible for freeing the this memory via C.free.
func allocServerConfigMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfServerConfigValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfServerConfigValue = unsafe.Sizeof([1]C.UA_ServerConfig{})

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *ServerConfig) PassRef() *C.UA_ServerConfig {
	if x == nil {
		x = (*ServerConfig)(allocServerConfigMemory(1))
	}
	return (*C.UA_ServerConfig)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Server) Ref() *C.UA_Server {
	if x == nil {
		return nil
	}
	return (*C.UA_Server)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Server) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewServerRef converts the C object reference into a raw struct reference without wrapping.
func NewServerRef(ref unsafe.Pointer) *Server {
	return (*Server)(ref)
}

// NewServer allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewServer() *Server {
	return (*Server)(allocServerMemory(1))
}

// allocServerMemory allocates memory for type C.UA_Server in C.
// The caller is responsible for freeing the this memory via C.free.
func allocServerMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfServerValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfServerValue = unsafe.Sizeof([1]C.UA_Server{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Server) PassRef() *C.UA_Server {
	if x == nil {
		x = (*Server)(allocServerMemory(1))
	}
	return (*C.UA_Server)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *NodeTypeLifecycle) Ref() *C.UA_NodeTypeLifecycle {
	if x == nil {
		return nil
	}
	return (*C.UA_NodeTypeLifecycle)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *NodeTypeLifecycle) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNodeTypeLifecycleRef converts the C object reference into a raw struct reference without wrapping.
func NewNodeTypeLifecycleRef(ref unsafe.Pointer) *NodeTypeLifecycle {
	return (*NodeTypeLifecycle)(ref)
}

// NewNodeTypeLifecycle allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewNodeTypeLifecycle() *NodeTypeLifecycle {
	return (*NodeTypeLifecycle)(allocNodeTypeLifecycleMemory(1))
}

// allocNodeTypeLifecycleMemory allocates memory for type C.UA_NodeTypeLifecycle in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNodeTypeLifecycleMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNodeTypeLifecycleValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNodeTypeLifecycleValue = unsafe.Sizeof([1]C.UA_NodeTypeLifecycle{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *NodeTypeLifecycle) PassRef() *C.UA_NodeTypeLifecycle {
	if x == nil {
		x = (*NodeTypeLifecycle)(allocNodeTypeLifecycleMemory(1))
	}
	return (*C.UA_NodeTypeLifecycle)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *String) Ref() *C.UA_String {
	if x == nil {
		return nil
	}
	return (*C.UA_String)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *String) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewStringRef converts the C object reference into a raw struct reference without wrapping.
func NewStringRef(ref unsafe.Pointer) *String {
	return (*String)(ref)
}

// NewString allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewString() *String {
	return (*String)(allocStringMemory(1))
}

// allocStringMemory allocates memory for type C.UA_String in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStringMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStringValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStringValue = unsafe.Sizeof([1]C.UA_String{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *String) PassRef() *C.UA_String {
	if x == nil {
		x = (*String)(allocStringMemory(1))
	}
	return (*C.UA_String)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *ByteString) Ref() *C.UA_String {
	if x == nil {
		return nil
	}
	return (*C.UA_String)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *ByteString) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewByteStringRef converts the C object reference into a raw struct reference without wrapping.
func NewByteStringRef(ref unsafe.Pointer) *ByteString {
	return (*ByteString)(ref)
}

// NewByteString allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewByteString() *ByteString {
	return (*ByteString)(allocStringMemory(1))
}

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *ByteString) PassRef() *C.UA_String {
	if x == nil {
		x = (*ByteString)(allocStringMemory(1))
	}
	return (*C.UA_String)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *NodeId) Ref() *C.UA_NodeId {
	if x == nil {
		return nil
	}
	return (*C.UA_NodeId)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *NodeId) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewNodeIdRef converts the C object reference into a raw struct reference without wrapping.
func NewNodeIdRef(ref unsafe.Pointer) *NodeId {
	return (*NodeId)(ref)
}

// NewNodeId allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewNodeId() *NodeId {
	return (*NodeId)(allocNodeIdMemory(1))
}

// allocNodeIdMemory allocates memory for type C.UA_NodeId in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNodeIdMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNodeIdValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNodeIdValue = unsafe.Sizeof([1]C.UA_NodeId{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *NodeId) PassRef() *C.UA_NodeId {
	if x == nil {
		x = (*NodeId)(allocNodeIdMemory(1))
	}
	return (*C.UA_NodeId)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *QualifiedName) Ref() *C.UA_QualifiedName {
	if x == nil {
		return nil
	}
	return (*C.UA_QualifiedName)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *QualifiedName) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewQualifiedNameRef converts the C object reference into a raw struct reference without wrapping.
func NewQualifiedNameRef(ref unsafe.Pointer) *QualifiedName {
	return (*QualifiedName)(ref)
}

// NewQualifiedName allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewQualifiedName() *QualifiedName {
	return (*QualifiedName)(allocQualifiedNameMemory(1))
}

// allocQualifiedNameMemory allocates memory for type C.UA_QualifiedName in C.
// The caller is responsible for freeing the this memory via C.free.
func allocQualifiedNameMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfQualifiedNameValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfQualifiedNameValue = unsafe.Sizeof([1]C.UA_QualifiedName{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *QualifiedName) PassRef() *C.UA_QualifiedName {
	if x == nil {
		x = (*QualifiedName)(allocQualifiedNameMemory(1))
	}
	return (*C.UA_QualifiedName)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *LocalizedText) Ref() *C.UA_LocalizedText {
	if x == nil {
		return nil
	}
	return (*C.UA_LocalizedText)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *LocalizedText) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewLocalizedTextRef converts the C object reference into a raw struct reference without wrapping.
func NewLocalizedTextRef(ref unsafe.Pointer) *LocalizedText {
	return (*LocalizedText)(ref)
}

// NewLocalizedText allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewLocalizedText() *LocalizedText {
	return (*LocalizedText)(allocLocalizedTextMemory(1))
}

// allocLocalizedTextMemory allocates memory for type C.UA_LocalizedText in C.
// The caller is responsible for freeing the this memory via C.free.
func allocLocalizedTextMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfLocalizedTextValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfLocalizedTextValue = unsafe.Sizeof([1]C.UA_LocalizedText{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *LocalizedText) PassRef() *C.UA_LocalizedText {
	if x == nil {
		x = (*LocalizedText)(allocLocalizedTextMemory(1))
	}
	return (*C.UA_LocalizedText)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *DataType) Ref() *C.UA_DataType {
	if x == nil {
		return nil
	}
	return (*C.UA_DataType)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *DataType) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewDataTypeRef converts the C object reference into a raw struct reference without wrapping.
func NewDataTypeRef(ref unsafe.Pointer) *DataType {
	return (*DataType)(ref)
}

// NewDataType allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewDataType() *DataType {
	return (*DataType)(allocDataTypeMemory(1))
}

// allocDataTypeMemory allocates memory for type C.UA_DataType in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDataTypeMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDataTypeValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfDataTypeValue = unsafe.Sizeof([1]C.UA_DataType{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *DataType) PassRef() *C.UA_DataType {
	if x == nil {
		x = (*DataType)(allocDataTypeMemory(1))
	}
	return (*C.UA_DataType)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *Variant) Ref() *C.UA_Variant {
	if x == nil {
		return nil
	}
	return (*C.UA_Variant)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *Variant) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewVariantRef converts the C object reference into a raw struct reference without wrapping.
func NewVariantRef(ref unsafe.Pointer) *Variant {
	return (*Variant)(ref)
}

// NewVariant allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewVariant() *Variant {
	return (*Variant)(allocVariantMemory(1))
}

// allocVariantMemory allocates memory for type C.UA_Variant in C.
// The caller is responsible for freeing the this memory via C.free.
func allocVariantMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfVariantValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfVariantValue = unsafe.Sizeof([1]C.UA_Variant{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *Variant) PassRef() *C.UA_Variant {
	if x == nil {
		x = (*Variant)(allocVariantMemory(1))
	}
	return (*C.UA_Variant)(unsafe.Pointer(x))
}

// Ref returns a reference to C object as it is.
func (x *DataValue) Ref() *C.UA_DataValue {
	if x == nil {
		return nil
	}
	return (*C.UA_DataValue)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *DataValue) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewDataValueRef converts the C object reference into a raw struct reference without wrapping.
func NewDataValueRef(ref unsafe.Pointer) *DataValue {
	return (*DataValue)(ref)
}

// NewDataValue allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewDataValue() *DataValue {
	return (*DataValue)(allocDataValueMemory(1))
}

// allocDataValueMemory allocates memory for type C.UA_DataValue in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDataValueMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDataValueValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfDataValueValue = unsafe.Sizeof([1]C.UA_DataValue{})

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *DataValue) PassRef() *C.UA_DataValue {
	if x == nil {
		x = (*DataValue)(allocDataValueMemory(1))
	}
	return (*C.UA_DataValue)(unsafe.Pointer(x))
}

// allocObjectAttributesMemory allocates memory for type C.UA_ObjectAttributes in C.
// The caller is responsible for freeing the this memory via C.free.
func allocObjectAttributesMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfObjectAttributesValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfObjectAttributesValue = unsafe.Sizeof([1]C.UA_ObjectAttributes{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ObjectAttributes) Ref() *C.UA_ObjectAttributes {
	if x == nil {
		return nil
	}
	return x.refe8e3a938
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ObjectAttributes) Free() {
	if x != nil && x.allocse8e3a938 != nil {
		x.allocse8e3a938.(*cgoAllocMap).Free()
		x.refe8e3a938 = nil
	}
}

// NewObjectAttributesRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewObjectAttributesRef(ref unsafe.Pointer) *ObjectAttributes {
	if ref == nil {
		return nil
	}
	obj := new(ObjectAttributes)
	obj.refe8e3a938 = (*C.UA_ObjectAttributes)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ObjectAttributes) PassRef() (*C.UA_ObjectAttributes, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe8e3a938 != nil {
		return x.refe8e3a938, nil
	}
	meme8e3a938 := allocObjectAttributesMemory(1)
	refe8e3a938 := (*C.UA_ObjectAttributes)(meme8e3a938)
	allocse8e3a938 := new(cgoAllocMap)
	allocse8e3a938.Add(meme8e3a938)

	var cspecifiedAttributes_allocs *cgoAllocMap
	refe8e3a938.specifiedAttributes, cspecifiedAttributes_allocs = (C.UA_UInt32)(x.SpecifiedAttributes), cgoAllocsUnknown
	allocse8e3a938.Borrow(cspecifiedAttributes_allocs)

	var cdisplayName_allocs *cgoAllocMap
	refe8e3a938.displayName, cdisplayName_allocs = *(*C.UA_LocalizedText)(unsafe.Pointer(&x.DisplayName)), cgoAllocsUnknown
	allocse8e3a938.Borrow(cdisplayName_allocs)

	var cdescription_allocs *cgoAllocMap
	refe8e3a938.description, cdescription_allocs = *(*C.UA_LocalizedText)(unsafe.Pointer(&x.Description)), cgoAllocsUnknown
	allocse8e3a938.Borrow(cdescription_allocs)

	var cwriteMask_allocs *cgoAllocMap
	refe8e3a938.writeMask, cwriteMask_allocs = (C.UA_UInt32)(x.WriteMask), cgoAllocsUnknown
	allocse8e3a938.Borrow(cwriteMask_allocs)

	var cuserWriteMask_allocs *cgoAllocMap
	refe8e3a938.userWriteMask, cuserWriteMask_allocs = (C.UA_UInt32)(x.UserWriteMask), cgoAllocsUnknown
	allocse8e3a938.Borrow(cuserWriteMask_allocs)

	var ceventNotifier_allocs *cgoAllocMap
	refe8e3a938.eventNotifier, ceventNotifier_allocs = (C.UA_Byte)(x.EventNotifier), cgoAllocsUnknown
	allocse8e3a938.Borrow(ceventNotifier_allocs)

	x.refe8e3a938 = refe8e3a938
	x.allocse8e3a938 = allocse8e3a938
	return refe8e3a938, allocse8e3a938

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ObjectAttributes) PassValue() (C.UA_ObjectAttributes, *cgoAllocMap) {
	if x.refe8e3a938 != nil {
		return *x.refe8e3a938, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ObjectAttributes) Deref() {
	if x.refe8e3a938 == nil {
		return
	}
	x.SpecifiedAttributes = (uint32)(x.refe8e3a938.specifiedAttributes)
	x.DisplayName = *(*LocalizedText)(unsafe.Pointer(&x.refe8e3a938.displayName))
	x.Description = *(*LocalizedText)(unsafe.Pointer(&x.refe8e3a938.description))
	x.WriteMask = (uint32)(x.refe8e3a938.writeMask)
	x.UserWriteMask = (uint32)(x.refe8e3a938.userWriteMask)
	x.EventNotifier = (byte)(x.refe8e3a938.eventNotifier)
}

// allocVariableAttributesMemory allocates memory for type C.UA_VariableAttributes in C.
// The caller is responsible for freeing the this memory via C.free.
func allocVariableAttributesMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfVariableAttributesValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfVariableAttributesValue = unsafe.Sizeof([1]C.UA_VariableAttributes{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *VariableAttributes) Ref() *C.UA_VariableAttributes {
	if x == nil {
		return nil
	}
	return x.refd219e026
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *VariableAttributes) Free() {
	if x != nil && x.allocsd219e026 != nil {
		x.allocsd219e026.(*cgoAllocMap).Free()
		x.refd219e026 = nil
	}
}

// NewVariableAttributesRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewVariableAttributesRef(ref unsafe.Pointer) *VariableAttributes {
	if ref == nil {
		return nil
	}
	obj := new(VariableAttributes)
	obj.refd219e026 = (*C.UA_VariableAttributes)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *VariableAttributes) PassRef() (*C.UA_VariableAttributes, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd219e026 != nil {
		return x.refd219e026, nil
	}
	memd219e026 := allocVariableAttributesMemory(1)
	refd219e026 := (*C.UA_VariableAttributes)(memd219e026)
	allocsd219e026 := new(cgoAllocMap)
	allocsd219e026.Add(memd219e026)

	var cspecifiedAttributes_allocs *cgoAllocMap
	refd219e026.specifiedAttributes, cspecifiedAttributes_allocs = (C.UA_UInt32)(x.SpecifiedAttributes), cgoAllocsUnknown
	allocsd219e026.Borrow(cspecifiedAttributes_allocs)

	var cdisplayName_allocs *cgoAllocMap
	refd219e026.displayName, cdisplayName_allocs = *(*C.UA_LocalizedText)(unsafe.Pointer(&x.DisplayName)), cgoAllocsUnknown
	allocsd219e026.Borrow(cdisplayName_allocs)

	var cdescription_allocs *cgoAllocMap
	refd219e026.description, cdescription_allocs = *(*C.UA_LocalizedText)(unsafe.Pointer(&x.Description)), cgoAllocsUnknown
	allocsd219e026.Borrow(cdescription_allocs)

	var cwriteMask_allocs *cgoAllocMap
	refd219e026.writeMask, cwriteMask_allocs = (C.UA_UInt32)(x.WriteMask), cgoAllocsUnknown
	allocsd219e026.Borrow(cwriteMask_allocs)

	var cuserWriteMask_allocs *cgoAllocMap
	refd219e026.userWriteMask, cuserWriteMask_allocs = (C.UA_UInt32)(x.UserWriteMask), cgoAllocsUnknown
	allocsd219e026.Borrow(cuserWriteMask_allocs)

	var cvalue_allocs *cgoAllocMap
	refd219e026.value, cvalue_allocs = *(*C.UA_Variant)(unsafe.Pointer(&x.Value)), cgoAllocsUnknown
	allocsd219e026.Borrow(cvalue_allocs)

	var cdataType_allocs *cgoAllocMap
	refd219e026.dataType, cdataType_allocs = *(*C.UA_NodeId)(unsafe.Pointer(&x.DataType)), cgoAllocsUnknown
	allocsd219e026.Borrow(cdataType_allocs)

	var cvalueRank_allocs *cgoAllocMap
	refd219e026.valueRank, cvalueRank_allocs = (C.UA_Int32)(x.ValueRank), cgoAllocsUnknown
	allocsd219e026.Borrow(cvalueRank_allocs)

	var carrayDimensionsSize_allocs *cgoAllocMap
	refd219e026.arrayDimensionsSize, carrayDimensionsSize_allocs = (C.size_t)(x.ArrayDimensionsSize), cgoAllocsUnknown
	allocsd219e026.Borrow(carrayDimensionsSize_allocs)

	var carrayDimensions_allocs *cgoAllocMap
	refd219e026.arrayDimensions, carrayDimensions_allocs = *(**C.UA_UInt32)(unsafe.Pointer(&x.ArrayDimensions)), cgoAllocsUnknown
	allocsd219e026.Borrow(carrayDimensions_allocs)

	var caccessLevel_allocs *cgoAllocMap
	refd219e026.accessLevel, caccessLevel_allocs = (C.UA_Byte)(x.AccessLevel), cgoAllocsUnknown
	allocsd219e026.Borrow(caccessLevel_allocs)

	var cuserAccessLevel_allocs *cgoAllocMap
	refd219e026.userAccessLevel, cuserAccessLevel_allocs = (C.UA_Byte)(x.UserAccessLevel), cgoAllocsUnknown
	allocsd219e026.Borrow(cuserAccessLevel_allocs)

	var cminimumSamplingInterval_allocs *cgoAllocMap
	refd219e026.minimumSamplingInterval, cminimumSamplingInterval_allocs = (C.UA_Double)(x.MinimumSamplingInterval), cgoAllocsUnknown
	allocsd219e026.Borrow(cminimumSamplingInterval_allocs)

	var chistorizing_allocs *cgoAllocMap
	refd219e026.historizing, chistorizing_allocs = (C.UA_Boolean)(x.Historizing), cgoAllocsUnknown
	allocsd219e026.Borrow(chistorizing_allocs)

	x.refd219e026 = refd219e026
	x.allocsd219e026 = allocsd219e026
	return refd219e026, allocsd219e026

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x VariableAttributes) PassValue() (C.UA_VariableAttributes, *cgoAllocMap) {
	if x.refd219e026 != nil {
		return *x.refd219e026, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *VariableAttributes) Deref() {
	if x.refd219e026 == nil {
		return
	}
	x.SpecifiedAttributes = (uint32)(x.refd219e026.specifiedAttributes)
	x.DisplayName = *(*LocalizedText)(unsafe.Pointer(&x.refd219e026.displayName))
	x.Description = *(*LocalizedText)(unsafe.Pointer(&x.refd219e026.description))
	x.WriteMask = (uint32)(x.refd219e026.writeMask)
	x.UserWriteMask = (uint32)(x.refd219e026.userWriteMask)
	x.Value = *(*Variant)(unsafe.Pointer(&x.refd219e026.value))
	x.DataType = *(*NodeId)(unsafe.Pointer(&x.refd219e026.dataType))
	x.ValueRank = (Int32)(x.refd219e026.valueRank)
	x.ArrayDimensionsSize = (uint)(x.refd219e026.arrayDimensionsSize)
	x.ArrayDimensions = (*uint32)(unsafe.Pointer(x.refd219e026.arrayDimensions))
	x.AccessLevel = (byte)(x.refd219e026.accessLevel)
	x.UserAccessLevel = (byte)(x.refd219e026.userAccessLevel)
	x.MinimumSamplingInterval = (Double)(x.refd219e026.minimumSamplingInterval)
	x.Historizing = (Boolean)(x.refd219e026.historizing)
}

// allocObjectTypeAttributesMemory allocates memory for type C.UA_ObjectTypeAttributes in C.
// The caller is responsible for freeing the this memory via C.free.
func allocObjectTypeAttributesMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfObjectTypeAttributesValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfObjectTypeAttributesValue = unsafe.Sizeof([1]C.UA_ObjectTypeAttributes{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ObjectTypeAttributes) Ref() *C.UA_ObjectTypeAttributes {
	if x == nil {
		return nil
	}
	return x.refe0f167e4
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ObjectTypeAttributes) Free() {
	if x != nil && x.allocse0f167e4 != nil {
		x.allocse0f167e4.(*cgoAllocMap).Free()
		x.refe0f167e4 = nil
	}
}

// NewObjectTypeAttributesRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewObjectTypeAttributesRef(ref unsafe.Pointer) *ObjectTypeAttributes {
	if ref == nil {
		return nil
	}
	obj := new(ObjectTypeAttributes)
	obj.refe0f167e4 = (*C.UA_ObjectTypeAttributes)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ObjectTypeAttributes) PassRef() (*C.UA_ObjectTypeAttributes, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refe0f167e4 != nil {
		return x.refe0f167e4, nil
	}
	meme0f167e4 := allocObjectTypeAttributesMemory(1)
	refe0f167e4 := (*C.UA_ObjectTypeAttributes)(meme0f167e4)
	allocse0f167e4 := new(cgoAllocMap)
	allocse0f167e4.Add(meme0f167e4)

	var cspecifiedAttributes_allocs *cgoAllocMap
	refe0f167e4.specifiedAttributes, cspecifiedAttributes_allocs = (C.UA_UInt32)(x.SpecifiedAttributes), cgoAllocsUnknown
	allocse0f167e4.Borrow(cspecifiedAttributes_allocs)

	var cdisplayName_allocs *cgoAllocMap
	refe0f167e4.displayName, cdisplayName_allocs = *(*C.UA_LocalizedText)(unsafe.Pointer(&x.DisplayName)), cgoAllocsUnknown
	allocse0f167e4.Borrow(cdisplayName_allocs)

	var cdescription_allocs *cgoAllocMap
	refe0f167e4.description, cdescription_allocs = *(*C.UA_LocalizedText)(unsafe.Pointer(&x.Description)), cgoAllocsUnknown
	allocse0f167e4.Borrow(cdescription_allocs)

	var cwriteMask_allocs *cgoAllocMap
	refe0f167e4.writeMask, cwriteMask_allocs = (C.UA_UInt32)(x.WriteMask), cgoAllocsUnknown
	allocse0f167e4.Borrow(cwriteMask_allocs)

	var cuserWriteMask_allocs *cgoAllocMap
	refe0f167e4.userWriteMask, cuserWriteMask_allocs = (C.UA_UInt32)(x.UserWriteMask), cgoAllocsUnknown
	allocse0f167e4.Borrow(cuserWriteMask_allocs)

	var cisAbstract_allocs *cgoAllocMap
	refe0f167e4.isAbstract, cisAbstract_allocs = (C.UA_Boolean)(x.IsAbstract), cgoAllocsUnknown
	allocse0f167e4.Borrow(cisAbstract_allocs)

	x.refe0f167e4 = refe0f167e4
	x.allocse0f167e4 = allocse0f167e4
	return refe0f167e4, allocse0f167e4

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x ObjectTypeAttributes) PassValue() (C.UA_ObjectTypeAttributes, *cgoAllocMap) {
	if x.refe0f167e4 != nil {
		return *x.refe0f167e4, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ObjectTypeAttributes) Deref() {
	if x.refe0f167e4 == nil {
		return
	}
	x.SpecifiedAttributes = (uint32)(x.refe0f167e4.specifiedAttributes)
	x.DisplayName = *(*LocalizedText)(unsafe.Pointer(&x.refe0f167e4.displayName))
	x.Description = *(*LocalizedText)(unsafe.Pointer(&x.refe0f167e4.description))
	x.WriteMask = (uint32)(x.refe0f167e4.writeMask)
	x.UserWriteMask = (uint32)(x.refe0f167e4.userWriteMask)
	x.IsAbstract = (Boolean)(x.refe0f167e4.isAbstract)
}

// allocWriteValueMemory allocates memory for type C.UA_WriteValue in C.
// The caller is responsible for freeing the this memory via C.free.
func allocWriteValueMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfWriteValueValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfWriteValueValue = unsafe.Sizeof([1]C.UA_WriteValue{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *WriteValue) Ref() *C.UA_WriteValue {
	if x == nil {
		return nil
	}
	return x.ref4be23793
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *WriteValue) Free() {
	if x != nil && x.allocs4be23793 != nil {
		x.allocs4be23793.(*cgoAllocMap).Free()
		x.ref4be23793 = nil
	}
}

// NewWriteValueRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewWriteValueRef(ref unsafe.Pointer) *WriteValue {
	if ref == nil {
		return nil
	}
	obj := new(WriteValue)
	obj.ref4be23793 = (*C.UA_WriteValue)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *WriteValue) PassRef() (*C.UA_WriteValue, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref4be23793 != nil {
		return x.ref4be23793, nil
	}
	mem4be23793 := allocWriteValueMemory(1)
	ref4be23793 := (*C.UA_WriteValue)(mem4be23793)
	allocs4be23793 := new(cgoAllocMap)
	allocs4be23793.Add(mem4be23793)

	var cnodeId_allocs *cgoAllocMap
	ref4be23793.nodeId, cnodeId_allocs = *(*C.UA_NodeId)(unsafe.Pointer(&x.NodeId)), cgoAllocsUnknown
	allocs4be23793.Borrow(cnodeId_allocs)

	var cattributeId_allocs *cgoAllocMap
	ref4be23793.attributeId, cattributeId_allocs = (C.UA_UInt32)(x.AttributeId), cgoAllocsUnknown
	allocs4be23793.Borrow(cattributeId_allocs)

	var cindexRange_allocs *cgoAllocMap
	ref4be23793.indexRange, cindexRange_allocs = *(*C.UA_String)(unsafe.Pointer(&x.IndexRange)), cgoAllocsUnknown
	allocs4be23793.Borrow(cindexRange_allocs)

	var cvalue_allocs *cgoAllocMap
	ref4be23793.value, cvalue_allocs = *(*C.UA_DataValue)(unsafe.Pointer(&x.Value)), cgoAllocsUnknown
	allocs4be23793.Borrow(cvalue_allocs)

	x.ref4be23793 = ref4be23793
	x.allocs4be23793 = allocs4be23793
	return ref4be23793, allocs4be23793

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x WriteValue) PassValue() (C.UA_WriteValue, *cgoAllocMap) {
	if x.ref4be23793 != nil {
		return *x.ref4be23793, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *WriteValue) Deref() {
	if x.ref4be23793 == nil {
		return
	}
	x.NodeId = *(*NodeId)(unsafe.Pointer(&x.ref4be23793.nodeId))
	x.AttributeId = (uint32)(x.ref4be23793.attributeId)
	x.IndexRange = *(*String)(unsafe.Pointer(&x.ref4be23793.indexRange))
	x.Value = *(*DataValue)(unsafe.Pointer(&x.ref4be23793.value))
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}
