//
// Copyright (c) 2017 - 2021, © Badassops LLC
// All rights reserved.
//
// Unauthorized copying of this file, via any medium is strictly prohibited
// * Proprietary and confidential *
//
package utils

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"vars"
)

// Function to exit if an error occured
func ExitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+fmt.Sprint(err))
		log.Printf("-< %s >-\n", fmt.Sprint(err))
		os.Exit(1)
	}
}

// Function to exit if pointer is nill
func ExitIfNill(ptr interface{}) {
	if ptr == nil {
		fmt.Fprintln(os.Stderr, "Error: got a nil pointer.")
		log.Printf("-< Error: got a nil pointer. >-\n")
		os.Exit(1)
	}
}

// Function to print the given message to stdout and log file
func StdOutAndLog(message string) {
	fmt.Printf("-< %s >-\n", message)
	log.Printf("-< %s >-\n", message)
	return
}

// Function to check if the user that runs the app is root
func IsRoot() {
	if os.Geteuid() != 0 {
		StdOutAndLog(fmt.Sprintf("%s must be run as root.", vars.MyProgname))
		os.Exit(1)
	}
}

// Function to log any reveived signal
func SignalHandler() {
	interrupt := make(chan os.Signal, 1)
	// we handle only these signal: SIGINT(2) - SIGTRAP(5) - SIGKILL(9) - SIGTERM(15)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTRAP, syscall.SIGKILL, syscall.SIGTERM)
	go func() {
		sigId := <-interrupt
		StdOutAndLog(fmt.Sprintf("received %v %d", sigId, sigId))
		os.Exit(0)
	}()
}

// Function to write a log if debug was enabled
func WriteDebug(debug string, messsage string) {
	debugMode, err := strconv.ParseBool(debug)
	if err != nil {
		return
	}
	if debugMode == true {
		log.Printf("Debug -< %s >-\n", messsage)
	}
	return
}

// Function check if file exist, it must be a directory and the parent directory misty exist
func CheckFileExist(fqdn string) bool {
	// can not be under /,
	if strings.Count(filepath.Clean(fqdn), "/") == 1 {
		fmt.Printf("Errored, file can not be under root (/), %s\n", filepath.Clean(fqdn))
		return false
	}

	// check file exist
	fileInfo, err := os.Stat(fqdn)
	if err != nil {
		// check if the parent directoy exist, in case the file will be created by a differen process
		dirPath, _ := path.Split(fqdn)
		// fmt.Printf(" %s <-> %s\n", dirPath, fileName)
		if _, err := os.Stat(dirPath); err != nil {
			fmt.Printf("Errored, parent directory %s does not exist\n", dirPath)
			return false
		}
		return true
	}

	// check is not a directory
	if fileInfo.IsDir() {
		fmt.Printf("Errored, given file name is a directory, %s\n")
		return false
	}
	return true
}
