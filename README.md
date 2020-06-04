
# URLProbe v 2.0


*Take a list of URLs and returns their HTTP response codes*

***Warning: This is a work in progress.***

## Install

With Go:

```bash
▶ go get -u github.com/1ndianl33t/urlprobe
```
## usage
```bash
▶ assetfinder --subs-only test.com | waybackurls | gf redirect | urlprobe -c 1000 -t 05
▶ Subfinder -d test.com | gau | gf ssrf | urlprobe -c 1000 -t 05
▶ cat urls.txt | urlprobe 
▶ cat urls.txt | urlprobe -c 100 -t 10
```
## Output
```bash
[Status code] Content-Length : URL
[200] L 24 https://example.com/robots.txt
```
## Demo
https://asciinema.org/a/335804


### Contact
[![Twitter](https://img.shields.io/badge/twitter-@1ndianl33t-blue.svg)](https://twitter.com/1ndianl33t)

