package utils

import "fmt"

func GetMediaInfoMessage(post, web, id string) string {
	if web == "" || id == "" {
		return fmt.Sprintf("Post: %s", post)
	}
	return fmt.Sprintf("Post: %s\n%s: @%s", post, web, id)
}
