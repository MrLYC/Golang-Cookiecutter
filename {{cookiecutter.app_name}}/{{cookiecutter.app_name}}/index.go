package {{cookiecutter.app_name}}

import (
	"github.com/gin-gonic/gin"
)

// IndexView : index of auth server
func IndexView(c *gin.Context) {
	c.String(200, "hello {{cookiecutter.app_name}}")
}
