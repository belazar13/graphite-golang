// Example of using the graphiteNop feature in action:
//
//	package main
//
//	import (
//		"log"
//		"github.com/belazar13/graphite-golang"
//		"time"
//	)
//
//	func main() {
//
//		Graphite, err := graphite.NewGraphite("graphite.example.com", 2004)
//		if err != nil {
//			log.Fatalf("%v", err)
//		}
//
//		metric := graphite.NewTaggedMetric("some.value", 100, time.Now().Unix()-100)
//		metric.AddTag(&graphite.Tag{Name: "tag1", Value: "value1"})
//		metric.AddTag(&graphite.Tag{Name: "tag2", Value: "value2"})
//
//		Graphite.SendTaggedMetric(metric)
//	}

package graphite
