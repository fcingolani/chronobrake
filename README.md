# ChronoBrake

[![Docker Automated buil](https://img.shields.io/docker/automated/fcingolani/chronobrake.svg)](https://hub.docker.com/r/fcingolani/chronobrake/)

ChronoBrake will hold your request for a specified amount of time.

## Usage

Start the server

```
$ go get github.com/fcingolani/chronobrake
$ $GOPATH/bin/chronobrake
{"level":"info","msg":"Listening on 3000","time":"2016-10-18T11:57:45-03:00"}
```
Then make a request, using the desired hold time in milliseconds as its path. For example this request will take 10 seconds to complete:

```
$ curl http://localhost:3000/10000
10000
```

## Options

ChronoBrake accepts the following arguments:

```
  -port string
        server port (default "3000")
```

## Docker

```
$ docker run -d -p 3000:3000 fcingolani/chronobrake
b3eb446e25cd071007bbe01c53f5ae32aa6b6769c0d9c51d7d19fbc2f4fdea7b
$ curl http://localhost:3000/5000
5000
```
