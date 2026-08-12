package list

import (
	"context"
	"fmt"
	"saveup/internal/storage"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type List struct {
	Client storage.S3Client
}

func formatSize(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func formatAge(t time.Time) string {
	age := time.Since(t)

	switch {
	case age < time.Minute:
		return "<1m"
	case age < time.Hour:
		return fmt.Sprintf("%dm", int(age.Minutes()))
	case age < 24*time.Hour:
		return fmt.Sprintf("%dh", int(age.Hours()))
	default:
		return fmt.Sprintf("%dd", int(age.Hours()/24))
	}
}

func (l *List) Run(ctx context.Context) error {
	output, err := l.Client.List(ctx)
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}

	// Sort from newest
	sort.Slice(
		output.Contents, func(i, j int) bool {
			return output.Contents[i].LastModified.After(aws.ToTime(output.Contents[j].LastModified))
		},
	)

	fmt.Printf("%-40s %-10s %s\n", "BACKUP", "SIZE", "AGE")

	for _, object := range output.Contents {
		name := aws.ToString(object.Key)
		size := formatSize(aws.ToInt64(object.Size))
		age := formatAge(aws.ToTime(object.LastModified))

		fmt.Printf("%-40s %-10s %s\n", name, size, age)
	}

	return nil
}
