package graphite

import (
	"fmt"
	"strings"
	"time"
)

// TaggedMetric is a struct that defines the relevant properties of a graphite metric for Pickle format
type TaggedMetric struct {
	Name      string
	Value     int
	Timestamp int64
	Tags      []*Tag
}

func NewTaggedMetric(name string, value int, timestamp int64) *TaggedMetric {
	return &TaggedMetric{
		Name:      name,
		Value:     value,
		Timestamp: timestamp,
		Tags:      []*Tag{},
	}
}

func (metric *TaggedMetric) AddTag(tag *Tag) {
	metric.Tags = append(metric.Tags, tag)
}

func (metric *TaggedMetric) String() string {
	return fmt.Sprintf(
		"%s %d %s",
		metric.Name,
		metric.Value,
		time.Unix(metric.Timestamp, 0).Format("2006-01-02 15:04:05"),
	)
}

func (metric *TaggedMetric) GetTags() string {
	if len(metric.Tags) == 0 {
		return ""
	}

	var tags []string
	for _, tag := range metric.Tags {
		tags = append(tags, tag.String())
	}

	return strings.Join(tags, ";")
}
