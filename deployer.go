//
// Copyright (c) 2017 - 2022, © Badassops LLC
// All rights reserved.
//
// Release under the BSD 3-Clause License
// https://opensource.org/licenses/BSD-3-Clause ♥
//

package main

import (
	"fmt"
	"net/http"
	"os"

	"api"
	"config"
	"initialize"
	"logs"
	"utils"
	"vars"
)

func main() {
	// must run as root
	//utils.IsRoot()

	// get given argument overwrite the values in the configuration file
	initialize.InitArgs()

	// get value from the configuration file, default or overwritten
	config.Init(vars.GivenValues["configFile"])

	// initialize the logger system
	logs.Init()

	// initialize the aoi system
	api.Init()

	// install a signale handler so we capture issue if the application dies
	utils.SignalHandler()

	// start the api server
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Server.Ip, config.Server.Port), nil)
	if err != nil {
		//e := fmt.Sprintf("http.ListenAndServer failed `%s`, panic follows...", err.Error())
		//logs.Error(errors.New(e))
		utils.ExitIfError(err)
	}

	// should never be reached
	os.Exit(0)
}
