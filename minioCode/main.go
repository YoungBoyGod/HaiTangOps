package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// 将错误处理独立出来便于维护
func handleError(err error) {
	log.Printf("Error: %v\n", err)
}

func ListBucketsAndObjects(minioClient *minio.Client, endpoint string, errChan chan error) {
	defer close(errChan) //确保关闭通道

	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		errChan <- err
		return
	}
	var wg sync.WaitGroup
	for _, bucket := range buckets {
		wg.Add(1)
		go func(bucket minio.BucketInfo) {
			defer wg.Done()
			if err := minioClient.MakeBucket(context.Background(), bucket.Name, minio.MakeBucketOptions{}); err != nil {
				errChan <- err

			}

		}(bucket)

	}
	wg.Wait()
	close(errChan) // 所有任务完成后关闭错误通道
}

func listObjects(minioClient *minio.Client, bucket minio.BucketInfo, endpoint string, errChan chan error) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	objectCh := minioClient.ListObjects(ctx, bucket.Name, minio.ListObjectsOptions{
		Prefix:    "",
		Recursive: true,
	})

	for object := range objectCh {
		if object.Err != nil {
			errChan <- object.Err
			continue
		}

		url := fmt.Sprintf("http://%s/%s/%s", endpoint, bucket.Name, object.Key)
		fmt.Printf("%s %s %s %s\n", bucket.Name, object.Key, object.LastModified, url)
	}
	return nil
}

func main() {

	// 使用环境变量获取敏感信息
	endpoint := getEnv("MINIO_ENDPOINT", "192.168.10.215:9000")
	accessKeyID := getEnv("MINIO_ACCESS_KEY", "agusX529pYFYvwgAAVwo")
	secretAccessKey := getEnv("MINIO_SECRET_KEY", "8QKyvVe1t80uQxBLhAJbrIqgKOuJ2panMRRvSbyz")
	useSSL := getEnv("MINIO_USE_SSL", "false") == "true"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		handleError(err)
		return
	}
	errChan := make(chan error, 10)
	ListBucketsAndObjects(minioClient, endpoint, errChan)
	go listObjects(minioClient, bucket, endpoint, errChan)

	for err := range errChan {
		handleError(err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
