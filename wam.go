package main

import (
	"fmt"
	"os"
	"github.com/taikedz/web-archives-manager/wam"
)

func main() {
	fore, aft := goargs.SplitTokensBefore("--", os.Args[1:])
	fmt.Printf("It works! %s / %s\n", fore, aft)
}