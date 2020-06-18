package awssqs

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// Message type contains exported fields
type Message struct {
	Body          []byte
	ReceiptHandle string
}

// New returns pointer to a new Message
func New() *Message {
	return new(Message)
}

// GetSingleMessage function returns single message from specified SQS queue
func (msg *Message) GetSingleMessage(sqsSession *sqs.SQS, queueURL *string) {
	result, err := sqsSession.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
	})
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	for _, m := range result.Messages {
		msg.Body = append(msg.Body, []byte(*m.Body)...)
	}

	if len(result.Messages) == 0 {
		fmt.Println("Received no messages")
		return
	}

	resultDelete, err := sqsSession.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: result.Messages[0].ReceiptHandle,
	})

	if err != nil {
		fmt.Println("Delete Error ", err)
		return
	}

	fmt.Println("Message Deleted ", resultDelete)

}
