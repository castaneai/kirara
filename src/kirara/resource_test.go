package kirara

import (
	"testing"
)

func TestTiqav(t *testing.T) {
	r := Tiqav {}
	posts, err := r.SearchByKeyword("ちくわ", 1)
	if err != nil {
		t.Errorf("search by keyword failed.")
	}
	for i, post := range posts {
		t.Logf("%d: %v", i, post)
	}
}
