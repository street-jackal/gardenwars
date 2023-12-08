package google

import "google.golang.org/api/customsearch/v1"

// ImageSearch searches for images using the Google Custom Search API
type ImageSearch struct {
	CustomSearchService *customsearch.Service
}

// Search searches for images using the Google Custom Search API
func (i *ImageSearch) Search(query string) ([]string, error) {
	results, err := i.CustomSearchService.Cse.List().ImgSize("MEDIUM").Q(query).Do()
	if err != nil {
		return nil, err
	}

	images := make([]string, len(results.Items))
	for i, item := range results.Items {
		images[i] = item.Link
	}

	return images, nil
}
