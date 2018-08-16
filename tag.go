package graphite

type Tag struct {
	Name  string
	Value string
}

func (tag Tag) String() string {
	return tag.Name + "=" + tag.Value
}
