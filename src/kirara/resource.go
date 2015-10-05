package kirara

type Tag string

type Post struct {
	PostID string
	ImageURLs []string
	ThumbnailImageURL string
	Tags []Tag
}

type Resource interface {
	ID() string
	Name() string

	SearchByKeyword(keyword string, pageNumber int) ([]Post, error)
	SearchByTags(tags []Tag, pageNumber int) ([]Post, error)
	AutoComplete(input string) ([]Tag, error)
}