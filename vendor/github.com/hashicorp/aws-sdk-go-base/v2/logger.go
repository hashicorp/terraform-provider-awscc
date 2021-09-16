package awsbase

import (
	"fmt"
	"log"

	"github.com/aws/smithy-go/logging"
)

type debugLogger struct{}

func (l debugLogger) Logf(classification logging.Classification, format string, v ...interface{}) {
	log.Printf("[%s] [aws-sdk-go-v2] %s", classification, fmt.Sprintf(format, v...))
}
