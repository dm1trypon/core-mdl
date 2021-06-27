package adduserdata

import (
	"strconv"

	addUserDataRecv "github.com/dm1trypon/core-mdl/internal/app/proto/adduserdata/recieve"
	addUserDataResp "github.com/dm1trypon/core-mdl/internal/app/proto/adduserdata/response"
	"github.com/dm1trypon/core-mdl/internal/pkg/orm"
)

func (m *AddUserData) Create(ormInst *orm.ORM) *AddUserData {
	m = &AddUserData{
		lc:      "ADD_USER_DATA",
		ormInst: ormInst,
	}

	return m
}

func (m *AddUserData) Run(data addUserDataRecv.Data) interface{} {
	id, err := strconv.Atoi(data.ID)
	if err != nil {
		return m.makeError(-10050, err.Error())
	}

	ok := m.ormInst.AddUserData(id, data.Nickname)
	if !ok {
		return m.makeError(-10050, "Error in query AddUserData")
	}

	return addUserDataResp.Data{
		Success: ok,
	}
}

func (m *AddUserData) makeError(code int16, err string) addUserDataResp.Errors {
	errors := addUserDataResp.Errors{}

	errData := addUserDataResp.Error{
		Code:    code,
		Message: err,
		Data:    addUserDataResp.ErrorData{},
	}

	errors = append(errors, errData)
	return errors
}
