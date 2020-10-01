package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/tomnomnom/gahttp"
	"net/http"
	"os"
	"time"
)

var concurrency int
var times int
var status int

func Banner() {
	color.HiCyan(`
____ _____________.____   __________              ___.
|    |   \______   \    |  \______   \_______  ____\_ |__   ____
|    |   /|       _/    |   |     ___/\_  __ \/  _ \| __ \_/ __ \
|    |  / |    |   \    |___|    |     |  | \(  <_> ) \_\ \  ___/
|______/  |____|_  /_______ \____|     |__|   \____/|___  /\___  >
\/        \/                           \/     \/   V 2.1`)
	color.HiYellow("           URLProbe:- Urls Status Code & ContentLength Checker")
	color.HiRed("              https://github.com/1ndianl33t")

	color.HiCyan("-------------------------------------------------------------------------")
}
func printStatus(req *http.Request, resp *http.Response, err error) {
	if err != nil {
		color.HiRed(err.Error())
		return
	}

	StatusCheck(req, resp)

}

func ParseArguments() {
	flag.IntVar(&concurrency, "c", 500, "Number of workers to use..default 500")
	flag.IntVar(&status, "s", 1, "If enabled..then check for specific status")
	flag.IntVar(&times, "t", 05, "Set rate limit")
	flag.Parse()
}

func colorPrintMessage(statusCode int, message string) {
	switch {
	case statusCode == 404:
		color.HiRed(message)
	case statusCode >= 200 && statusCode <= 299:
		color.HiGreen(message)
	case statusCode >= 300 && statusCode <= 308:
		color.HiBlue(message)
	case statusCode >= 400 && statusCode <= 511:
		color.HiCyan(message)
	}
}

func StatusCheck(req *http.Request, resp *http.Response) {
	message := fmt.Sprintf("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)

	if status != 1 {
		if status == resp.StatusCode {
			if status == 404 {
				colorPrintMessage(resp.StatusCode, message)
			} else {
				color.HiCyan(message)
			}

		}
	} else {
		colorPrintMessage(resp.StatusCode, message)
	}
}

func main() {
	Banner()
	ParseArguments()
	p := gahttp.NewPipeline()
	p.SetConcurrency(concurrency)
	p.SetRateLimit(time.Duration(times) * time.Second)
	urls := gahttp.Wrap(printStatus, gahttp.CloseBody)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		p.Get(sc.Text(), urls)
	}
	p.Done()
	p.Wait()
}
