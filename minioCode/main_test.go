package main

import (
	"context"
	"testing"

	"github.com/minio/minio-go/v7"
)

func BenchmarkListObjects(b *testing.B) {
	// 初始化必要的变量，比如minioClient、ctx、bucket.Name等
	// 注意：真实环境下，这些初始化操作应放在setup函数中，并使用b.ResetTimer()避免计入测试时间
	minioClient, ctx, bucketName := initMinioClientAndContext()

	// 避免每次测试都执行初始化
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
			Prefix:    "",
			Recursive: true,
		})
		if err != nil {
			b.Fatal(err) // 在基准测试中遇到错误应终止测试
		}
	}
}

// initMinioClientAndContext 是一个示例函数，你需要根据实际情况实现它来初始化minioClient和context
func initMinioClientAndContext() (*minio.Client, context.Context, string) {
	// 实现初始化逻辑...
	return nil, nil, ""
}
