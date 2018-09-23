package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-al")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
