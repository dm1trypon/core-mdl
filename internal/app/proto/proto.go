package proto

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dm1trypon/core-mdl/internal/app/proto/adduserdata"
	addUserDataRecv "github.com/dm1trypon/core-mdl/internal/app/proto/adduserdata/recieve"
	"github.com/dm1trypon/core-mdl/internal/app/proto/getuserdata"
	getUserDataRecv "github.com/dm1trypon/core-mdl/internal/app/proto/getuserdata/recieve"
	logger "github.com/dm1trypon/easy-logger"
	"github.com/qri-io/jsonschema"
)

func (p *Proto) Create() *Proto {
	p = &Proto{
		lc: "PROTO",
		schemas: []SchemaData{
			{
				Message:    getUserDataRecv.InternalMessage{},
				SchemaMain: getUserDataRecv.SchemaMain,
				SchemaData: getUserDataRecv.SchemaData,
			},
			{
				Message:    addUserDataRecv.InternalMessage{},
				SchemaMain: addUserDataRecv.SchemaMain,
				SchemaData: addUserDataRecv.SchemaData,
			},
		},
		getUDProtoInst: new(getuserdata.GetUserDataProto).Create(),
		addUDProtoInst: new(adduserdata.AddUserDataProto).Create(),
	}

	return p
}

func (p *Proto) Parse(body []byte) (interface{}, interface{}) {
	ctx := context.Background()
	rs := &jsonschema.Schema{}

	for _, schemaData := range p.schemas {
		if err := json.Unmarshal([]byte(schemaData.SchemaMain), rs); err != nil {
			logger.ErrorJ(p.lc, fmt.Sprint("Wrong JSON PROTO schema: ", err.Error()))
			continue
		}

		errs, err := rs.ValidateBytes(ctx, body)
		if err != nil {
			logger.ErrorJ(p.lc, fmt.Sprint("Wrong JSON PROTO message: ", err.Error()))
			continue
		}

		for _, err := range errs {
			logger.ErrorJ(p.lc, fmt.Sprint("Error validate JSON PROTO: ", err.Error()))
		}

		if len(errs) > 0 {
			continue
		}

		switch schemaData.Message.(type) {
		case getUserDataRecv.InternalMessage:
			return p.getUDProtoInst.OnMessage(body, schemaData.SchemaData)
		case addUserDataRecv.InternalMessage:
			return p.addUDProtoInst.OnMessage(body, schemaData.SchemaData)
		}
	}

	return nil, nil
}

func (p *Proto) Pack(recv, result interface{}) []byte {
	var respBody interface{}

	switch recv.(type) {
	case getUserDataRecv.InternalMessage:
		respBody = p.getUDProtoInst.MakeBody(recv.(getUserDataRecv.InternalMessage), result)
	case addUserDataRecv.InternalMessage:
		respBody = p.addUDProtoInst.MakeBody(recv.(addUserDataRecv.InternalMessage), result)
	}

	body, err := json.Marshal(respBody)
	if err != nil {
		logger.ErrorJ(p.lc, fmt.Sprint("Error packing JSON: ", err.Error()))
		return nil
	}

	return body
}
