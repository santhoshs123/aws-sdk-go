//go:build go1.7
// +build go1.7

package v4

import (
	"net/http"

	"github.com/santhoshs123/aws-sdk-go/aws"
)

func requestContext(r *http.Request) aws.Context {
	return r.Context()
}
