package proxy

import "net/url"

type MLModel struct {
	Name string
	Url  *url.URL
}
