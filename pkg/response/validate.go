package response

import (
	"strings"
)

type Validator interface {
	ResponseValidator(err string) []string
}

func ResponseValidator(err string) []string {
	return strings.Split(err,"\n")
}