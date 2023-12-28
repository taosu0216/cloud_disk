package utils

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func COSUploadFile(r *http.Request) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("----------------------------------------------", "env load fail err is : ", err)
		return "", err

	}
	u, _ := url.Parse(os.Getenv("COSPATH"))
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("SECRETID"),
			SecretKey: os.Getenv("SecretKey"),
		},
	})
	file, fileHeader, err := r.FormFile("file")
	key := "cloud_disk/" + GenerateUUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(
		context.Background(), key, file, nil,
	)
	if err != nil {
		log.Fatalln("-------------------------------------------", "upload fail err is : ", err)
		return "", err
	}
	return os.Getenv("COSPATH") + "/" + key, nil
}
