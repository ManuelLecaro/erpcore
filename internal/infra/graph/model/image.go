package model

import (
	"fmt"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
)

func CreateNewImage(id, name, description, url string) *Image {
	return &Image{
		ID:          id,
		Name:        name,
		Description: description,
		URL:         url,
	}
}

func ToImageDTO(image domain.Image) *Image {
	return &Image{
		ID:          fmt.Sprintf("%d", image.ID),
		Name:        image.Name,
		Description: image.Description,
		URL:         image.URL,
	}
}

func FromImageDTO(image Image) *domain.Image {
	return &domain.Image{
		Name:        image.Name,
		Description: image.Description,
		URL:         image.URL,
	}
}

func ToDomainImageFromImput(image *NewImage) *domain.Image {
	return &domain.Image{
		Name:        image.Name,
		Description: image.Description,
		URL:         image.URL,
	}
}

func ToDomainImageFromInputArray(imagesArray []*NewImage) (images []domain.Image) {
	images = []domain.Image{}

	for _, im := range imagesArray {
		images = append(images, *ToDomainImageFromImput(im))
	}

	return
}
