package util

import (
	"errors"
	"strings"
)

func ExtractUrlParam(url string, paramPos int) (string, error) {
	parts := strings.Split(url, "/")
	if len(parts) < paramPos+1 || parts[paramPos] == "" {
		return "", errors.New("no param")
	}
	return parts[paramPos], nil
}
