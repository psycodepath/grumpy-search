package sitemap

import (
	"testing"
)

func TestGetSitemap(t *testing.T) {

	entries, err := GetSitemap("https://grumpy-learning.com/sitemap.xml")

	if err != nil {
		t.Error(err.Error())
	}
	// stupid test but i don't got the source so don't know how many pages there should be
	if len(entries) < 1029 {
		t.Errorf("Expected Entries to be at least 1029 %d", len(entries))
	}
}
