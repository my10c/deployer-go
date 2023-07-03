//
// Copyright (c) 2017 - 2021, © Badassops LLC
// All rights reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// * Proprietary and confidential *
//
package Initialize

import (
	"fmt"
	"os"
	"strconv"

	"deployer.badassops.com/Help"
	"deployer.badassops.com/Utils"
	"deployer.badassops.com/Variables"

	"github.com/akamensky/argparse"
)

// Function to process the given args
func InitArgs() {
	parser := argparse.NewParser(Variables.MyProgname, Variables.MyTask)

	configFile := parser.String("c", "configFile",
		&argparse.Options{
			Required: false,
			Help:     "Path to the configuration file",
			Default:  Variables.DefaultConfigFile,
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
			Default:  Variables.DefaultLog,
		})

	logMaxSize := parser.String("M", "logMaxSize",
		&argparse.Options{
			Required: false,
			Help:     "Max size of the log file (MB). Default: " + strconv.Itoa(Variables.DefaultLogMaxSize),
			//Default:  strconv.Itoa(Variables.DefaultLogMaxSize),
		})

	logMaxBackups := parser.String("B", "logMaxBackups",
		&argparse.Options{
			Required: false,
			Help:     "Max log file count. Default: " + strconv.Itoa(Variables.DefaultLogMaxBackups),
			//Default:  strconv.Itoa(Variables.DefaultLogMaxBackups),
		})

	logMaxAge := parser.String("A", "logMaxAge",
		&argparse.Options{
			Required: false,
			Help:     "Max days to keep a log file. Default: " + strconv.Itoa(Variables.DefaultLogMaxAge),
			//Default:  strconv.Itoa(Variables.DefaultLogMaxAge),
		})

	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *showSetup {
		Help.HelpSetup()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("%s\n", Variables.MyVersion)
		os.Exit(0)
	}

	if *showInfo {
		fmt.Printf("%s\n", Variables.MyInfo)
		fmt.Printf("%s\n", Variables.MyTask)
		os.Exit(0)
	}

	if configFile != nil {
		Variables.GivenValues["configFile"] = fmt.Sprintf("%s", *configFile)
	}

	//if *logFile != "" {
	Variables.GivenValues["logfile"] = fmt.Sprintf("%s", *logFile)
	//}

	if *logMaxSize != "" {
		Variables.GivenValues["maxsize"] = fmt.Sprintf("%s", *logMaxSize)
	}

	if *logMaxBackups != "" {
		Variables.GivenValues["maxbackups"] = fmt.Sprintf("%s", *logMaxBackups)
	}

	if *logMaxAge != "" {
		Variables.GivenValues["maxage"] = fmt.Sprintf("%s", *logMaxAge)
	}

	if !Utils.CheckFileExist(Variables.GivenValues["configFile"]) {
		fmt.Printf("Configuration file %s, does not exist\n", Variables.GivenValues["configFile"])
		os.Exit(1)
	}

	if !Utils.CheckFileExist(Variables.GivenValues["logfile"]) {
		os.Exit(1)
	}
}
