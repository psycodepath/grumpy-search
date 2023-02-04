package sitemap

import (
	sitemap "github.com/oxffaa/gopher-parse-sitemap"
)

func GetSitemap(url string) ([]sitemap.Entry, error) {
	entries := make([]sitemap.Entry, 500)

	err := sitemap.ParseFromSite(url, func(e sitemap.Entry) error {
		entries = append(entries, e)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return entries, nil
}
