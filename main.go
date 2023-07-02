package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	architecture := flag.String("a", runtime.GOARCH, "Set architecture")
	platform := flag.String("p", runtime.GOOS, "Set platform")
	fileOutput := flag.String("o", fmt.Sprintf("build%s", PATH_SEPARATOR), "Set output file locate")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	var fileExtension string

	switch *platform {
	case "windows":
		fileExtension = ".exe"
	default:
		fileExtension = ""
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	slicedPath := strings.Split(path, PATH_SEPARATOR)
	projectName := slicedPath[len(slicedPath)-1]

	var outputFileName string
	if strings.HasSuffix(*fileOutput, "\\") || strings.HasSuffix(*fileOutput, "/") {
		outputFileName = fmt.Sprintf("%[1]s_%[2]s-%[3]s%[4]s", projectName, *platform, *architecture, fileExtension)
	} else {
		outputFileName = ""
	}

	buildCommand := fmt.Sprintf("go build -o %[1]s%[3]s%[2]s%[4]s", path, *fileOutput, PATH_SEPARATOR, outputFileName)

	if envErr := os.Setenv("GOARCH", *architecture); envErr != nil {
		fmt.Println(err)
	}

	if envErr := os.Setenv("GOOS", *platform); envErr != nil {
		fmt.Println(err)
	}

	fmt.Printf("Building %[1]s for %[2]s/%[3]s...\n\n", projectName, *platform, *architecture)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", buildCommand)
	default:
		cmd = exec.Command("sh", "-c", buildCommand)
	}

	var stderr bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr
	if cmdErr := cmd.Run(); cmdErr != nil {
		fmt.Printf("%s\n%s", cmdErr.Error(), stderr.String())
	} else {
		fmt.Printf("Build completed!\nFile saved at %[1]s%[2]s!\n", *fileOutput, outputFileName)
	}
}
