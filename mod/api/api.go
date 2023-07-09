//
// Copyright (c) 2017 - 2022, © Badassops LLC
// All rights reserved.
//
// Release under the BSD 3-Clause License
// https://opensource.org/licenses/BSD-3-Clause ♥
//

package api

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"config"
	"logs"
	"msg"
)

func getIp(r *http.Request) string {
	a := r.RemoteAddr
	host, _, err := net.SplitHostPort(a)
	if err == nil {
		return host
	}
	return a
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Make sure the method is GET.
	if r.Method != "GET" {
		err := msg.ENotFound
		if logs.Response != nil {
			logs.Response(r, err)
		}
		w.WriteHeader(msg.GetStatus(err))
		return
	}

	// Always authenticate, i.e. the request must contain an Auth header
	// containing the token as set in the configuration.
	auth := r.Header.Get("Auth")
	if auth != config.Api.Auth {
		err := msg.EUnauthorized
		if logs.Response != nil {
			logs.Response(r, err)
		}
		w.WriteHeader(msg.GetStatus(err))
		return
	}

	// Check whether the remote host is allowed access.
	if len(config.Api.AclP) > 0 {
		ip := net.ParseIP(getIp(r))
		ok := false
		for _, v := range config.Api.AclP {
			fmt.Println(v, v.Contains(ip))
			if v.Contains(ip) {
				ok = true
				break
			}
		}
		if !ok {
			err := msg.EUnauthorized
			logs.Response(r, err)
			w.WriteHeader(msg.GetStatus(err))
			return
		}
	}

	// The path must start with the prefix as set in the configuration.
	path := r.URL.Path
	if !strings.HasPrefix(path, config.Api.Prefix) {
		err := msg.ENotFound
		if logs.Response != nil {
			logs.Response(r, err)
		}
		w.WriteHeader(msg.GetStatus(err))
		return
	}

	// There must be at least one argument to pass to the script.
	if len(r.URL.Query()) < 1 {
		err := msg.ENotFound
		if logs.Response != nil {
			logs.Response(r, err)
		}
		w.WriteHeader(msg.GetStatus(err))
		return
	}

	// Since we only need the name, then strip it to get the command (script).
	path = strings.TrimPrefix(path, config.Api.Prefix)
	cmd, exists := config.Api.Cmds[path]
	if !exists {
		err := msg.ENotFound
		if logs.Response != nil {
			logs.Response(r, err)
		}
		w.WriteHeader(msg.GetStatus(err))
		return
	}

	// WARNING We pass the arguments as is.
	// It is up to the script to do whatever it has to do with it.
	exe := exec.Command(cmd, r.URL.RawQuery)

	// If we're debugging, it's nice to print out the result of the script.
	// That is, if it does print out something.
	if logs.Debug != nil {
		exe.Stdout = os.Stdout
	}

	err := exe.Run()
	if err != nil {
		if logs.Response != nil {
			logs.Response(r, err)
		}
		w.WriteHeader(msg.GetStatus(err))
		return
	}
	if logs.Response != nil {
		logs.Response(r, nil, fmt.Sprintf("%s %s", cmd, r.URL.RawQuery))
	}
}

func Init() {
	logs.AddPrefix("Api.")

	// This handles every URL.
	http.HandleFunc("/", handle)
}
