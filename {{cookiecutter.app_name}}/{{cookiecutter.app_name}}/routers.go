package {{cookiecutter.app_name}}

import (
	"github.com/gin-gonic/gin"
)

func route(enging *gin.Engine) {
	enging.GET("/", IndexView)
}
