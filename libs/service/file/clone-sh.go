package file_service

import "fmt"

func cloneSh(url, branch string) string {
	return fmt.Sprintf(`
	GIT_SSH_COMMAND="ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no" git clone -b %s %s /code --progress
	`, branch, url)
}
