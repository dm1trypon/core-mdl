package proto

import (
	"github.com/dm1trypon/core-mdl/internal/app/proto/adduserdata"
	"github.com/dm1trypon/core-mdl/internal/app/proto/getuserdata"
)

type Proto struct {
	lc             string
	schemas        []SchemaData
	getUDProtoInst *getuserdata.GetUserDataProto
	addUDProtoInst *adduserdata.AddUserDataProto
}

type SchemaData struct {
	Message    interface{}
	SchemaMain string
	SchemaData string
}
