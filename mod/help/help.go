//
// Copyright (c) 2017 - 2021, © Badassops LLC
// All rights reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// * Proprietary and confidential *
//
package help

import (
	"fmt"
	"os"

	"vars"
)

func HelpSetup() {
	fmt.Printf("%s", vars.MyInfo)
	fmt.Print(`
Configuration file must be valid json files with this structure:
{
	"server": {
		"port": 9091,           The port to listen to.
		"ip": "172.16.240.151"  IP to bind to.
	},
	"logs": {
		"logfile": "/var/log/deployer.log",  Absolute path of the log file.
		"maxsize": 512,                      Maximum size of the log file before rotated (MB).
		"maxage": 30,                        Maximum file age before rotated (days).
		"maxbackups": 28                     Maximum number of backups.
		"debug": false,                      Whether to log debug messages.
		"info": true,                        Whether to log informational messages.
		"response-ok": true,                 Whether to log successful requests.
		"response-error": true               Whether to log error requests.
	},
	"api": {
		"prefix": "/api/v1/",                        URI prefix for the APIs.
		"auth": "QBiaxVGbYqofjAQj"                   Authorization token.
		"acl": ["127.0.0.0/16", "172.16.240.0/24"],  IPs allowed to connect in CIDR notation.
		"cmds": {                                    List of API name and script to execute.
			"blue": "/var/deployer/blue.sh",
			"green": "/var/deployer/green.sh"
		}
	}
}

Example on how to access the server:

	curl --header \"Auth: QBiaxVGbYqofjAQj\" http://127.0.0.1:9091/api/v1/blue?foo:bar

`)
	os.Exit(0)
}
