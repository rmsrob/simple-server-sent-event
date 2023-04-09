![WIP](https://img.shields.io/badge/status-wip-red)

[![Go Reference](https://pkg.go.dev/badge/github.com/rrobrms/simple-server-sent-event.svg)](https://pkg.go.dev/github.com/rrobrms/simple-server-sent-event)
[![Go Report Card](https://goreportcard.com/badge/github.com/rrobrms/simple-server-sent-event)](https://goreportcard.com/report/github.com/rrobrms/simple-server-sent-event)
[![Coverage Status](https://coveralls.io/repos/github/rrobrms/simple-server-sent-event/badge.svg?branch=master)](https://coveralls.io/github/rrobrms/simple-server-sent-event?branch=master)
<a href="https://gitpod.io/#https://github.com/rrobrms/simple-server-sent-event" target="_blank">
  <img
    src="https://img.shields.io/badge/Open%20with-Gitpod-908a85?logo=gitpod"
    alt="Contribute with Gitpod"
  />
</a>

# simple-server-sent-event

> Open a SSE connection and get a response message with its formatted time when hit an endpoint

## Install
### Prerequisite
> **Note**
>
>Required:
- [Golang V1.20](https://go.dev/doc/install)

```sh
git clone git@github.com:rrobrms/simple-server-sent-event.git
cd simple-server-sent-event
go mod tidy         # to download modules
```


## Usage
> Start the server
```sh
go run cmd/main.go
```

> Init the sse connection with an "keep-alive" header
- GET the endpoint `http://localhost:1323/event`
> Then when you hit the desired endpoint, we can compute the message
- GET on the endpoint `http://localhost:1323/post-event`

> Cleanned logs

```sh
⇨ http server started on [::]:1323
Client connected
data: 08-04-2023 23:59:25 CEST - Hello, world!

Client connected
data: 08-04-2023 23:59:36 CEST - Hello, world!

Client connected
data: 08-04-2023 23:59:57 CEST - Hello, world!

Client connected
data: 09-04-2023 00:00:12 CEST - Hello, world!
```