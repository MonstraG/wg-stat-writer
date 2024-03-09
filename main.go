package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	argsWithoutProg := os.Args[1:]
	folder := "results/"
	if len(argsWithoutProg) > 0 {
		folder = argsWithoutProg[0]
	}
	if !strings.HasSuffix(folder, "/") {
		folder += "/"
	}

	out, err := exec.Command("wg", "show").Output()
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().UTC().Format(time.RFC3339)

	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%v%v", folder, date)

	err = os.WriteFile(path, out, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
