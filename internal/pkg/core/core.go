package core

import (
	"time"

	"github.com/dm1trypon/core-mdl/internal/app/methods"
	"github.com/dm1trypon/core-mdl/internal/pkg/orm"
	"github.com/dm1trypon/db-mdl/dbpgconnector"
	"github.com/dm1trypon/rmq-mdl/rmqconnector"
)

func (c *Core) Create() *Core {
	c = &Core{
		lc:           "CORE",
		rmqConnInst:  nil,
		dbPgConnInst: nil,
		ormInst:      nil,
		methodsInst:  nil,
	}

	return c
}

func (c *Core) Run() {
	c.startDBPGModule()
	c.startRMQModule()
	c.ormInst = new(orm.ORM).Create(c.dbPgConnInst)
	c.methodsInst = new(methods.Methods).Create(c.ormInst)
}

func (c *Core) startRMQModule() {
	c.rmqConnInst = new(rmqconnector.RMQConnector).Create()

	cfg := rmqconnector.Config{
		Username:             "guest",
		Password:             "guest",
		Host:                 "127.0.0.1",
		Port:                 5672,
		TLS:                  true,
		ReconnectionInterval: time.Duration(2 * time.Second),
		Certs:                rmqconnector.Certs{},
		Events: []rmqconnector.Event{
			{
				Consuming: true,
				Kind:      "reciever",
				Exchange:  "reciever",
				Queue:     "reciever_queue",
			},
			{
				Consuming: false,
				Kind:      "respondent",
				Exchange:  "respondent",
				Queue:     "",
			},
		},
	}

	c.rmqConnInst.SetConfig(cfg)
	go c.rmqConnInst.Run()
	<-c.rmqConnInst.GetChConnected()
	go c.rmqHandler()
}

func (c *Core) rmqHandler() {
	for {
		recvMsg := <-c.rmqConnInst.GetChNextMsg()
		respMsg := c.methodsInst.NextMessage([]byte(recvMsg.Body))

		if respMsg == nil {
			continue
		}

		c.rmqConnInst.Publish(respMsg, "respondent", "application/json")
	}
}

func (c *Core) startDBPGModule() {
	c.dbPgConnInst = new(dbpgconnector.DBPGConnector).Create()

	cfg := dbpgconnector.Config{
		Username:             "postgres",
		Password:             "mpassword",
		Host:                 "localhost",
		Port:                 5432,
		DbName:               "db_game",
		SSLMode:              0,
		ConnectTimeout:       10,
		PingInterval:         2 * time.Second,
		ReconnectionInterval: 2 * time.Second,
		Certs:                dbpgconnector.Certs{},
	}
	c.dbPgConnInst.SetConfig(cfg)

	go c.dbPgConnInst.Run()
	<-c.dbPgConnInst.GetChConnected()

	// Configuring the access level and transaction isolation level.
	settings := map[uint8]bool{
		0: false,
	}
	c.dbPgConnInst.SetDBPGToolsList(settings)
}
