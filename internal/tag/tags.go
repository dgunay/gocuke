package tag

import "github.com/cucumber/messages-go/v16"

type Tags map[string]bool

func NewTags(tags ...string) Tags {
	res := map[string]bool{}
	for _, tag := range tags {
		res[tag] = true
	}
	return res
}

func NewTagsFromPickleTags(pickleTags []*messages.PickleTag) Tags {
	res := map[string]bool{}
	for _, tag := range pickleTags {
		res[tag.Name] = true
	}
	return res
}
