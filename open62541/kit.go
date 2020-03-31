package open62541

/*
#cgo LDFLAGS: -lopen62541
#include "open62541/types_generated.h"
*/
import "C"
import "unsafe"

// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
var TYPES = (*[373]DataType)(unsafe.Pointer(&C.UA_TYPES))[:373:373]

func (v *DataValue) SetValue(value Variant) {
	v.value = C.UA_Variant(value)
}

func (v *DataValue) SetStatus(status uint32) {
	v.status = C.UA_StatusCode(status)
}
