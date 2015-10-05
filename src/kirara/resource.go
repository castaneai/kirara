package kirara

import (
	"net/url"
	"time"
)

type Tag string

type Post struct {
	PostID string
	ImageURLs []url.URL
	ThumbnailImageURL url.URL
	Tags []Tag
	CreatedAt time.Time
}

type Resource interface {
	ID() string
	Name() string

	SearchByKeyword(keyword string, pageNumber int) ([]Post, error)
	SearchByTags(tags []Tag, pageNumber int) ([]Post, error)
	AutoComplete(input string) ([]Tag, error)
}