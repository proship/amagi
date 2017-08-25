package fileStorage

import (
	"fmt"
	"io"
	"os"

	"github.com/b-eee/amagi/services/fileStorage/backends"

	utils "github.com/b-eee/amagi"
	minio "github.com/minio/minio-go"
)

type (
	// File file storage interface
	File struct {
		ObjectName string
		BucketName string
		File       io.Reader
	}
)

// PutObject put object to file storage
func (fs *File) PutObject() (interface{}, error) {
	req := backends.FileObject{
		File:       fs.File,
		ObjectName: fs.ObjectName,
		BucketName: fs.BucketName,
	}
	resp, err := backends.PutObject(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetObject get object from storage
func (fs *File) GetObject() (interface{}, error) {
	req := backends.FileObject{
		BucketName: fs.BucketName,
		ObjectName: fs.ObjectName,
	}

	return backends.GetObject(req)
}

// MIOExtractAndStoreObject extract object to minio *Object and store locally
func MIOExtractAndStoreObject(object interface{}, filepath string) error {
	localfile, err := os.Create(filepath)
	if err != nil {
		utils.Error(fmt.Sprintf("error MIOExtractAndStoreObject create %v", err))
		return err
	}

	if _, err := io.Copy(localfile, object.(*minio.Object)); err != nil {
		utils.Error(fmt.Sprintf("error MIOExtractAndStoreObject copy %v", err))
		return err
	}

	return nil
}

// InitializeServerStorages initliaze or create client to server storages for server common storages
func InitializeServerStorages() error {

	return nil
}