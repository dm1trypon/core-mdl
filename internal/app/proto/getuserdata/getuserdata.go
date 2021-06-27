package getuserdata

import (
	"context"
	"encoding/json"
	"fmt"

	getUserDataRecv "github.com/dm1trypon/core-mdl/internal/app/proto/getuserdata/recieve"
	getUserDataResp "github.com/dm1trypon/core-mdl/internal/app/proto/getuserdata/response"
	logger "github.com/dm1trypon/easy-logger"
	"github.com/qri-io/jsonschema"
)

func (g *GetUserDataProto) Create() *GetUserDataProto {
	g = &GetUserDataProto{
		lc: "GET_USER_DATA_PROTO",
	}

	return g
}

func (g *GetUserDataProto) OnMessage(body []byte, schemaData string) (interface{}, interface{}) {
	ctx := context.Background()
	msg := getUserDataRecv.InternalMessage{}

	if err := json.Unmarshal(body, &msg); err != nil {
		return nil, nil
	}

	rsData := &jsonschema.Schema{}
	if err := json.Unmarshal([]byte(schemaData), rsData); err != nil {
		logger.ErrorJ(g.lc, fmt.Sprint("Wrong JSON DATA schema: ", err.Error()))
		return nil, nil
	}

	errors := getUserDataResp.Errors{}

	errs, err := rsData.ValidateBytes(ctx, []byte(msg.Data))
	if err != nil {
		logger.ErrorJ(g.lc, fmt.Sprint("Error DATA validate: ", err.Error()))
		return nil, nil
	}

	for _, err := range errs {
		logger.ErrorJ(g.lc, fmt.Sprint("Error validate: ", err.Error()))
		errData := getUserDataResp.Error{
			Code:    -32700,
			Message: err.Error(),
			Data:    getUserDataResp.ErrorData{},
		}

		errors = append(errors, errData)
	}

	if len(errors) > 0 {
		return msg, errors
	}

	data := getUserDataRecv.Data{}

	if err := json.Unmarshal(msg.Data, &data); err != nil {
		logger.ErrorJ(g.lc, fmt.Sprint("Wrong message DATA: ", err.Error()))
		return nil, nil
	}

	return msg, data
}

func (g *GetUserDataProto) MakeBody(recv getUserDataRecv.InternalMessage, result interface{}) interface{} {
	var respBody interface{}

	switch result.(type) {
	case getUserDataResp.Data:
		respBody = getUserDataResp.IMSuccess{
			ID:       recv.ID,
			GID:      recv.CID,
			CID:      recv.GID,
			Expires:  recv.Expires,
			Features: recv.Features,
			Locale:   recv.Locale,
			Domain:   recv.Domain,
			Event:    recv.Event,
			Data:     result.(getUserDataResp.Data),
			Saga: getUserDataResp.Saga{
				SID:     recv.Saga.SID,
				Name:    recv.Saga.Name,
				State:   recv.Saga.State,
				Expires: recv.Expires,
			},
			Src: getUserDataResp.Src{
				Name: recv.Src.Name,
				Version: getUserDataResp.Version{
					Major: recv.Src.Version.Major,
					Minor: recv.Src.Version.Minor,
					Patch: recv.Src.Version.Patch,
				},
			},
			ReplyTo: getUserDataResp.ReplyTo{
				Exchange:   recv.ReplyTo.Exchange,
				RoutingKey: recv.ReplyTo.RoutingKey,
				Event:      recv.ReplyTo.Event,
			},
			Tokens: getUserDataResp.Tokens{
				ID:     recv.Tokens.ID,
				Access: recv.Tokens.Access,
			},
		}
	case getUserDataResp.Errors:
		respBody = getUserDataResp.IMError{
			ID:       recv.ID,
			GID:      recv.CID,
			CID:      recv.GID,
			Expires:  recv.Expires,
			Features: recv.Features,
			Locale:   recv.Locale,
			Domain:   recv.Domain,
			Event:    recv.Event,
			Errors:   result.(getUserDataResp.Errors),
			Saga: getUserDataResp.Saga{
				SID:     recv.Saga.SID,
				Name:    recv.Saga.Name,
				State:   recv.Saga.State,
				Expires: recv.Expires,
			},
			Src: getUserDataResp.Src{
				Name: recv.Src.Name,
				Version: getUserDataResp.Version{
					Major: recv.Src.Version.Major,
					Minor: recv.Src.Version.Minor,
					Patch: recv.Src.Version.Patch,
				},
			},
			ReplyTo: getUserDataResp.ReplyTo{
				Exchange:   recv.ReplyTo.Exchange,
				RoutingKey: recv.ReplyTo.RoutingKey,
				Event:      recv.ReplyTo.Event,
			},
			Tokens: getUserDataResp.Tokens{
				ID:     recv.Tokens.ID,
				Access: recv.Tokens.Access,
			},
		}
	}

	return respBody
}
