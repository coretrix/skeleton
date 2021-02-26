package helper

import (
	"errors"
	"strings"
)

const (
	png  = "png"
	jpeg = "jpeg"
)

var supportedFileExtensions = []string{
	png, jpeg,
}

func GetFileExtensionFromBase64Header(base64Header string) (string, error) {
	for _, ext := range supportedFileExtensions {
		if strings.Contains(base64Header, ext) {
			return "." + ext, nil
		}
	}

	return "", errors.New("unsupported image type")
}
