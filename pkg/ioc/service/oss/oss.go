package oss

import (
	"time"

	"github.com/summer-solutions/orm"
)

type Client interface {
	GetObjectURL(bucket string, object *Object) string
	GetObjectSignedURL(bucket string, object *Object, expires time.Time) string
	UploadObjectFromFile(ormService *orm.Engine, bucket, localFile string) Object
	UploadObjectFromBase64(ormService *orm.Engine, bucket, content, extension string) Object
	UploadImageFromFile(ormService *orm.Engine, bucket, localFile string) Object
	UploadImageFromBase64(ormService *orm.Engine, bucket, image, extension string) Object
}

const BucketShow = "hymn-show"

var BucketsList = map[string]uint64{
	BucketShow: 1,
}

type Object struct {
	ID         uint64
	StorageKey string
	Data       interface{}
}
