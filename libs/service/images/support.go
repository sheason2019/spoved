package images

var optional_compile_images = []string{"node:16-alpine", "golang:1.20.0-alpine3.17"}
var optional_deploy_images = []string{"node:16-alpine", "golang:1.20.0-alpine3.17", "root/spoved-nginx"}

func ValidateImage(image string, method string) (support bool) {
	var images []string
	if method == "compile" {
		images = optional_compile_images
	} else if method == "deploy" {
		images = optional_deploy_images
	} else {
		return false
	}

	for _, v := range images {
		if image == v {
			return true
		}
	}
	return false
}
