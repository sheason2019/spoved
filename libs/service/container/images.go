package container_service

var optional_images = []string{"node:16-alpine"}

func ValidateImage(image string) (support bool) {
	for _, v := range optional_images {
		if image == v {
			return true
		}
	}
	return false
}
