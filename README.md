# parrot

[![Build Status](https://travis-ci.org/dekokun/parrot.svg?branch=master)](https://travis-ci.org/dekokun/parrot)[![Go Report Card](https://goreportcard.com/badge/github.com/dekokun/parrot)](https://goreportcard.com/report/github.com/dekokun/parrot)[![codecov](https://codecov.io/gh/dekokun/parrot/branch/master/graph/badge.svg)](https://codecov.io/gh/dekokun/parrot)[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Parrot is a dummy responce server.

## Usage

```
Usage of parrot:
  -port int
          listen port (default 8080)
```

## Example

### docker

```bash
$ docker pull dekokun/parrot
$ docker run -p 80:8080 dekokun/parrot
$ curl --header 'Foo:Bar' localhost
```

### docker compose

proxy server: nginx
app server: parrot

```bash
$ cd example
$ docker-compose up
$ curl --header 'Foo:Bar' localhost
```

### listen port 8080

```bash
$ parrot -port 8080
```

### Change responce header

```bash
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

```bash
$ curl --header 'X-Forwarded-For: x.x.x.x' localhost:8080
{"Accept":["*/*"],"Host":["localhost:8080"],"User-Agent":["curl/7.43.0"],"X-Forwarded-For":["x.x.x.x"]}
```

## Installation

```
$ go get github.com/dekokun/parrot
```

## License

MIT

## Author

Shintaro Kurachi (a.k.a. dekokun)
