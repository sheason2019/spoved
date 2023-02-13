package images

var optional_images = []string{"node:16-alpine", "golang:1.20.0-alpine3.17"}

func ValidateImage(image string) (support bool) {
	for _, v := range optional_images {
		if image == v {
			return true
		}
	}
	return false
}
