package retention

import (
	"context"
	"fmt"
	"saveup/internal/config"
	"saveup/internal/storage"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func Run(ctx context.Context, cfg *config.Config) error {
	result, err := storage.List(ctx, cfg)
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

	err = storage.BulkDelete(ctx, objectIdentifiers, cfg)
	if err != nil {
		return fmt.Errorf("bulk delete: %w", err)
	}

	return nil
}
