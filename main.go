package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

const defaultResultsFolder = "results"

func getResultsFolder() string {
	folder := defaultResultsFolder

	if len(os.Args) > 1 {
		folder = os.Args[1]
	}

	return path.Clean(folder)
}

func main() {
	out, err := exec.Command("wg", "show").Output()
	if err != nil {
		log.Fatal(err)
	}

	resultsFolder := getResultsFolder()

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
