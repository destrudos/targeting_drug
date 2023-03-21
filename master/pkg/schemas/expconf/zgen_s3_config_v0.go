// Code generated by gen.py. DO NOT EDIT.

package expconf

import (
	"github.com/santhosh-tekuri/jsonschema/v2"

	"github.com/determined-ai/determined/master/pkg/schemas"
)

func (s S3ConfigV0) Bucket() string {
	if s.RawBucket == nil {
		panic("You must call WithDefaults on S3ConfigV0 before .Bucket")
	}
	return *s.RawBucket
}

func (s *S3ConfigV0) SetBucket(val string) {
	s.RawBucket = &val
}

func (s S3ConfigV0) AccessKey() *string {
	return s.RawAccessKey
}

func (s *S3ConfigV0) SetAccessKey(val *string) {
	s.RawAccessKey = val
}

func (s S3ConfigV0) SecretKey() *string {
	return s.RawSecretKey
}

func (s *S3ConfigV0) SetSecretKey(val *string) {
	s.RawSecretKey = val
}

func (s S3ConfigV0) EndpointURL() *string {
	return s.RawEndpointURL
}

func (s *S3ConfigV0) SetEndpointURL(val *string) {
	s.RawEndpointURL = val
}

func (s S3ConfigV0) Prefix() *string {
	return s.RawPrefix
}

func (s *S3ConfigV0) SetPrefix(val *string) {
	s.RawPrefix = val
}

func (s S3ConfigV0) ParsedSchema() interface{} {
	return schemas.ParsedS3ConfigV0()
}

func (s S3ConfigV0) SanityValidator() *jsonschema.Schema {
	return schemas.GetSanityValidator("http://determined.ai/schemas/expconf/v0/s3.json")
}

func (s S3ConfigV0) CompletenessValidator() *jsonschema.Schema {
	return schemas.GetCompletenessValidator("http://determined.ai/schemas/expconf/v0/s3.json")
}
