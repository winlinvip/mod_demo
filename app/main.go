package main

import (
	"fmt"
	"os"
	"private.me/mdi/logger"
)

func main() {
	tank := "console"
	if len(os.Args) > 1 {
		tank = os.Args[1]
	}

	l,err := logger.New(tank)
	if err != nil {
		fmt.Println(fmt.Sprintf("err is %+v", err))
		return
	}
	l.Printf("Hello, World")
}
