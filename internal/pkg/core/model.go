package core

import (
	"github.com/dm1trypon/core-mdl/internal/app/methods"
	"github.com/dm1trypon/core-mdl/internal/pkg/orm"
	"github.com/dm1trypon/db-mdl/dbpgconnector"
	"github.com/dm1trypon/rmq-mdl/rmqconnector"
)

type Core struct {
	lc           string
	rmqConnInst  *rmqconnector.RMQConnector
	dbPgConnInst *dbpgconnector.DBPGConnector
	ormInst      *orm.ORM
	methodsInst  *methods.Methods
}
