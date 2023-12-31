//
// Copyright (c) 2017 - 2022, © Badassops LLC
// All rights reserved.
//
// Release under the BSD 3-Clause License
// https://opensource.org/licenses/BSD-3-Clause ♥
//

package vars

import (
	"os"
	"path"
	"strconv"
	"time"
)

var (
	// varaibles start with a kapital letter are global
	now = time.Now()

	MyVersion   = "0.5"
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
