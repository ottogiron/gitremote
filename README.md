# Git remote

Executes  git operations in a remote server. It requires git installed in the remote server.

[![Build Status](https://travis-ci.org/ottogiron/gitremote.svg?branch=master)](https://travis-ci.org/ottogiron/gitremote)
[![GoDoc](https://godoc.org/github.com/ottogiron/gitremote?status.svg)](https://godoc.org/github.com/ottogiron/gitremote)
[![Go Report Card](https://goreportcard.com/badge/github.com/ottogiron/gitremote)](https://goreportcard.com/report/github.com/ottogiron/gitremote)

## Installation

TODO

## Getting Started


### Run the server in the remote server

```bash
gitr serve --port=8082
```

### Execute a git command

#### Clone

```bash
gitr clone <git_url> [path] --remote_host_url=<my_host_url>
```


