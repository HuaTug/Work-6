package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"HuaTug.com/cmd/publish/dal/db"
	"HuaTug.com/config/cache"
	"HuaTug.com/kitex_gen/publishs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	endpoint        = "localhost:9000"
	accessKeyID     = "minioadmin"
	secretAccessKey = "minioadmin"
	chunkSize       = 5 * 1024 * 1024 //5MB
)

type Uploadvideoservice struct {
	ctx context.Context
}

func NewUploadService(ctx context.Context) *Uploadvideoservice {
	return &Uploadvideoservice{ctx: ctx}
}

func (v *Uploadvideoservice) UploadFile(req *publishs.UpLoadVideoRequest) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	key := "video_id"
	Id := cache.GenerateID(key)
	minioClient, err := minio.NewCore(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return errors.WithMessage(err, "Minio server init failed")
	}
	bucketName := req.BucketName
	var filePath string
	var objectName string
	objectName = req.ObjectName

	switch req.ContentType {
	case "video/mp4":
		filePath = "/home/xuzh/Videos/" + req.Path
		objectName = req.ObjectName + ".mp4"
	case "png", "jpg", "jpeg":
		filePath = "/home/xuzh/Pictures/" + req.Path
		objectName = req.ObjectName + ".jpg"
	}

	fmt.Println(filePath)
	src, err := os.Open(filePath)
	if err != nil {
		return errors.WithMessage(err, "Failed to open file")
	}
	defer src.Close()

	wg.Add(1)
	go func() {
		defer wg.Done()
		exists, err3 := minioClient.BucketExists(context.Background(), bucketName)
		if err3 == nil && exists {
			logrus.Printf("Bucket %s already exists\n", bucketName)
		} else {
			err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		}
	}()
	wg.Wait()
	if err != nil {
		return errors.WithMessage(err, "MakeBucket failed")
	}
	uploadID, err := minioClient.NewMultipartUpload(context.Background(), bucketName, objectName, minio.PutObjectOptions{})
	if err != nil {
		return errors.WithMessage(err, "MultipartUpload file failed")
	}
	var parts []minio.ObjectPart
	buffer := make([]byte, chunkSize)
	for partNumber := 1; ; partNumber++ {
		n, err := src.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				// 其他错误，记录错误信息并返回
				return errors.WithMessage(err, "Error while reading file:")
			}
		}
		if n == 0 {
			break
		}

		//将每一个切片进行并发的上传
		part, err := minioClient.PutObjectPart(context.Background(), bucketName, objectName, uploadID, partNumber, bytes.NewReader(buffer[:n]), int64(n), minio.PutObjectPartOptions{})
		if err != nil {
			return errors.WithMessage(err, "Fail to uploadfile")
		}

		/*ToDo: 虽然src.Read() 是按顺序读取文件的，但在每次读取之后，都会启动一个 goroutine 来并发地执行上传操作，
		多个 goroutine 可能会同时向 parts 切片中添加元素。因此，为了避免多个 goroutine 同时修改 parts 切片而导致的竞态条件，你需要在对 parts 切片进行读写操作时加锁。
		*/

		mu.Lock()
		parts = append(parts, part)
		mu.Unlock()
	}
	var completeParts []minio.CompletePart
	mu.Lock()
	for _, part := range parts {
		completeParts = append(completeParts, minio.CompletePart{
			PartNumber: part.PartNumber,
			ETag:       part.ETag,
		})
	}
	mu.Unlock()
	_, err = minioClient.CompleteMultipartUpload(context.Background(), bucketName, objectName, uploadID, completeParts, minio.PutObjectOptions{})
	if err != nil {
		logrus.Info(err)
		return err
	}
	logrus.Info("UpLoad file Success!")

	/* 	 	go func() {
		_, err = minioClient.PutObject(context.Background(), bucketName, objectName, src, -1, minio.PutObjectOptions{})
		if err != nil {
			logrus.Info(err)
		}
		wg.Done()
	}()  */

	publish := &publishs.Video{
		VideoId:     Id,
		PlayUrl:     filePath,
		CoverUrl:    req.CoverUrl,
		PublishTime: time.Now().Format(time.DateTime),
		Title:       req.Title,
		AuthorId:    req.UserId,
	}
	if err := db.VideoCreate(v.ctx, &publishs.VideoCreateRequest{
		Video: publish,
	}); err != nil {
		return errors.WithMessage(err, "Fail to Write Sql")
	}
	logrus.Info("文件上传成功")
	return nil
}
