package musel

import (
	"slices"
	"strings"
)

type Control struct {
	Name         string
	SearchURL    string
	UpdateURL    string
	Placeholder  string
	SelectedKeys []string
}

func (c *Control) StringValue() string {
	return strings.Join(c.SelectedKeys, ",")
}

func (c *Control) RemoveKey(key string) {
	for i, k := range c.SelectedKeys {
		if k == key {
			c.SelectedKeys = slices.Delete(c.SelectedKeys, i, i+1)
			break
		}
	}
}

func ControlSelectedKeysFromString(s string) []string {
	raw := strings.Split(s, ",")
	res := make([]string, 0, len(raw))
	for _, s := range raw {
		ts := strings.TrimSpace(s)
		if len(ts) > 0 {
			res = append(res, ts)
		}
	}
	return res
}

type Option struct {
	Key   string
	Title string
}

type Options struct {
	SearchQuery string
	ControlName string
	SelectURL   string
	List        []Option
	EmptyText   string
}
