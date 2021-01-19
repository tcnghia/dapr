// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

// This binary can be optimized further to be a tiny container image
// that can serve as a very fast-loading Init container.
//
// For now, in the PoC, I don't mind yet about the one restart that
// happens. It may even be faster than an Init container.
//
// So, for the sake of hackery, I bundle this into the sidecar binary
// itself. Yuck.

const (
	varDapr = "/var/dapr"
	// Ideally this should be a socket, and we can use HTTP over
	// Unix socket to communicate intelligently with the sidecar.
	//
	// But for now this is a dumb file.
	readyPath = varDapr + "/ready"
)

func launchCmd() error {
	fmt.Println("--- APP LAUNCHER log ---")
	log.Info("Waiting for Dapr to be ready...")
	for {
		f, err := os.Open(readyPath)
		if err == nil {
			f.Close()
			message, _ := ioutil.ReadFile(readyPath)
			log.Info("Received message from sidecar:", string(message))
			break
		}
		fmt.Println(".")
		time.Sleep(100 * time.Millisecond)
	}
	log.Info("Dapr is ready")
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Infof("Launching container with command %+v\n", cmd)
	fmt.Println("\n--- APP log --- ")
	return cmd.Run()
}
