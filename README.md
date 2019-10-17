# Unofficial Go TBA API

This is a wrapper of [The Blue Alliance Read-Only v3 API](https://www.thebluealliance.com/apidocs) for [Golang](https://golang.org).
It supports the `If-Modified-Since` header to prefer cached data over new data and a custom `User-Agent` header.
Authentication with the client is done by passing your API key to each function on calling.
