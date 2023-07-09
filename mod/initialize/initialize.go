//
// Copyright (c) 2017 - 2022, © Badassops LLC
// All rights reserved.
//
// Release under the BSD 3-Clause License
// https://opensource.org/licenses/BSD-3-Clause ♥
//

package initialize

import (
	"fmt"
	"os"
	"strconv"

	"help"
	"utils"
	"vars"

	"github.com/akamensky/argparse"
)

// Function to process the given args
func InitArgs() {
	parser := argparse.NewParser(vars.MyProgname, vars.MyTask)

	configFile := parser.String("c", "configFile",
		&argparse.Options{
			Required: false,
			Help:     "Path to the configuration file",
			Default:  vars.DefaultConfigFile,
		})

	showSetup := parser.Flag("S", "showconfig",
		&argparse.Options{
			Required: false,
			Help:     "Show configuration example",
			Default:  false,
		})

	showVersion := parser.Flag("v", "version",
		&argparse.Options{
			Required: false,
			Help:     "Show version",
			Default:  false,
		})

	showInfo := parser.Flag("i", "info",
		&argparse.Options{
			Required: false,
			Help:     "Show information",
			Default:  false,
		})

	logFile := parser.String("l", "logFile",
		&argparse.Options{
			Required: false,
			Help:     "Path to the log file",
			Default:  vars.DefaultLog,
		})

	logMaxSize := parser.String("M", "logMaxSize",
		&argparse.Options{
			Required: false,
			Help:     "Max size of the log file (MB). Default: " + strconv.Itoa(vars.DefaultLogMaxSize),
			//Default:  strconv.Itoa(vars.DefaultLogMaxSize),
		})

	logMaxBackups := parser.String("B", "logMaxBackups",
		&argparse.Options{
			Required: false,
			Help:     "Max log file count. Default: " + strconv.Itoa(vars.DefaultLogMaxBackups),
			//Default:  strconv.Itoa(vars.DefaultLogMaxBackups),
		})

	logMaxAge := parser.String("A", "logMaxAge",
		&argparse.Options{
			Required: false,
			Help:     "Max days to keep a log file. Default: " + strconv.Itoa(vars.DefaultLogMaxAge),
			//Default:  strconv.Itoa(vars.DefaultLogMaxAge),
		})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *showSetup {
		help.HelpSetup()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("%s\n", vars.MyVersion)
		os.Exit(0)
	}

	if *showInfo {
		fmt.Printf("%s\n", vars.MyInfo)
		fmt.Printf("%s\n", vars.MyTask)
		os.Exit(0)
	}

	if configFile != nil {
		vars.GivenValues["configFile"] = fmt.Sprintf("%s", *configFile)
	}

	//if *logFile != "" {
	vars.GivenValues["logfile"] = fmt.Sprintf("%s", *logFile)
	//}

	if *logMaxSize != "" {
		vars.GivenValues["maxsize"] = fmt.Sprintf("%s", *logMaxSize)
	}

	if *logMaxBackups != "" {
		vars.GivenValues["maxbackups"] = fmt.Sprintf("%s", *logMaxBackups)
	}

	if *logMaxAge != "" {
		vars.GivenValues["maxage"] = fmt.Sprintf("%s", *logMaxAge)
	}

	if !utils.CheckFileExist(vars.GivenValues["configFile"]) {
		fmt.Printf("Configuration file %s, does not exist\n", vars.GivenValues["configFile"])
		os.Exit(1)
	}

	if !utils.CheckFileExist(vars.GivenValues["logfile"]) {
		os.Exit(1)
	}
}
