package images

var OptionalCompileImages = []string{"node:16-alpine", "golang:1.20.0-alpine3.17"}
var OptionalDeployImages = []string{"node:16-alpine", "golang:1.20.0-alpine3.17", "root/spoved-nginx"}

func ValidateImage(image string, method string) (support bool) {
	var images []string
	if method == "compile" {
		images = OptionalCompileImages
	} else if method == "deploy" {
		images = OptionalDeployImages
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
