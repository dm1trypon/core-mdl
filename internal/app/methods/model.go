package methods

import (
	"github.com/dm1trypon/core-mdl/internal/app/methods/adduserdata"
	"github.com/dm1trypon/core-mdl/internal/app/methods/getuserdata"
	"github.com/dm1trypon/core-mdl/internal/app/proto"
)

type Methods struct {
	lc              string
	protoInst       *proto.Proto
	getUserDataInst *getuserdata.GetUserData
	addUserDataInst *adduserdata.AddUserData
}
