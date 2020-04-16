package aws

import (
	"fmt"

	awsarn "github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
)

type Instance interface {
	Create(svc iamiface.IAMAPI) error
	Update(svc iamiface.IAMAPI) error
	Delete(svc iamiface.IAMAPI) error
	ARN() awsarn.ARN
	IsCreated() bool
}

// ARNify turns a list of string inputs into a list of parsed ARNs
func ARNify(input ...string) ([]awsarn.ARN, error) {
	arns := []awsarn.ARN{}
	for i, str := range input {
		if !awsarn.IsARN(str) {
			return arns, fmt.Errorf("input '%s' at index '%d' is not a valid ARN", str, i)
		}
		arn, _ := awsarn.Parse(str)
		arns = append(arns, arn)
	}
	return arns, nil
}

type InstanceError struct {
	Code    ErrorCode
	Message string
}

func (is InstanceError) Error() string {
	return is.Message
}

func (is InstanceError) IsOfErrorCode(code ErrorCode) bool {
	return is.Code == code
}

type ErrorCode string

const ErrAWSInstanceNotYetCreated ErrorCode = "AWS Instance has not been created"

func NewInstanceError(code ErrorCode, msg string) InstanceError {
	return InstanceError{Code: code, Message: msg}
}

func NewInstanceNotYetCreatedError(msg string) InstanceError {
	return NewInstanceError(ErrAWSInstanceNotYetCreated, msg)
}