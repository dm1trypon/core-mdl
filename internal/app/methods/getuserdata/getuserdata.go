package getuserdata

import (
	"encoding/json"
	"strconv"

	getUserDataRecv "github.com/dm1trypon/core-mdl/internal/app/proto/getuserdata/recieve"
	getUserDataResp "github.com/dm1trypon/core-mdl/internal/app/proto/getuserdata/response"
	"github.com/dm1trypon/core-mdl/internal/pkg/orm"
)

func (m *GetUserData) Create(ormInst *orm.ORM) *GetUserData {
	m = &GetUserData{
		lc:      "GET_USER_DATA",
		ormInst: ormInst,
	}

	return m
}

func (m *GetUserData) Run(data getUserDataRecv.Data) interface{} {
	errors := getUserDataResp.Errors{}

	var dataID string

	if err := json.Unmarshal(data.ID, &dataID); err != nil {
		errData := getUserDataResp.Error{
			Code:    -10050,
			Message: err.Error(),
			Data:    getUserDataResp.ErrorData{},
		}

		errors = append(errors, errData)

		return errors
	}

	id, err := strconv.Atoi(dataID)
	if err != nil {
		errData := getUserDataResp.Error{
			Code:    -10050,
			Message: err.Error(),
			Data:    getUserDataResp.ErrorData{},
		}

		errors = append(errors, errData)

		return errors
	}

	userData, ok := m.ormInst.GetUserData(id)
	if !ok {
		errData := getUserDataResp.Error{
			Code:    -10050,
			Message: "Error in query GetUserData",
			Data:    getUserDataResp.ErrorData{},
		}

		errors = append(errors, errData)

		return errors
	}

	if len(userData) < 1 {
		errData := getUserDataResp.Error{
			Code:    -10050,
			Message: "Empty result query GetUserData",
			Data:    getUserDataResp.ErrorData{},
		}

		errors = append(errors, errData)

		return errors
	}

	var username json.RawMessage

	if userData[0].Username.Valid {
		res, err := json.Marshal(userData[0].Username.String)
		if err != nil {
			username = json.RawMessage(nil)
		}

		username = json.RawMessage(res)
	} else {
		username = json.RawMessage(nil)
	}

	return getUserDataResp.Data{
		Username: username,
	}
}
