// +build !integration

package pipeline

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mpucholblasco/s3logsbeat/aws"

	"github.com/elastic/beats/libbeat/common"
)

func TestGetKeyFields(t *testing.T) {
	bucket := "mybucket"
	key := "myenvironment-myapp/myawsregion/myfile.gz"
	sqsMessage := &SQSMessage{
		sqs: &SQS{
			keyRegexFields: regexp.MustCompile(`^(?P<environment>[^\-]+)-(?P<application>[^/]+)/`),
		},
	}
	s3object := NewS3Object(aws.NewS3Object(bucket, key), sqsMessage)
	keyFields, err := s3object.GetKeyFields()
	expectedKeyFields := common.MapStr{
		"environment": "myenvironment",
		"application": "myapp",
	}

	assert.NoError(t, err)
	assert.Equal(t, expectedKeyFields, *keyFields)
}
