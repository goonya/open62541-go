all:
	~/go/bin/c-for-go open62541.yaml

clean:
	rm -f open62541/cgo_helpers.go open62541/cgo_helpers.h open62541/cgo_helpers.c
	rm -f open62541/const.go open62541/doc.go open62541/types.go
	rm -f open62541/open62541.go

test:
	cd example && go build