# Git remote

Executes  git operations in a remote server. It requires git installed in the remote server.

[![Build Status](https://travis-ci.org/ottogiron/gitremote.svg?branch=master)](https://travis-ci.org/ottogiron/gitremote)
[![GoDoc](https://godoc.org/github.com/ottogiron/gitremote?status.svg)](https://godoc.org/github.com/ottogiron/gitremote)
[![Go Report Card](https://goreportcard.com/badge/github.com/ottogiron/gitremote)](https://goreportcard.com/report/github.com/ottogiron/gitremote)

## Installation

Find binaries for linux and MacOS in the [releases page](https://github.com/ottogiron/gitremote/releases)

## Getting Started


### Run the server in the remote host

***Configuration (default location $HOME/.gitremote.yaml):***

```yaml
---
port: 2183
allowed-directories:
    - /path/in/server/to/allowed/git/repo

allowed-commands:
    - git status
    - git log
    - git pull
```

```bash
gitr serve
```

### Run git comands in a client machine


```bash
gitr exec --command='git status' --working-dir="/path/to/git/repo/in/remote/server" --host-address="localhost:2183"
gitr exec --command='git add .' --working-dir="/path/to/git/repo/in/remote/server" --host-address="localhost:2183"
gitr exec --command='git commit -m "Add latest changes remotely"' --working-dir="/path/to/git/repo/in/remote/server" --host-address="localhost:2183"
gitr exec --command='git push' --working-dir="/path/to/git/repo/in/remote/server" --host-address="localhost:2183"
```

### Commands Usage

## Serve

```bash
git serve [flags]

Starts the git remote server

Example:.
gitr serve --port=2183

Usage:
  gitremote serve [flags]

Flags:
  -a, --allowed-commands stringSlice      The list of allowed git commands
  -d, --allowed-directories stringSlice   The list of allowed git executable directories
  -c, --cert-file string                  The TLS cert file (default "server.pem")
  -k, --key-file string                   The TLS key file (default "server.key")
  -p, --port int                          Port for the rpc service (default 2183)
  -t, --tls                               Connection uses TLS if true, else plain TCP

Global Flags:
      --config string   config file (default is $HOME/.gitremote.yaml)
```

## Exec

```bash
gitr exec [flags]

Executes a git command in a  remote server:

Example:.
gitr exec --command="git status" --working-dir=/home/otto/myproject --host-address=myhost:2183

The above executes translates to "git status" command running in the /home/otto/myproject directory

Usage:
  gitremote exec [flags]

Flags:
  -c, --command string        git command to be executed example: git status
  -a, --host-address string   git remote server address (default "localhost:2183")
  -w, --working-dir string    Working directory of the git command to run (default ".")

Global Flags:
      --config string   config file (default is $HOME/.gitremote.yaml)
```

