# Unofficial Go TBA API

![GitHub](https://img.shields.io/github/license/akrantz01/go-tba)
[![GoDoc](https://godoc.org/github.com/akrantz01/go-tba?status.svg)](https://godoc.org/github.com/akrantz01/go-tba)

This is a wrapper of [The Blue Alliance Read-Only v3 API](https://www.thebluealliance.com/apidocs) for [Golang](https://golang.org).
It supports the `If-Modified-Since` header to prefer cached data over new data and a custom `User-Agent` header.
Authentication with the client is done by passing your API key to each function on calling.
Unlike the [auto-generated one](https://github.com/TBA-API/tba-api-client-go), this only uses the standard library and thus requires no dependencies.
