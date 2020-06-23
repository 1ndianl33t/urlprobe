
# URLProbe v 2.1


*Take a list of URLs and returns their HTTP response codes*

***Warning: This is a work in progress.***

### What's New ( update)
```bash
Added -s flag to print specific status code. 
Changed colour of HTTP response 404 as red and other interesting responses like 401 and 403 to Cyan colour
```
## Install

With Go:

```bash
▶ go get -u github.com/1ndianl33t/urlprobe
```
## usage
```bash
▶ cat urls.txt | urlprobe 
▶ cat urls.txt | urlprobe -c 1000 -t 10
```
### Print Only Certain Status Codes

If you only want to print results that returned a certain status code, you can
use the `-s`  option:

```
▶ cat links | urlprobe -c 1000 -t 05 -s 200
```

## Example
```bash

▶ assetfinder --subs-only test.com | waybackurls | gf redirect | urlprobe -c 1000 -t 05
▶ Subfinder -d test.com | gau | gf ssrf | urlprobe -c 1000 -t 05

```
### Tips

grep 200ok urls 
```bash
cat hosts | urlprobe -c 1000 -t 01 | grep 200 | awk '{print $5}' | tee -a lol-200.txt 
```
## Output
```bash
[Status code] Content-Length : URL
[200] L 24 https://example.com/robots.txt
```
## Demo
https://asciinema.org/a/335804

### Donations
You can encourage me to contribute more to the open source with donations.

- Paypal - [https://www.paypal.me/1ndianl33t](https://www.paypal.me/1ndianl33t)

- GooglePay,Paytm -

`8085778875`

### Contributers

[![Twitter](https://img.shields.io/badge/twitter-@shivangx01b-blue.svg)](https://twitter.com/shivangx01b)


### Contact
[![Twitter](https://img.shields.io/badge/twitter-@1ndianl33t-blue.svg)](https://twitter.com/1ndianl33t)

