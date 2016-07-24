package resource

import (
	"fmt"
	"github.com/franela/goreq"
)

type Tiqav struct {}

func (t *Tiqav) ID() string {
	return "tiqav"
}

func (t *Tiqav) Name() string {
	return "Tiqav"
}

type searchResponseData struct {
	Id string
	Ext string
	SourceUrl string
	Width int
	Height int
}

func getPostTags(postID string) ([]Tag, error) {
	res, err := goreq.Request{Uri: fmt.Sprintf("http://api.tiqav.com/images/%s/tags.json", postID)}.Do()
	if err != nil {
		return nil, err
	}
	var tags []Tag
	res.Body.FromJsonTo(&tags)
	return tags, nil
}

func (t *Tiqav) SearchByKeyword(keyword string, pageNumber int) ([]Post, error) {
	res, err := goreq.Request{Uri: fmt.Sprintf("http://api.tiqav.com/search.json?q=%s", keyword)}.Do()
	if err != nil {
		return nil, err
	}

	var data []searchResponseData
	res.Body.FromJsonTo(&data)
	var posts []Post
	for _, d := range data {
		tags, err := getPostTags(d.Id)
		if err != nil {
			return nil, err
		}
		posts = append(posts, Post {
			PostID: d.Id,
			Images: []Image { {ImageID: d.Id, URL: fmt.Sprintf("http://img.tiqav.com/%s.%s", d.Id, d.Ext), Width: d.Width, Height: d.Height} },
			ThumbnailImageURL: fmt.Sprintf("http://img.tiqav.com/%s.th.jpg", d.Id),
			Tags: tags,
		})
	}
	return posts, nil
}

func (t *Tiqav) SearchByTags(tags []Tag, pageNumber int) ([]Post, error) {
	return []Post {}, nil
}

func (t *Tiqav) AutoComplete(input string) ([]Tag, error) {
	return []Tag{}, nil
}
