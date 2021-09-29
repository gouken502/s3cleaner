package s3_all_delete

import (
	
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ----------------------------------------------------
// 全体変数
// ----------------------------------------------------
var (
	
)


/*
   @brief  	エントリポイント

   @param 	-

   @return  -

   @note

   @exception
			-
   @history
*/

// Handle xxx

func s3_all_delete(s3bucket string) {
	//-----------------------------------------------
	// S3からファイル取得
	//-----------------------------------------------
	//awsセッション
	sess := session.Must(session.NewSession())

	// S3 clientを作成
	s3Client := s3.New(sess)

	// ファイルの存在確認
	loo, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(s3bucket),
		//Prefix: aws.String("/"),
	})
	if err != nil {
		return
	}

	// データが存在しなかった場合
	if *loo.KeyCount == 0 {
		return
	}

	for _, item := range loo.Contents {

		key := *item.Key

		if !strings.Contains(key, ".bin") {
			continue
		}

		// Object取得
		_, err := s3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(s3bucket),
			Key:    aws.String(key),
		})

		if err != nil {
			return
		}

		err = deleteS3object(key,s3bucket)
		if err != nil {
			return
		}
	}

	return
}

func deleteS3object(key string,s3bucket string) error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// S3 clientを作成
	s3Client := s3.New(sess)

	_, err := s3Client.DeleteObject(
		&s3.DeleteObjectInput{
			Bucket: aws.String(s3bucket),
			Key:    aws.String(key),
		})
	if err != nil {
		return err
	}

	return nil
}

