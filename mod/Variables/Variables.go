//
// Copyright (c) 2017 - 2021, © Badassops LLC
// All rights reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// * Proprietary and confidential *
//
package Variables

import (
	"os"
	"path"
	"strconv"
	"time"
)

var (
	// varaibles start with a kapital letter are global
	now = time.Now()

	MyVersion   = "0.3"
	MyProgname  = path.Base(os.Args[0])
	myAuthor    = "Marc Krisnanto and Luc Suryo"
	myCopyright = "Copyright 2017 - " + strconv.Itoa(now.Year()) + " © Badassops LLC"
	myLicense   = "Badassops LLC proprietary ♥"
	myEmail     = "<marc@badassops.com> and <luc@badassops.com>"
	MyInfo      = MyProgname + " " + MyVersion + "\n" +
		myCopyright + "\nLicense: " + myLicense +
		"\nWritten by " + myAuthor + "\n" + 
        "Authors email: " + myEmail + "\n"
	MyTask = "A blue - green deployment API server"

	// defaults
	GivenValues map[string]string

	DefaultHome          = "/etc/deployer"
	DefaultConfigFile    = DefaultHome + "/config.json"
	DefaultLog           = "/var/log/deployer.log"
	DefaultLogMaxSize    = 512 // MB
	DefaultLogMaxBackups = 28  // count
	DefaultLogMaxAge     = 30  // days
)

func init() {
	// setup the default value, these are hardcoded.
	GivenValues = make(map[string]string)
	GivenValues["configFile"] = DefaultConfigFile
	GivenValues["logfile"] = DefaultLog
}
