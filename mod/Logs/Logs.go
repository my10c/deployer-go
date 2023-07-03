//
// Copyright (c) 2017 - 2021, © Badassops LLC
// All rights reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// * Proprietary and confidential *
//
package Logs

import (
	"log"
	"net"
	"net/http"
	"runtime"
	"strconv"
	"strings"

	"deployer.badassops.com/Config"
	"deployer.badassops.com/Msg"
	"deployer.badassops.com/Variables"

	"gopkg.in/natefinch/lumberjack.v2"
)

// AddPrefix adds a function name prefix for logging.
// When debugging, Logs will include the function names in the call stack up to
// the function which name starts with a prefix registerd with this function.
//
//  Example
//
//    package foo
///    func helper_bar () {
//        if Logs.Info != nil {
//            Logs.Info("bla bla")  //--> logs "I (foo.say_boo < foo.helper_bar) bla bla"
//         }
//    }
///    func say_boo () {
//       helper_bar()
//    }
///    func int() {
//        Logs.AddPrefix("foo.say_")
//    }
//
func AddPrefix(s ...string) {
	prefixes = append(prefixes, s...)
}

var prefixes = []string{"main."}

func _funcName(depth int) string {
	pc, _, _, ok := runtime.Caller(depth)
	info := runtime.FuncForPC(pc)
	if ok && info != nil {
		fname := runtime.FuncForPC(pc).Name()
		pos := strings.LastIndex(fname, "/")
		if pos > 1 {
			fname = fname[pos+1:]
		}
		fname = strings.TrimSuffix(fname, ".0")
		for _, prefix := range prefixes {
			if strings.HasPrefix(fname, prefix) {
				for depth > 3 {
					depth--
					pc, _, _, ok := runtime.Caller(depth)
					info := runtime.FuncForPC(pc)
					if !ok || info == nil {
						break
					}
					f := runtime.FuncForPC(pc).Name()
					pos := strings.LastIndex(f, "/")
					if pos > 1 {
						f = f[pos+1:]
					}
					if f == "" {
						break
					}
					fname = fname + " > " + f
				}
				return "(" + fname + ")"
			}
		}
	}
	return ""
}

// Get the function name in the call stack that matches a prefix.
func funcName() string {
	flast := _funcName(3)
	for i := 4; i < 7; i++ {
		fname := _funcName(i)
		if fname != "" {
			return fname
		}
	}
	return flast
}

var removeLines = strings.NewReplacer(
	"\r\n", "\\r\\n",
	"\r", "\\r",
	"\n", "\\n")

func tidy(s string) string {
	return strings.TrimSpace(removeLines.Replace(s))
}

// Debug logs a debug message.
// Caller must check whether Debug is nil before calling it. The idea; no time
// and memory is wasted building a message string (e.g. using fmt.Sprintf) that
// will simply not be used.
// Example
//     if Logs.Debug != nil {
//         Logs.Debug(fmt.Sprintf("Foo %s", bar))
//     }
var Debug = debug

var debug = func(msg string) {
	var b strings.Builder
	b.Grow(128)
	b.WriteString("D ")
	b.WriteString(funcName())
	b.WriteString(" ")
	b.WriteString(tidy(msg))
	log.Println(b.String())
}

// Info logs a system informational message.
// Caller must check whether Info is nil before calling it. The idea; no time
// and memory is wasted building a message string (e.g. using fmt.Sprintf) that
// will simply not be used.
// Example
//     if Logs.Info != nil {
//         Logs.Info(fmt.Sprintf("Foo %s", bar))
//     }
var Info = info

var info = func(msg string) {
	var b strings.Builder
	b.Grow(128)
	b.WriteString("I ")
	b.WriteString(funcName())
	b.WriteString(" ")
	b.WriteString(tidy(msg))
	log.Println(b.String())
}

// Error logs a system error message.
// Example
//     err := doSomething()
//     if err != nil {
//         Logs.Error(err)
//     }
func Error(err error) {
	var b strings.Builder
	b.Grow(128)
	b.WriteString("E ")
	b.WriteString(funcName())
	b.WriteString(" ")
	b.WriteString(tidy(err.Error()))
	log.Println(b.String())
}

// Whether to log a HTTP 200 response
var responseOK = true

// Whether to log a non HTTP 200 response
var responseError = true

// Don't need the port part of RemoteAddr.
func ip(r *http.Request) string {
	a := r.RemoteAddr
	host, _, err := net.SplitHostPort(a)
	if err == nil {
		return host
	}
	return a
}

// Response logs a response for a HTTP request.
var Response = response

var response = func(r *http.Request, err error, args ...string) {
	status := 200
	msg := "OK"
	if err != nil {
		status = Msg.GetStatus(err)
		if len(args) > 0 {
			msg = tidy(args[0])
		} else {
			msg = tidy(err.Error())
		}
	} else {
		if len(args) > 0 {
			msg = tidy(args[0])
		}
	}
	var b strings.Builder
	b.Grow(128)
	b.WriteString(strconv.Itoa(status))
	b.WriteString(" ")
	b.WriteString(r.Method)
	b.WriteString(" ")
	b.WriteString(r.URL.Path)
	b.WriteString(" ")
	b.WriteString(ip(r))
	b.WriteString(" ")
	b.WriteString(msg)
	log.Println(b.String())
}

func Init() {
	AddPrefix("Logs.")

	if !Config.Logs.Info {
		Info = nil
	} else {
		Info = info
	}

	if !Config.Logs.Debug {
		Debug = nil
	} else {
		Debug = debug
	}

	if Config.Logs.ResponseOK {
		responseOK = true
	} else {
		responseOK = false
	}
	if Config.Logs.ResponseError {
		responseError = true
	} else {
		responseError = false
	}

	if !responseOK && !responseError {
		Response = nil
	} else {
		Response = response
	}

	// Override values if given via the command line arguments
	x := Variables.GivenValues["maxsize"] // = strconv.Itoa(DefaultLogMaxSize)
	if x != "" {
		Config.Logs.FileMaxSize, _ = strconv.Atoi(x)
	}
	x = Variables.GivenValues["maxbackups"] // = strconv.Itoa(DefaultLogMaxSize)
	if x != "" {
		Config.Logs.FileMaxBackups, _ = strconv.Atoi(x)
	}
	x = Variables.GivenValues["maxage"] // = strconv.Itoa(DefaultLogMaxSize)
	if x != "" {
		Config.Logs.FileMaxAge, _ = strconv.Atoi(x)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(&lumberjack.Logger{
		Filename:   Config.Logs.FilePath,
		MaxSize:    Config.Logs.FileMaxSize,
		MaxBackups: Config.Logs.FileMaxBackups,
		MaxAge:     Config.Logs.FileMaxAge,
		Compress:   true,
	})
}
