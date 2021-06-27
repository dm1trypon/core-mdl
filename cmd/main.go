package main

import (
	"github.com/dm1trypon/core-mdl/internal/pkg/core"
	logger "github.com/dm1trypon/easy-logger"
)

// LC - logging's category
const LC = "MAIN"

func main() {
	logCfg := logger.Cfg{
		AppName: "CORE_MDL",
		LogPath: "",
		Level:   0,
	}

	logger.SetConfig(logCfg)

	logger.InfoJ(LC, "STARTING SERVICE")

	stop := make(chan bool)

	coreInst := new(core.Core).Create()
	coreInst.Run()

	<-stop
}
