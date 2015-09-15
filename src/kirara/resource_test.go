package kirara

import (
	"testing"
)

type testResource struct {}

func (t *testResource) ID() string {
	return "test-resource"
}

func (t *testResource) Name() string {
	return "Test Resource"
}

func (t *testResource) SearchByKeyword(keyword string, pageNumber int) ([]Post, error) {
	return []Post {}, nil
}

func (t *testResource) SearchByTags(tags []Tag, pageNumber int) ([]Post, error) {
	return []Post {}, nil
}

func (t *testResource) AutoComplete(input string) ([]Tag, error) {
	return []Tag {}, nil
}

func TestInterface(t *testing.T) {
	var r Resource
	r = &testResource{}
	if r.ID() != "test-resource" {
		t.Errorf("リソースIDが不正")
	}
	if r.Name() != "Test Resource" {
		t.Errorf("リソース名が不正")
	}
}
