package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

const defaultResultsFolder = "results"

func getResultsFolder() string {
	pathArg := flag.String("path", defaultResultsFolder, "Folder to save results to")
	flag.Parse()
	return path.Clean(*pathArg)
}

func main() {
	resultsFolder := getResultsFolder()

	out, err := exec.Command("wg", "show").Output()
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(resultsFolder, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().UTC().Format(time.RFC3339)
	resultsFilePath := path.Join(resultsFolder, date)

	err = os.WriteFile(resultsFilePath, out, 666)
	if err != nil {
		log.Fatal(err)
	}
}
