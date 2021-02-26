package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type DevController struct {
}

func (controller *DevController) ShowENV(c *gin.Context) {
	envVar := c.Query("env")

	c.JSON(http.StatusOK, os.Getenv(envVar))
}
