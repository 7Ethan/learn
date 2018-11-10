package main

import (
	"os"
	"os/exec"
	"syscall"
)

func checkErr(err error) error {
	if err != nil {
		return err
		os.Exit(1)
	}
	return nil
}

func main() {
	bin, err := exec.LookPath("ls")
	checkErr(err)

	args := []string{"ls", "-a", "-l"}

	env := os.Environ()

	execErr := syscall.Exec(bin, args, env)
	if execErr != nil {
		panic(err)
	}
}
