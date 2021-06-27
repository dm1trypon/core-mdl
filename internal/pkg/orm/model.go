package orm

import (
	"database/sql"

	"github.com/dm1trypon/db-mdl/dbpgconnector"
)

type ORM struct {
	lc           string
	dbPgConnInst *dbpgconnector.DBPGConnector
}

type GetUserData struct {
	Username sql.NullString `db:"username"`
}
