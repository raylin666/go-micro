package qiniu

import "github.com/raylin666/go-utils/upload/qiniu"

var (
	q *qiniu.Options
)

func New(accessKey string, secretKey string, bucket string, zone string, c *qiniu.Config) *qiniu.Options {
	q = qiniu.New(accessKey, secretKey, bucket, zone, c)
	return q
}

func Get() *qiniu.Options {
	return q
}
