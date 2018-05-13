package services

import (
	"os"
	"io"
	"log"
	//  "crypto/sha256"

	"mime/multipart"
    "crypto/rand"
    "encoding/base64"
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"

	// env "../environment"
	"../config"
)


type gcpFileUploader struct {
	    // Some fields
	}
type localFileUploader struct {
	    // Some fields
	}

type FileUploader interface {
	FileSave(*multipart.FileHeader) string
}

var sharedLocalFileUploader *localFileUploader = &localFileUploader{}
var sharedGCPFileUploader *gcpFileUploader = &gcpFileUploader{}

func GetFileUploader() FileUploader {
	// switch setting.GetInstance().RunMode  {
	// case setting.Development: return sharedLocalFileUploader
	// case setting.Production: return sharedGCPFileUploader
	// default: return sharedLocalFileUploader
	// }
	c := config.GetInstance()
	switch c.AssetConfig.UseGCS {
		case true: return sharedGCPFileUploader
		case false: return sharedLocalFileUploader
		default: return sharedLocalFileUploader
	}
}


func(self *gcpFileUploader) FileSave(file *multipart.FileHeader) (string) {
	// TODO: gcp用のアップローダを定義


	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// アップロードされたファイルをオープン
	src, err := file.Open()
	if err != nil {
		log.Println(err)
	}
	defer src.Close()

	// ここでパスの確認
	fileName := getSafeFileName()
	// filePath := "uploads/image/" +  fileName

	// uploads
	bucketName := config.GetInstance().AssetConfig.GCSBucket
	err = gcsWrite(client, bucketName, fileName, src)
	if err != nil {
	}

	return fileName
}

func gcsWrite(client *storage.Client, bucket, object string, f  multipart.File) error {
	ctx := context.Background()

	obj := client.Bucket(bucket).Object(object)
	wc := obj.NewWriter(ctx)

	if _, err := io.Copy(wc, f); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}

	// 権限をpublicに
	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		log.Println(err.Error())
		return nil
	}
	// [END upload_file]
	return nil
}

func(self *localFileUploader) FileSave(file *multipart.FileHeader) (string) {
	// アップロードされたファイルをオープン
	src, err := file.Open()
	if err != nil {
		log.Println(err)
	}
	defer src.Close()
	

	// ここでパスの確認
	fileName := getSafeFileName()
// 	filePath := config.GetInstance().AssetConfig.GetPrefix() + "/" + fileName
	filePath := "uploads/image/" +  fileName

	// サーバー上に保存するファイルを作成
	// ここをループ
	dst, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		log.Println(err)
	}


	return fileName
}

func getSafeFileName() string {
	// todo: dbに問い合わせて重複のないFILE_PATHを取得する処理を書く
	return getFileName()
}

func getFileName() string {
    c := 40
    b := make([]byte, c)
    rand.Read(b)
	fileName := base64.URLEncoding.EncodeToString(b)

	return fileName
}

// func(self *localFileUploader) Temp (os.File) {
// }
