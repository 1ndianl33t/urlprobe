
package main

import (
        "bufio"
        "flag"
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
func printStatus(req *http.Request, resp *http.Response,  err error) {
        if err != nil {
                return
        }
        StatusCheck(req, resp)

}

func ParseArguments() {
    flag.IntVar(&concurrency, "c", 50, "Number of workers to use..default 50")
    flag.IntVar(&status, "s", 1, "If enabled..then check for specific status")
    flag.IntVar(&times, "t", 05, "Set rate limit")
    flag.Parse()
}


func StatusCheck(req *http.Request, resp *http.Response) {
        if status != 1 {
            if status == resp.StatusCode {
                if status == 404 {
                    color.HiRed("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)
                } 
                if status != 404 {
                    color.HiCyan("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)
                } 
                
            }
        } else {
            if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
                    color.HiGreen("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)

            }
            if resp.StatusCode >= 300 && resp.StatusCode <= 308 {
                    color.HiBlue("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)

            }
            if resp.StatusCode >= 400 && resp.StatusCode <= 451 {
                    if resp.StatusCode == 404 {
                        color.HiRed("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)
                    } else {
                        color.HiCyan("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)
                    }
                    
            }

            if resp.StatusCode >= 500 && resp.StatusCode <= 511 {
                    color.HiCyan("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)

            }
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
