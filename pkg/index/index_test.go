package index

import (
	"strings"
	"testing"
)

func TestGetSitemap(t *testing.T) {
	var url, eRobots, eTitle, eContent string

	url = "https://grumpy-learning.com/blog/2018/11/08/why-grumpyconf-has-no-sponsors/"
	eRobots = "index, follow"
	eTitle = "Why GrumpyConf Has No Sponsors — Grumpy Learning — Where even the CSS is grumpy"
	eContent = "The event is not cheap"

	p, err := GetPage(url)

	if err != nil {
		t.Error(err)
	}

	if p.Url != url {
		t.Errorf("\nExpected url to be set to \"%s\"\nActual: \"%s\"", url, p.Url)
	}

	if p.Robots != eRobots {
		t.Errorf("\nExpected robots to be set to \"%s\"\nActual: \"%s\"", eRobots, p.Robots)
	}

	if p.Title != eTitle {
		t.Errorf("\nExpected title to be set to \"%s\"\nActual: \"%s\"", eTitle, p.Title)
	}

	if strings.Contains(p.Content, "The event is not cheap") == false {
		t.Errorf("\nExpected content  to contain \"%s\"", eContent)
	}

}

func TestPage_Type(t *testing.T) {
	p := Page{}

	if p.Type() != "Page" {
		t.Errorf("\nExpected Type to return Page\nActual: \"%s\"", p.Type())
	}
}
