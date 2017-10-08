package main

import (
	"time"

	"github.com/gaeshi-gaeshi/thermo-smart/WebService"
)

func main() {
	WebService.Start()
	time.Sleep(time.Minute)
}
