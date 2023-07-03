//
// Copyright (c) 2017 - 2021, © Badassops LLC
// All rights reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// * Proprietary and confidential *
//

package main

import (
	"fmt"
	"net/http"
	"os"

	"deployer.badassops.com/Api"
	"deployer.badassops.com/Config"
	"deployer.badassops.com/Initialize"
	"deployer.badassops.com/Logs"
	"deployer.badassops.com/Utils"
	"deployer.badassops.com/Variables"
)

func main() {
	// must run as root
	//Utils.IsRoot()

	// get given argument overwrite the values in the configuration file
	Initialize.InitArgs()

	// get value from the configuration file, default or overwritten
	Config.Init(Variables.GivenValues["configFile"])

	// initialize the logger system
	Logs.Init()

	// initialize the aoi system
	Api.Init()

	// install a signale handler so we capture issue if the application dies
	Utils.SignalHandler()

	// start the api server
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", Config.Server.Ip, Config.Server.Port), nil)
	if err != nil {
		//e := fmt.Sprintf("http.ListenAndServer failed `%s`, panic follows...", err.Error())
		//Logs.Error(errors.New(e))
		Utils.ExitIfError(err)
	}

	// should never be reached
	os.Exit(0)
}
