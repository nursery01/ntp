package main

import (
	"os/exec"
	"time"
)

func main() {
	for {
		a1 := exec.Command("ntpdate", "time.apple.com")
		a1.Run()
		time.Sleep(24 * time.Hour * 30)
	}

}
