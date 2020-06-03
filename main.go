
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

func Banner() {
        color.HiCyan(`
 ____ _____________.____   __________              ___.
|    |   \______   \    |  \______   \_______  ____\_ |__   ____
|    |   /|       _/    |   |     ___/\_  __ \/  _ \| __ \_/ __ \
|    |  / |    |   \    |___|    |     |  | \(  <_> ) \_\ \  ___/
|______/  |____|_  /_______ \____|     |__|   \____/|___  /\___  >
                 \/        \/                           \/     \/   V 2.0`)
        color.HiYellow("           URLProbe:- Urls Status Code Checker")
        color.HiRed("              https://github.com/1ndianl33t")

        color.HiCyan("------------------------------------------------------------------------->
}
func printStatus(req *http.Request, resp *http.Response, err error) {
        if err != nil {
                return
        }
        StatusCheck(req, resp)

}

func StatusCheck(req *http.Request, resp *http.Response) {
        if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
                color.HiGreen("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)

        }
        if resp.StatusCode >= 300 && resp.StatusCode <= 308 {
                color.HiBlue("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)

        }
        if resp.StatusCode >= 400 && resp.StatusCode <= 451 {
                color.HiRed("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)

        }

        if resp.StatusCode >= 500 && resp.StatusCode <= 511 {
                color.HiCyan("[%d] L %d : %s\n", resp.StatusCode, resp.ContentLength, req.URL)

        }
}

func main() {
        Banner()
        var concurrency = 50
        flag.IntVar(&concurrency, "c", 50, "")
        var times int = 05
        flag.IntVar(&times, "t", 05, "")
        flag.Parse()
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
