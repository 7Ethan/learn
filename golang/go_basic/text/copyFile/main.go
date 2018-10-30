package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func fileExiting(name string) bool {
	_, err := os.Stat(name)
	return err == nil || os.IsExist(err)
}

func copyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()

	return io.Copy(srcFile, dstFile)
}

func copyFileAction(src, dst string, showProgress bool, force bool) {
	if !force {
		if fileExiting(dst) {
			fmt.Printf("%s file is exiting,overwrite? y / n\n", dst)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()

			if strings.TrimSpace(string(data)) != "y" {
				return
			}
		}
	}

	copyFile(src, dst)
}

func main() {
	var showProgress, force bool

	flag.BoolVar(&showProgress, "v", false, "show what is being done")
	flag.BoolVar(&force, "f", false, "force copy when is file exiting")

	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	copyFileAction(flag.Arg(0), flag.Arg(1), showProgress, force)
}
