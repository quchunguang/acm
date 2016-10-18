# acm

[![Build Status](https://secure.travis-ci.org/quchunguang/acm.png?branch=master)](http://travis-ci.org/quchunguang/acm)
[![Go Report Card](https://goreportcard.com/badge/github.com/quchunguang/acm)](https://goreportcard.com/report/github.com/quchunguang/acm)
[![GoDoc](https://godoc.org/github.com/quchunguang/acm?status.svg)](https://godoc.org/github.com/quchunguang/acm)
[![Join the chat at https://gitter.im/acm-icpc/Lobby](https://badges.gitter.im/acm-icpc/Lobby.svg)](https://gitter.im/acm-icpc/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

ACM-ICPC problems' gallery.

## usage
You can run all programs by `go test`, or matching one of them with regular expression, like

```
go test -run TestSum4Zero
```

This will solve target problem with all different methods and given the performance data.

## arguments
You can optionally passing environment variables instead of command line arguments in some problem, like

```
N=100 go test -run TestSum4Zero
```
