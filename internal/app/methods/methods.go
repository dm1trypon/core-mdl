package methods

import (
	"github.com/dm1trypon/core-mdl/internal/app/methods/adduserdata"
	"github.com/dm1trypon/core-mdl/internal/app/methods/getuserdata"
	"github.com/dm1trypon/core-mdl/internal/app/proto"
	addUserDataRecv "github.com/dm1trypon/core-mdl/internal/app/proto/adduserdata/recieve"
	getUserDataRecv "github.com/dm1trypon/core-mdl/internal/app/proto/getuserdata/recieve"
	"github.com/dm1trypon/core-mdl/internal/pkg/orm"
)

func (m *Methods) Create(ormInst *orm.ORM) *Methods {
	m = &Methods{
		lc:              "METHODS",
		protoInst:       new(proto.Proto).Create(),
		getUserDataInst: new(getuserdata.GetUserData).Create(ormInst),
		addUserDataInst: new(adduserdata.AddUserData).Create(ormInst),
	}

	return m
}

func (m *Methods) NextMessage(body []byte) []byte {
	msg, data := m.protoInst.Parse(body)
	if msg == nil || data == nil {
		return nil
	}

	switch data.(type) {
	case getUserDataRecv.Data:
		return m.protoInst.Pack(msg, m.getUserDataInst.Run(data.(getUserDataRecv.Data)))
	case addUserDataRecv.Data:
		return m.protoInst.Pack(msg, m.addUserDataInst.Run(data.(addUserDataRecv.Data)))
	}

	return m.protoInst.Pack(msg, data)
}
