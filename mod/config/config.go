//
// Copyright (c) 2017 - 2022, © Badassops LLC
// All rights reserved.
//
// Release under the BSD 3-Clause License
// https://opensource.org/licenses/BSD-3-Clause ♥
//

package config

import (
	"encoding/json"
	"net"
	"os"

	"utils"
)

// ServerT is the server configuration.
type ServerT struct {
	Port int    `json:"port"` // server will listen to this port
	Ip   string `json:"ip"`   // server IP address
}

// LogsT is the Logs module configuration.
type LogsT struct {
	FilePath       string `json:"logfile"`        // Absolute path of the log file.
	FileMaxSize    int    `json:"maxsize"`        // Maximum size of the log file before rotated (MB).
	FileMaxAge     int    `json:"maxage"`         // Maximum file age before rotated (days).
	FileMaxBackups int    `json:"maxbackups"`     // Maximum number of backups.
	FileCompress   bool   `json:"compress"`       // Whether to compress rotated files.
	FileLocalTime  bool   `json:"localtime"`      // Whether to use server localtime (false means UTC).
	Debug          bool   `json:"debug"`          // Whether to log debug messages.
	Info           bool   `json:"info"`           // Whether to log informational messages.
	ResponseOK     bool   `json:"response-ok"`    // Whether to log request OK responses.
	ResponseError  bool   `json:"response-error"` // Whether to log request error responses.
}

// ApiT is the Api module configuration.
type ApiT struct {
	Prefix string            `json:"prefix"` // api URL prefix, e.g. "/api/v1/"
	Auth   string            `json:"auth"`   // authentication token, e.g. "QBiaxVGbYqofjAQVmk7qAmiI3JlrC1cOFWgpJVwj"
	Acl    []string          `json:"acl"`    // allowed remoted IP addresses
	Cmds   map[string]string `json:"cmds"`   // list of commands, e.g. "blue":"/usr/local/sbin/api/blue"
	AclP   []*net.IPNet      // parsed Acl
}

var (
	// Server configuration.
	Server ServerT

	// Logs configuration.
	Logs LogsT

	// Api configuration.
	Api ApiT
)

func Init(path string) {
	var conf struct {
		Server ServerT `json:"server"`
		Logs   LogsT   `json:"logs"`
		Api    ApiT    `json:"api"`
	}

	f, err := os.Open(path)
	if err != nil {
		//log.Printf("Config.Init %s: %s", path, err.Error())
		//panic(err)
		utils.ExitIfError(err)
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&conf)
	if err != nil {
		//log.Printf("Config.Init %s: %s", path, err.Error())
		//panic(err)
		utils.ExitIfError(err)
	}
	Server = conf.Server
	Logs = conf.Logs
	Api = conf.Api

	// parse the CIDR notations for quicker comparison by Api.go
	// also check whether they are valid
	for _, v := range Api.Acl {
		_, ipv4Net, err := net.ParseCIDR(v)
		if err != nil {
			//log.Printf("Config.Init %s: %s", v, err)
			//panic(err)
			utils.ExitIfError(err)
		}
		Api.AclP = append(Api.AclP, ipv4Net)
	}
}
