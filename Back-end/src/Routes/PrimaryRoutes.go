package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func NOTIMPLEMENTED(c *gin.Context, route string) {
	c.JSON(http.StatusOK, "This route has not been implemented yet! route: `/" + route + "` But hello anyways!")
}

func Index(c *gin.Context) {
	NOTIMPLEMENTED(c, "");
}