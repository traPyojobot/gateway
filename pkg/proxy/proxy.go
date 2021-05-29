package proxy

import "net/url"

type MLModel struct {
	Name  string
	Group string
	Url   url.URL
}
