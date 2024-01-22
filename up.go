package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)

// Update updates axon by downloading the latest executable from github, and renaming the
// old executable to `axon-old` so that it can be deleted by `DeletePreviousInstallation`.
func update() {
	fmt.Println("Updating axon...")
	fmt.Println("Downloading the axon executable...")

	// * determining the os-specific url
	url := ""
	switch runtime.GOOS {
	case "windows":
		url = "https://github.com/shravanasati/axon/releases/latest/download/axon-windows-amd64.exe"
	case "linux":
		url = "https://github.com/shravanasati/axon/releases/latest/download/axon-linux-amd64"
	case "darwin":
		url = "https://github.com/shravanasati/axon/releases/latest/download/axon-darwin-amd64"
	default:
		fmt.Println("Your OS isn't supported by axon.")
		return
	}

	// * sending a request
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error: Unable to download the executable. Check your internet connection.")
		fmt.Println(err.Error())
		return
	}

	defer res.Body.Close()

	// * determining the executable path
	downloadPath, e := os.UserHomeDir()
	if e != nil {
		fmt.Println("Error: Unable to determine axon's location.")
		fmt.Println(e.Error())
		return
	}
	downloadPath += "/.axon/axon"
	if runtime.GOOS == "windows" {
		downloadPath += ".exe"
	}

	os.Rename(downloadPath, downloadPath+"-old")

	exe, er := os.Create(downloadPath)
	if er != nil {
		fmt.Println("Error: Unable to access file permissions.")
		fmt.Println(er.Error())
		return
	}
	defer exe.Close()

	// * writing the received content to the axon executable
	_, errr := io.Copy(exe, res.Body)
	if errr != nil {
		fmt.Println("Error: Unable to write the executable.")
		fmt.Println(errr.Error())
		return
	}

	// * performing an additional `chmod` utility for linux and mac
	if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		os.Chmod(downloadPath, 0755)
	}

	fmt.Println("axon was updated successfully.")
}

// DeletePreviousInstallation deletes previous installation if it exists.
func deletePreviousInstallation() {
	axonDir, _ := os.UserHomeDir()
	axonDir += "/.axon"

	files, _ := ioutil.ReadDir(axonDir)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), "-old") {
			// fmt.Println("found existsing installation")
			os.Remove(axonDir + "/" + f.Name())
		}
		// fmt.Println(f.Name())
	}
}
