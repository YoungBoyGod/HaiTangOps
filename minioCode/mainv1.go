package main

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

/*
description: 用于打印出minio public桶内文件的直接下载url.
author: YoungGodBoy
date:2024年6月25日
version: 1.0.0
commit:初始化代码和编写初步框架
version: 1.0.1
Todo: 需要结合写一个html 用于显示下载链接
*/

func mainv1() {
	endpoint := "192.168.10.215:9000"
	accessKeyID := "agusX529pYFYvwgAAVwo"
	secretAccessKey := "8QKyvVe1t80uQxBLhAJbrIqgKOuJ2panMRRvSbyz"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Println("err")
	}

	// log.Printf("hello %#v\n", minioClient) // minioClient is now set up

	buckets, err := minioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		objectCh := minioClient.ListObjects(ctx, bucket.Name, minio.ListObjectsOptions{
			Prefix:    "", // 默认为空,只显示带有以xx前缀的文件，相当于屏蔽其余内容
			Recursive: true,
		})
		for object := range objectCh {
			if object.Err != nil {
				fmt.Println(object.Err)
				return
			}
			// 获取每个object的访问url
			url := fmt.Sprintf("http://%s/%s/%s", endpoint, bucket.Name, object.Key)
			fmt.Println(bucket.Name, object.Key, object.LastModified, url)
		}
	}
}
