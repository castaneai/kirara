package resource

type Tag string

type Image struct {
	ImageID string
	URL string
	Width int
	Height int
}

type Post struct {
	PostID string
	Images []Image
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

type LoginRequest map[string]string

type Authenticator interface {
	Login(request LoginRequest) error
	Logout() error
}