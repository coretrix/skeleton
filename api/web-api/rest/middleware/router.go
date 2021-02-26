package middleware

import (
	"coretrix/skeleton/api/web-api/rest/controller"

	"github.com/gin-gonic/gin"
)

func Router(ginEngine *gin.Engine) {
	var devPanel *controller.DevPanelController
	{
		ginEngine.GET("/dev/action-list/", devPanel.GetActionListAction)
	}

	var dev *controller.DevController
	{
		ginEngine.GET("/dev/get-env/", dev.ShowENV)
	}
}
