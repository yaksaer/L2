package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func GetTime(server string) (*time.Time, error) {
	time, err := ntp.Time(server)
	if err != nil {
		return nil, err
	}
	return &time, nil
}

func main() {
	time, err := GetTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(time)
}
