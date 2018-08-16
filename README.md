graphite-golang
===============

This is a lightweight Graphite API client written in Go that implements Carbon
submission functionality. It allows to send Tags with metric in Pickle format.

## Installation

Use `go-get` to install graphite-golang
```
go get github.com/belazar13/graphite-golang
```

## External dependencies

This project has external dependencies for support Pickle

```
go get github.com/hydrogen18/stalecucumber
```

## Documentation

Like most every other Golang project, this projects documentation can be found
on godoc at [godoc.org/github.com/belazar13/graphite-golang](http://godoc.org/github.com/belazar13/graphite-golang).

## Examples

```go
package mylib

import (
    "github.com/belazar13/graphite-golang"
    "log"
)

func init() {

    // load your configuration file / mechanism
    config := newConfig()

    // try to connect a graphite server
    if config.GraphiteEnabled {
        Graphite, err = graphite.NewGraphite(config.Graphite.Host, config.Graphite.Port)
    } else {
        Graphite = graphite.NewGraphiteNop(config.Graphite.Host, config.Graphite.Port)
    }
    // if you couldn't connect to graphite, use a nop
    if err != nil {
        Graphite = graphite.NewGraphiteNop(config.Graphite.Host, config.Graphite.Port)
    }

    log.Printf("Loaded Graphite connection: %#v", Graphite)
    Graphite.SimpleSend("stats.graphite_loaded", "1")
}

func doWork() {
    // this will work just fine, regardless of if you're working with a graphite
    // nop or not
    Graphite.SimpleSend("stats.doing_work", "1")
}
```

## Example with Tag and Pickle protocol

```go
package main

import (
	"log"
	"github.com/belazar13/graphite-golang"
	"time"
)

func main() {

	Graphite, err := graphite.NewGraphite("graphite.example.com", 2004)
	if err != nil {
		log.Fatalf("%v", err)
	}

	metric := graphite.NewTaggedMetric("some.value", 100, time.Now().Unix()-100)
	metric.AddTag(&graphite.Tag{Name: "tag1", Value: "value1"})
	metric.AddTag(&graphite.Tag{Name: "tag2", Value: "value2"})

	Graphite.SendTaggedMetric(metric)
}

```
