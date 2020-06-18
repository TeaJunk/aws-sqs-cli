package awssqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// NewSession creates new AWS sessions with only limited parameters specified. Configuration migth be overwritten by environment variables.
// Please refer to AWS GO SDK
func NewSession(awsprofile, awsregion *string) *sqs.SQS {
	sOpts := session.Options{SharedConfigState: session.SharedConfigEnable}
	if *awsprofile != "" {
		sOpts.Profile = *awsprofile
	}
	sOpts.Config = aws.Config{
		Region: awsregion,
	}
	sess := session.Must(session.NewSessionWithOptions(sOpts))
	return sqs.New(sess)
}
