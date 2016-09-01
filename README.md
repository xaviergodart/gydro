# Gydro [![Build Status](https://travis-ci.org/xaviergodart/gydro.svg?branch=master)](https://travis-ci.org/xaviergodart/gydro) [![Go Report Card](https://goreportcard.com/badge/github.com/xaviergodart/gydro)](https://goreportcard.com/report/github.com/xaviergodart/gydro) [![GoDoc](https://godoc.org/github.com/xaviergodart/gydro?status.svg)](https://godoc.org/github.com/xaviergodart/gydro)

A lightweight API Gateway with zero dependencies for JSON services written in Go. You can see it as a very simple alternative to [Kong](https://getkong.org/), [Tyk](https://tyk.io/) or [API Umbrella](https://apiumbrella.io/).

**Gydro is still under development and should not be used in production.**

## Features

- No external databases needed
- Administration via RESTFul api
- Consumer auth by api key
- Round-robin load balancing per route between multiple backends
- Circuit breaker
- Persistent rate limiter
- Group based authorization

## TODO

- Add JWT auth
- Add support for an external database to allow multi host deployments
- Allow some configuration in order to override hardcoded default parameters...
- Add some tests...

## Under the hood

- [Tiedot](https://github.com/HouzuoGuo/tiedot) as datastore
- [BuntDB](https://github.com/tidwall/buntdb) for rate limiter middleware backend
- [Gorilla Mux](https://github.com/gorilla/mux) for api routing
- [Oxy](https://github.com/vulcand/oxy/) for forwarding and load balancing requests
- [Echo](https://github.com/labstack/echo) for admin api

## Installing

### From sources

Install Go >= 1.7, set your $GOPATH, and run :
```
go get github.com/xaviergodart/gydro
```

To start Gydro:
```
$GOPATH/bin/gydro
```

## Documentation

[Getting started](https://github.com/xaviergodart/gydro/wiki/Getting-started)
