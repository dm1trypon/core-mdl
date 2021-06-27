package adduserdata

import (
	"context"
	"encoding/json"
	"fmt"

	addUserDataRecv "github.com/dm1trypon/core-mdl/internal/app/proto/adduserdata/recieve"
	addUserDataResp "github.com/dm1trypon/core-mdl/internal/app/proto/adduserdata/response"
	logger "github.com/dm1trypon/easy-logger"
	"github.com/qri-io/jsonschema"
)

func (a *AddUserDataProto) Create() *AddUserDataProto {
	a = &AddUserDataProto{
		lc: "GET_USER_DATA_PROTO",
	}

	return a
}

func (a *AddUserDataProto) OnMessage(body []byte, schemaData string) (interface{}, interface{}) {
	ctx := context.Background()
	msg := addUserDataRecv.InternalMessage{}

	if err := json.Unmarshal(body, &msg); err != nil {
		return nil, nil
	}

	rsData := &jsonschema.Schema{}
	if err := json.Unmarshal([]byte(schemaData), rsData); err != nil {
		logger.ErrorJ(a.lc, fmt.Sprint("Wrong JSON DATA schema: ", err.Error()))
		return nil, nil
	}

	errors := addUserDataResp.Errors{}

	errs, err := rsData.ValidateBytes(ctx, []byte(msg.Data))
	if err != nil {
		logger.ErrorJ(a.lc, fmt.Sprint("Error DATA validate: ", err.Error()))
		return nil, nil
	}

	for _, err := range errs {
		logger.ErrorJ(a.lc, fmt.Sprint("Error validate: ", err.Error()))
		errData := addUserDataResp.Error{
			Code:    -32700,
			Message: err.Error(),
			Data:    addUserDataResp.ErrorData{},
		}

		errors = append(errors, errData)
	}

	if len(errors) > 0 {
		return msg, errors
	}

	data := addUserDataRecv.Data{}

	if err := json.Unmarshal(msg.Data, &data); err != nil {
		logger.ErrorJ(a.lc, fmt.Sprint("Wrong message DATA: ", err.Error()))
		return nil, nil
	}

	return msg, data
}

func (a *AddUserDataProto) MakeBody(recv addUserDataRecv.InternalMessage, result interface{}) interface{} {
	var respBody interface{}

	switch result.(type) {
	case addUserDataResp.Data:
		respBody = addUserDataResp.IMSuccess{
			ID:       recv.ID,
			GID:      recv.CID,
			CID:      recv.GID,
			Expires:  recv.Expires,
			Features: recv.Features,
			Locale:   recv.Locale,
			Domain:   recv.Domain,
			Event:    recv.Event,
			Data:     result.(addUserDataResp.Data),
			Saga: addUserDataResp.Saga{
				SID:     recv.Saga.SID,
				Name:    recv.Saga.Name,
				State:   recv.Saga.State,
				Expires: recv.Expires,
			},
			Src: addUserDataResp.Src{
				Name: recv.Src.Name,
				Version: addUserDataResp.Version{
					Major: recv.Src.Version.Major,
					Minor: recv.Src.Version.Minor,
					Patch: recv.Src.Version.Patch,
				},
			},
			ReplyTo: addUserDataResp.ReplyTo{
				Exchange:   recv.ReplyTo.Exchange,
				RoutingKey: recv.ReplyTo.RoutingKey,
				Event:      recv.ReplyTo.Event,
			},
			Tokens: addUserDataResp.Tokens{
				ID:     recv.Tokens.ID,
				Access: recv.Tokens.Access,
			},
		}
	case addUserDataResp.Errors:
		respBody = addUserDataResp.IMError{
			ID:       recv.ID,
			GID:      recv.CID,
			CID:      recv.GID,
			Expires:  recv.Expires,
			Features: recv.Features,
			Locale:   recv.Locale,
			Domain:   recv.Domain,
			Event:    recv.Event,
			Errors:   result.(addUserDataResp.Errors),
			Saga: addUserDataResp.Saga{
				SID:     recv.Saga.SID,
				Name:    recv.Saga.Name,
				State:   recv.Saga.State,
				Expires: recv.Expires,
			},
			Src: addUserDataResp.Src{
				Name: recv.Src.Name,
				Version: addUserDataResp.Version{
					Major: recv.Src.Version.Major,
					Minor: recv.Src.Version.Minor,
					Patch: recv.Src.Version.Patch,
				},
			},
			ReplyTo: addUserDataResp.ReplyTo{
				Exchange:   recv.ReplyTo.Exchange,
				RoutingKey: recv.ReplyTo.RoutingKey,
				Event:      recv.ReplyTo.Event,
			},
			Tokens: addUserDataResp.Tokens{
				ID:     recv.Tokens.ID,
				Access: recv.Tokens.Access,
			},
		}
	}

	return respBody
}
