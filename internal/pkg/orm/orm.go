package orm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dm1trypon/core-mdl/internal/pkg/orm/queries"
	"github.com/dm1trypon/db-mdl/dbpgconnector"
	logger "github.com/dm1trypon/easy-logger"
)

func (o *ORM) Create(dbPgConnInst *dbpgconnector.DBPGConnector) *ORM {
	o = &ORM{
		lc:           "ORM",
		dbPgConnInst: dbPgConnInst,
	}

	return o
}

func (o *ORM) AddUserData(id int, nickname string) bool {
	query := queries.AddUserData
	// Modify query
	query = strings.Replace(query, ":id", strconv.Itoa(id), -1)
	query = strings.Replace(query, ":nickname", nickname, -1)

	dbpgTools := o.dbPgConnInst.GetDBPGTools(0, false)
	if dbpgTools == nil {
		return false
	}

	_, ok := dbpgTools.Exec(query)

	return ok
}

func (o *ORM) GetUserData(id int) ([]GetUserData, bool) {
	query := queries.GetUserData
	// Modify query
	query = strings.Replace(query, ":id", strconv.Itoa(id), -1)

	dbpgTools := o.dbPgConnInst.GetDBPGTools(0, false)
	if dbpgTools == nil {
		return []GetUserData{}, false
	}

	rows, ok := dbpgTools.Query(query)
	if !ok {
		return []GetUserData{}, false
	}

	var result []GetUserData

	for rows.Next() {
		getUserData := GetUserData{}

		if err := rows.Scan(&getUserData.Username); err != nil {
			logger.ErrorJ(o.lc, fmt.Sprint("Error adding query's result to struct: ", err.Error()))
			continue
		}

		result = append(result, getUserData)
	}

	return result, true
}
