package retention

import (
	"context"
	"fmt"
	"saveup/internal/storage"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type Retention struct {
	Client storage.S3Client
}

func (r *Retention) Run(ctx context.Context) error {
	result, err := r.Client.List(ctx)
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}

	objectCount := len(result.Contents)
	if objectCount == MinObjectToExist {
		return nil
	}

	var objectIdentifiers []types.ObjectIdentifier

	now := time.Now()
	for i := 0; i < objectCount-MinObjectToExist; i++ {
		elapsedTime := now.Sub(*result.Contents[i].LastModified)
		if elapsedTime.Seconds() > MaxObjectLifeInSecond {
			objectKey := *result.Contents[i].Key
			objectIdentifiers = append(
				objectIdentifiers, types.ObjectIdentifier{
					Key: aws.String(objectKey),
				},
			)
		}
	}

	if len(objectIdentifiers) == 0 {
		return nil
	}

	err = r.Client.BulkDelete(ctx, objectIdentifiers)
	if err != nil {
		return fmt.Errorf("bulk delete: %w", err)
	}

	return nil
}
