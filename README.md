# parrot

[![Build Status](https://travis-ci.org/dekokun/parrot.svg?branch=master)](https://travis-ci.org/dekokun/parrot)[![Go Report Card](https://goreportcard.com/badge/github.com/dekokun/parrot)](https://goreportcard.com/report/github.com/dekokun/parrot)[![Coverage Status](https://coveralls.io/repos/github/dekokun/parrot/badge.svg?branch=travis_test)](https://coveralls.io/github/dekokun/parrot?branch=travis_test)[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Parrot is a dummy responce server.

## Example

### listen port 8080

```
$ parrot -p 8080
```

### Change responce header

```
$ curl -IXGET localhost:8080
HTTP/1.1 200 OK
Date: Thu, 13 Oct 2016 00:05:18 GMT
Content-Length: 73
Content-Type: text/plain; charset=utf-8

$ curl -IXGET localhost:8080\?Cache-Control=max-age=0
HTTP/1.1 200 OK
Cache-Control: max-age=0
Date: Thu, 13 Oct 2016 00:05:20 GMT
Content-Length: 73
Content-Type: text/plain; charset=utf-8
```

### Get request header

```
$ curl --header 'X-Forwarded-For: x.x.x.x' localhost:8080
{"Accept":["*/*"],"Host":["localhost:8080"],"User-Agent":["curl/7.43.0"],"X-Forwarded-For":["x.x.x.x"]}
```

