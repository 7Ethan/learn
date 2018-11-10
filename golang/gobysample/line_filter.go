package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)

	for scaner.Scan() {
		ucl := strings.ToUpper(scaner.Text())
		fmt.Fprintln(os.Stdout, ucl)
	}

	if err := scaner.Err(); err != nil {
		panic(err)
		os.Exit(1)
	}

	/*
		$ cat limit_rate.go| go run line_filter.go
		PACKAGE MAIN

		IMPORT (
				"FMT"
				"TIME"
		)

		FUNC MAIN() {
				REQ := MAKE(CHAN INT, 5)
				FOR I := 0; I < 5; I++ {
						REQ <- I
				}
				CLOSE(REQ)

				LIMITER := TIME.NEWTICKER(200 * TIME.MILLISECOND)

				FOR R := RANGE REQ {
						<-LIMITER.C
						FMT.PRINTLN("REQUEST : ", R, TIME.NOW())
				}
		}
	*/
}
