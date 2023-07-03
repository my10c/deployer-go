//
// Copyright (c) 2017 - 2021, © Badassops LLC
// All rights reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// * Proprietary and confidential *
//
package Api

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"deployer.badassops.com/Config"
	"deployer.badassops.com/Logs"
	"deployer.badassops.com/Msg"
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
		err := Msg.ENotFound
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// Always authenticate, i.e. the request must contain an Auth header
	// containing the token as set in the configuration.
	auth := r.Header.Get("Auth")
	if auth != Config.Api.Auth {
		err := Msg.EUnauthorized
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// Check whether the remote host is allowed access.
	if len(Config.Api.AclP) > 0 {
		ip := net.ParseIP(getIp(r))
		ok := false
		for _, v := range Config.Api.AclP {
			fmt.Println(v, v.Contains(ip))
			if v.Contains(ip) {
				ok = true
				break
			}
		}
		if !ok {
			err := Msg.EUnauthorized
			Logs.Response(r, err)
			w.WriteHeader(Msg.GetStatus(err))
			return
		}
	}

	// The path must start with the prefix as set in the configuration.
	path := r.URL.Path
	if !strings.HasPrefix(path, Config.Api.Prefix) {
		err := Msg.ENotFound
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// There must be at least one argument to pass to the script.
	if len(r.URL.Query()) < 1 {
		err := Msg.ENotFound
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// Since we only need the name, then strip it to get the command (script).
	path = strings.TrimPrefix(path, Config.Api.Prefix)
	cmd, exists := Config.Api.Cmds[path]
	if !exists {
		err := Msg.ENotFound
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}

	// WARNING We pass the arguments as is.
	// It is up to the script to do whatever it has to do with it.
	exe := exec.Command(cmd, r.URL.RawQuery)

	// If we're debugging, it's nice to print out the result of the script.
	// That is, if it does print out something.
	if Logs.Debug != nil {
		exe.Stdout = os.Stdout
	}

	err := exe.Run()
	if err != nil {
		if Logs.Response != nil {
			Logs.Response(r, err)
		}
		w.WriteHeader(Msg.GetStatus(err))
		return
	}
	if Logs.Response != nil {
		Logs.Response(r, nil, fmt.Sprintf("%s %s", cmd, r.URL.RawQuery))
	}
}

func Init() {
	Logs.AddPrefix("Api.")

	// This handles every URL.
	http.HandleFunc("/", handle)
}
