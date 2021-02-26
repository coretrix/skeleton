package mocks

import (
	"time"

	ossService "coretrix/skeleton/pkg/ioc/service/oss"

	"github.com/stretchr/testify/mock"
	"github.com/summer-solutions/orm"
)

type FakeOSSClient struct {
	mock.Mock
}

func (t *FakeOSSClient) GetObjectURL(bucket string, object *ossService.Object) string {
	return t.Called(bucket, object).String(0)
}

func (t *FakeOSSClient) GetObjectSignedURL(bucket string, object *ossService.Object, expires time.Time) string {
	return t.Called(bucket, object, expires).String(0)
}

func (t *FakeOSSClient) UploadObjectFromFile(_ *orm.Engine, bucket, localFile string) ossService.Object {
	return t.Called(bucket, localFile).Get(0).(ossService.Object)
}

func (t *FakeOSSClient) UploadObjectFromBase64(_ *orm.Engine, bucket, content, extension string) ossService.Object {
	return t.Called(bucket, content, extension).Get(0).(ossService.Object)
}

func (t *FakeOSSClient) UploadImageFromFile(_ *orm.Engine, bucket, localFile string) ossService.Object {
	return t.Called(bucket, localFile).Get(0).(ossService.Object)
}

func (t *FakeOSSClient) UploadImageFromBase64(_ *orm.Engine, bucket, image, extension string) ossService.Object {
	return t.Called(bucket, image, extension).Get(0).(ossService.Object)
}
