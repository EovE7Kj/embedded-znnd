package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"encoding/base64"
	"compress/gzip"
	"znn-uk/embedded" 
)

const (
  dataDir = "/root/znn"
)

func main() {

	err := os.MkdirAll(dataDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		os.Exit(1)
	}

	decodedBinary, err := base64.StdEncoding.DecodeString(strings.TrimSpace(embedded.EmbeddedBinary))
	if err != nil {
		fmt.Printf("Error decoding base64 string: %v\n", err)
		os.Exit(1)
	}

	gzipReader, err := gzip.NewReader(strings.NewReader(string(decodedBinary)))
	if err != nil {
		fmt.Printf("Error creating gzip reader: %v\n", err)
		os.Exit(1)
	}
	defer gzipReader.Close()

	tempFile, err := os.CreateTemp("", "znnd-amd64-*")
	if err != nil {
		fmt.Printf("Error creating temporary file: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, gzipReader)
	if err != nil {
		fmt.Printf("Error writing to temporary file: %v\n", err)
		os.Exit(1)
	}

	err = tempFile.Chmod(0755)
	if err != nil {
		fmt.Printf("Error setting executable permission: %v\n", err)
		os.Exit(1)
	}

	err = tempFile.Close()
	if err != nil {
		fmt.Printf("Error closing temporary file: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(tempFile.Name(), "--data", dataDir)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
}
