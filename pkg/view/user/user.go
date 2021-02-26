package user

import (
	"github.com/coretrix/hitrix/example/graph/model"
)

func GetUser() (*model.User, error) {
	//your logic here
	//ossService := ioc.GetOSSService()
	//ossService.GetObjectURL()

	return &model.User{
		ID: "1234",
	}, nil
}
