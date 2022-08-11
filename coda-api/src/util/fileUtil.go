package util

import (
	"context"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// GetFilePathForGcs Gcs用のパス取得
func GetFilePathForGcs(
	fileName string,
	fileHeder *multipart.FileHeader,
	folder string,
) string {
	if fileHeder != nil {
		ext := filepath.Ext(fileHeder.Filename)
		return folder + "/" + fileName + ext
	}
	return folder + "/" + fileName
}

// GetFileNameFromURL URLからファイ名を取得
func GetFileNameFromURL(
	imageURL string,
) string {
	return filepath.Base(imageURL)
}

// UploadNewFileToGcs Gcsに新しいファイルアップロードをする
func UploadNewFileToGcs(
	path string,
	file multipart.File,
) (string, error) {
	ctx := context.Background()
	client, err := getGcsClient(ctx)
	if err != nil {
		return "", err
	}
	sw := client.Bucket(os.Getenv("GCS_BUCKET")).Object(path).NewWriter(ctx)
	if _, err := io.Copy(sw, file); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}
	u, _ := url.Parse("/" + os.Getenv("GCS_BUCKET") + "/" + sw.Attrs().Name)
	return "https://storage.googleapis.com" + u.EscapedPath(), nil
}

// DeleteFileFromGcs 指定されたパスのGcsのファイルをDELETE
func DeleteFileFromGcs(
	path string,
) error {
	ctx := context.Background()
	client, err := getGcsClient(ctx)
	if err != nil {
		return err
	}
	obj := client.Bucket(os.Getenv("GCS_BUCKET")).Object(path)
	if err := obj.Delete(ctx); err != nil {
		return err
	}
	return nil
}

// getGcsClient Gcsのクライアント取得
func getGcsClient(ctx context.Context) (*storage.Client, error) {
	credentialFilePath := "./gcsKey/" + os.Getenv("GCS_KEY_FILE")
	return storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
}

// GetMultipartFiles 複数ファイルの取得
func GetMultipartFiles(c *gin.Context, propertyName string) ([]multipart.File, []*multipart.FileHeader, error) {
	err := c.Request.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, nil, err
	}
	formdata := c.Request.MultipartForm
	filHeaders := formdata.File[propertyName+"[]"]
	var files []multipart.File
	for _, h := range filHeaders {
		file, err := h.Open()
		if err != nil {
			return nil, nil, err
		}
		files = append(files, file)
	}
	return files, filHeaders, nil
}

// CloseMultipartFiles 複数ファイルのクローズ処理
func CloseMultipartFiles(files []multipart.File) {
	for _, f := range files {
		f.Close()
	}
}
