package sqs

import "github.com/santhoshs123/aws-sdk-go/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
