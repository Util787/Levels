package main

import (
	"fmt"
	"os"

	"l2.8/common"
)

func main() {
	curTime, err := common.GetNtpTime()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	fmt.Println(curTime)
}
