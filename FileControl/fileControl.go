package FileControl

import (
	"fmt"
)

type filePath struct {
	FileName string
}

func GetFilePath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("file path cannot be empty")
	}
	return fmt.Sprintf("File path is: %s", path), nil

}
