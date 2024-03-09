package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	out, err := exec.Command("wg show").Output()
	if err != nil {
		log.Fatal(err)
	}

	date := time.Now().Format(time.RFC3339)
	path := fmt.Sprintf("results/%v", date)
	err = os.WriteFile(path, out, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
