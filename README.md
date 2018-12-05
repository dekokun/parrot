# parrot

[![Build Status](https://travis-ci.org/dekokun/parrot.svg?branch=master)](https://travis-ci.org/dekokun/parrot)[![Go Report Card](https://goreportcard.com/badge/github.com/dekokun/parrot)](https://goreportcard.com/report/github.com/dekokun/parrot)[![codecov](https://codecov.io/gh/dekokun/parrot/branch/master/graph/badge.svg)](https://codecov.io/gh/dekokun/parrot)[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

Parrot is a dummy responce server.

## Usage

```
Usage of parrot:
  -port int
          listen port (default 30000)
```

## Example

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

## use with docker compose example

put below 2 files, docker-compose.yml and nginx.conf, to same dir and `docker-compose up`

### docker-compose.yml

```yml
version: '3'
services:
  parrot:
    image: dekokun/parrot
  proxy:
    image: 'nginx:1.15.7'
    volumes:
      - './nginx.conf:/etc/nginx/nginx.conf'
    depends_on:
      - parrot
    ports:
      - '80:80'
```

### nginx.conf
```Nginx
user  nginx;
worker_processes  1;

error_log  /dev/stderr info;
pid        /var/run/nginx.pid;


events {
    worker_connections  1024;
}

http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format ltsv "time:$time_local"
                "\tstatus:$status"
                "\treqtime:$request_time"
                "\tupstream:$upstream_addr"
                "\tupstream_status:$upstream_status"
                "\tvhost:$host";

    access_log  /dev/stdout  ltsv;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  65;

    #gzip  on;
    upstream myapp {
        server parrot:8080;
    }

    server {
        listen       80;
        server_name  localhost;
        location / {
            proxy_pass   http://myapp;
        }
    }

}
```

## License

MIT

## Author

Shintaro Kurachi (a.k.a. dekokun)
