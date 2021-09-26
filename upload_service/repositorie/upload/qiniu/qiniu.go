package qiniu

import "github.com/raylin666/go-utils/upload/qiniu"

var (
	q *qiniu.Options
)

type UploadPutRet struct {
	Hash         string `json:"hash"`
	Key          string `json:"key"`
	Fsize        int64	`json:"fsize"`
	Url 		 string `json:"url"`
	Bucket		 string `json:"bucket"`
	MimeType	 string `json:"mimeType"`
	Name	 	 string `json:"name"`
	Uuid	 	 string `json:"uuid"`
	Ext			 string `json:"ext"`
}

func New(accessKey string, secretKey string, bucket string, zone string, c *qiniu.Config) *qiniu.Options {
	q = qiniu.New(accessKey, secretKey, bucket, zone, c)
	q.PutPolicy.ReturnBody = `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","mimeType":"$(mimeType)","name":"$(fprefix)","uuid":"$(uuid)","ext":"$(ext)"}`
	q.PutRet = UploadPutRet{}
	return q
}

func Get() *qiniu.Options {
	return q
}
