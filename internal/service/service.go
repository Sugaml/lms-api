package service

import "github.com/sugaml/lms-api/internal/core/port"

type Service struct {
	repo port.Repository
}

// NewAnnocuncementService creates a new product service instance
func NewService(
	repo port.Repository,
) port.Service {
	return &Service{
		repo,
	}
}

type mapString map[string]string

// func BulkImageWithCache(urls []*domain.ImageResponse) []*domain.ImageResponse {
// 	for i, url := range urls {
// 		urls[i].Url = helper.ImageWithCache(url.Url)
// 	}
// 	return urls
// }

// func BulkCleanImage(urls []*domain.ImageResponse) []*domain.ImageResponse {
// 	for i, url := range urls {
// 		urls[i].Url = helper.CleanImage(url.Url)
// 	}
// 	return urls
// }
