package open62541

/*
#cgo LDFLAGS: -lopen62541
#include "open62541/types_generated.h"

void setHasValue(UA_DataValue *value, bool hasValue) {
	if (value == NULL) { return; }
	value->hasValue = hasValue;
}

void setHasStatus(UA_DataValue *value, bool hasStatus) {
	if (value == NULL) { return; }
	value->hasStatus = hasStatus;
}

*/
import "C"
import "unsafe"

type (
	Int64T = int64
	Uint64T = uint64
)

// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
var TYPES = (*[373]DataType)(unsafe.Pointer(&C.UA_TYPES))[:373:373]

func (v *DataValue) SetValue(value Variant) {
	v.value = C.UA_Variant(value)
}

func (v *DataValue) SetStatus(status uint32) {
	v.status = C.UA_StatusCode(status)
}

func (v *DataValue) SetHasValue(b bool) {
	cv := (C.UA_DataValue)(*v)
	cb := C.bool(b)
	C.setHasValue(&cv, cb)
}

func (v *DataValue) SetHasStatus(b bool) {
	cv := (C.UA_DataValue)(*v)
	cb := C.bool(b)
	C.setHasStatus(&cv, cb)
}
