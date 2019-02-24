package main

import (
	"fmt"
	"os"
	"time"

	"github.com/neranjana/content4go/logger"
)

func main() {
	logger.Init(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
	logger.Trace.Println("Trace message")
	logger.Info.Println("Info message")
	logger.Warning.Println("Warning message")
	logger.Error.Println("Error message")

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	fmt.Println("Unix timestamp")
	var timestamp int64
	timestamp = t.Unix()
	fmt.Println(timestamp)
	var nowTime time.Time
	nowTime = time.Unix(timestamp, 0)
	fmt.Println("time now")
	fmt.Println(nowTime.Day())

}
