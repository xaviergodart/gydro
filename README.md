# Gydro [![Build Status](https://travis-ci.org/xaviergodart/gydro.svg?branch=master)](https://travis-ci.org/xaviergodart/gydro) [![Go Report Card](https://goreportcard.com/badge/github.com/xaviergodart/gydro)](https://goreportcard.com/report/github.com/xaviergodart/gydro) [![GoDoc](https://godoc.org/github.com/xaviergodart/gydro?status.svg)](https://godoc.org/github.com/xaviergodart/gydro)

A lightweight API Gateway with zero dependencies for JSON services. You can see it as a very simple alternative to [Kong](https://getkong.org/), [Tyk](https://tyk.io/) or [API Umbrella](https://apiumbrella.io/).

## Features

- No external databases needed
- Administration via RESTFul api
- Consumer auth by api key
- Round-robin load balancing per route between multiple backends
- Persistent rate limiter

## TODO

- Add groups authorization
- Add circuit breaker middleware
- Add JWT auth
- Add some tests...

## Under the hood

- (https://github.com/HouzuoGuo/tiedot)[Tiedot] as datastore
- (https://github.com/tidwall/buntdb)[BuntDB] for rate limiter middleware backend
- (https://github.com/gorilla/mux)[Gorilla Mux] for api routing
- (https://github.com/vulcand/oxy/)[Oxy] for forwarding and load balancing requests
- (https://github.com/labstack/echo)[Echo] for admin api
