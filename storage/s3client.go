package storage

import "io"

// S3Client is a S3 client.
type S3Client interface {
	GetObject(objectName string) (S3Object, error)
	PutObject(objectName, contentType string, r io.Reader, size int64) (int, error)
	RemoveObject(objectName string) error
}

// S3Object is an object stored in S3.
type S3Object interface {
	io.Reader
	io.Seeker
	io.ReaderAt
	io.Closer
}
